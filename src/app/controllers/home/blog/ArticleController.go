package blog

import (
	"app"
	"app/models/background"
	models2 "app/models/home"
	"app/service/home"
	"config"
	"databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//文章专栏
func GetBlogArticle(c *gin.Context) {
	c.HTML(http.StatusOK, "default/article", gin.H{
		"Title": "欢迎使用GO语言编程",
	})
}

//Ajax请求文章列表
func PostArticleList(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64) //分类的id
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultPostForm("limit", "10"))
	keywords := c.PostForm("keywords")
	wheres := map[string]interface{}{
		"keywords": keywords,
	}
	var category_id int64 = 0
	if id != 0 {
		cate := new(models.Category)
		has, err := databases.Orm.Where("id=?", id).Get(cate)
		category_id = cate.Id
		if !has || err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"data": gin.H{
					"info": "该分类不存在，无法加载内容",
				},
			})
			return
		}
	}
	news, num, all, page_c := home.GetNewsByCategoryId(category_id, limit, page-1, wheres)
	var data = make([]map[string]interface{}, len(*news))
	for key, val := range *news {
		data[key] = make(map[string]interface{})
		data[key]["id"] = val.Id
		data[key]["title"] = val.Title
		data[key]["image"] = val.Img
		data[key]["intro"] = val.Intro
		data[key]["keywords"] = val.Keywords
		data[key]["author"] = models.GetArticleAuthorById(int(val.AuthorId))
		data[key]["cate_title"] = models.GetCategoryById(val.CateId).Title
		data[key]["cate_id"] = val.CateId
		data[key]["click_num"] = val.ClickNum
		data[key]["count_num"] = val.CountNum
		data[key]["source"] = val.Source
		data[key]["is_top"] = val.IsTop
		data[key]["count_num"] = val.CountNum
		data[key]["created_at"] = time.Time(val.CreatedAt).Format("2006-01-02")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"news": data,
			"num":  num,
			"all":  all,
			"page": page_c,
		},
	})
	return
}

//Ajax请求文章分类
func PostArticleCatgory(c *gin.Context) {
	_, category := models2.GetCategory()
	data := make([]map[string]interface{}, len(*category))
	for key, val := range *category {
		data[key] = make(map[string]interface{})
		data[key]["id"] = val.Id
		data[key]["title"] = val.Title
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"category": data,
		},
	})
	return
}

//文章推荐
func GetArticleRight(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultPostForm("limit", "10"))
	red_data, click_data := home.GetNewsByRight(limit)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"red":   red_data,
			"click": click_data,
		},
	})
	return
}

//文章详情
func GetBlogArticleDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data, preNext := home.GetArticleById(id)
	c.HTML(http.StatusOK, "default/detail", gin.H{
		"Title":   "欢迎使用GO语言编程",
		"Data":    data,
		"PreNext": preNext,
		"StrSub": func(str string) string {
			if str == "" {
				return "无"
			}
			if strings.Count(str, "")-1 > 10 {
				return app.SubString(str, 0, 8) + "..."
			}
			return str
		},
	})
}
