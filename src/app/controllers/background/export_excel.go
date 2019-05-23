package background

import (
	"app"
	"app/models/background"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"strconv"
)

//region Remark:导出 Author:tang
func GetSearchAllCsv(c *gin.Context) {
	fileName := "search.csv"
	//获取数据源
	keywords := c.Query("keywords")
	imgHost := c.Request.Host
	way := c.Query("way")
	typeKey := c.Query("type_key")
	limit, _ := strconv.ParseInt(models.ReadConfig("sys.paginate"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	if page == 0 {
		page = 1
	}
	list, _, _, _ := models.SearchArticleBykeys(keywords, imgHost, limit, page, way, typeKey)
	fmt.Println("list:", len(list))
	//处理数据格式
	searchStrArray := structs2StringArray(list)
	fmt.Println("searchStrArray:", len(searchStrArray))
	searchHeader := []string{"标题", "简介", "类别", "内容", "作者", "浏览量", "创建时间"}
	b := &bytes.Buffer{}
	b.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	wr := csv.NewWriter(b)
	wr.Write(searchHeader) //按行输出
	for i := 0; i < len(searchStrArray); i++ {
		wr.Write(searchStrArray[i])
	}
	footHeader := []string{"*注:作者的数字代表 1:白衣少侠, 2: 阿猛, 3: 池建, 4: 邹琴"}
	wr.Write(footHeader)
	wr.Flush()
	//导出下载
	app.DownCsv(c, fileName, b)
}

//endregion
func structs2StringArray(datas []models.Article) [][]string {
	var searchArr = make([][]string, 0)
	for _, val := range datas {
		searchArr = append(searchArr, []string{
			val.Title,
			val.Intro,
			val.CateName,
			string(val.Content),
			strconv.Itoa(int(val.AuthorId)),
			strconv.FormatFloat(float64(val.ClickNum), 'f', 2, 64),
			val.CreatedAt.String(),
		})
	}
	return searchArr
}
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
