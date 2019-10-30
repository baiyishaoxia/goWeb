package statistical

import (
	"app/models"
	"databases"
)

func GetDataCountSum() map[string]map[int]int64 {

	//统计资讯库
	article := new(models.Article)
	article_totals, _ := databases.Orm.Count(article)

	//统计管理员
	admin := new(models.BlogAdmin)
	admin_totals, _ := databases.Orm.Count(admin)

	fo := make(map[string]map[int]int64)
	fo["sum"] = make(map[int]int64)
	fo["sum"][0] = article_totals
	fo["sum"][1] = 0
	fo["sum"][2] = 0
	fo["sum"][3] = 0
	fo["sum"][4] = admin_totals
	return fo

	//---------组合---------
	//data := make(map[string][]int64)
	//总数
	//data["sum"] = append(data["sum"], article_totals)
	//data["sum"] = append(data["sum"], 0)
	//data["sum"] = append(data["sum"], 0)
	//data["sum"] = append(data["sum"], 0)
	//data["sum"] = append(data["sum"], admin_totals)
	//return data
}
