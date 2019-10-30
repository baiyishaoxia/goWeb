package background

import (
	"app/models"
	"databases"
	"fmt"
	"math"
)

//获取列表
func PagebannerCategoryList(page int, limit int, keywords string) (*[]models.BannerCategory, float64, float64, int) {
	data := new([]models.BannerCategory)
	err := databases.Orm.Table("banner_category").Desc("id")
	err1 := *err
	num, _ := err.Count()
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	//数据
	err2 := err1.Limit(limit, page*limit-limit).Find(data)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return data, float64(num), all, page + 1
}
