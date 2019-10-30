package background

import (
	"app"
	"app/models"
	"databases"
	"encoding/json"
	"fmt"
	"math"
)

//关联表
type BannerData struct {
	models.Banner         `xorm:"extends"`
	models.BannerCategory `xorm:"extends" json:"banner_category"`
}

//获取列表
func PageBannerList(keywords string, banner_category_id int64, limit int, page int) (*[]BannerData, float64, float64, int) {
	banner := new([]BannerData)
	err := databases.Orm.Table("banner").
		Join("LEFT", "`banner_category`", "`banner_category`.id = `banner`.banner_category_id").
		Asc("banner.sort").Desc("banner.id")
	if keywords != "" {
		err.Where("banner.title=?", keywords)
	}
	if banner_category_id != 0 {
		err.Where("banner.banner_category_id=?", banner_category_id)
	}
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
	err2 := err1.Limit(limit, page*limit).Find(banner)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return banner, float64(num), all, page + 1
}
func PageBannerListAjax(keywords string, banner_category_id int64, limit int, page int) string {
	banner := new([]BannerData)
	err := databases.Orm.Table("banner").
		Join("LEFT", "`banner_category`", "`banner_category`.id = `banner`.banner_category_id")
	if keywords != "" {
		err.Where("banner.title=?", keywords)
	}
	if banner_category_id != 0 {
		err.Where("banner.banner_category_id=?", banner_category_id)
	}
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
	err2 := err1.Limit(limit, page*limit-limit).Find(banner)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	value, _ := json.Marshal(banner)
	return app.TableJson(0, num, "", string(value[:]))
}
func HasBanners(ids []string) bool {
	item := new(models.Banner)
	count, err := databases.Orm.Table("`banner`").In("banner_category_id", ids).Count(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		return true
	}
	return false
}
