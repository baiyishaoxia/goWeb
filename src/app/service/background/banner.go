package background

import (
	"app"
	"databases"
	"fmt"
	"math"
)

//banner图内容表
type Banner struct {
	Id               int64    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	Title            string   `xorm:"VARCHAR(255)" json:"title"`              //标题
	Url              string   `xorm:"VARCHAR(255)" json:"url"`                //链接
	Image            string   `xorm:"VARCHAR(255)" json:"image"`              //图片
	Sort             int64    `xorm:"BIGINT"       json:"sort"`               //排序
	BannerCategoryId int64    `xorm:"BIGINT"       json:"banner_category_id"` //类别
	Intro            string   `xorm:"VARCHAR(255)" json:"intro"`              //简介
	Abstract         string   `xorm:"VARCHAR(255)" json:"abstract"`           //摘要
	Content          string   `xorm:"TEXT"    json:"content"`                 //内容
	CreatedAt        app.Time `xorm:"created" json:"created_at"`
	UpdatedAt        app.Time `xorm:"updated" json:"updated_at"`
}

type BannerData struct {
	Banner         `xorm:"extends"`
	BannerCategory `xorm:"extends" json:"banner_category"`
}

//获取列表
func PageBannerList(keywords string, banner_category_id int64, limit int, page int) (*[]BannerData, float64, float64, int) {
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
	return banner, float64(num), all, page + 1
}
func HasBanners(ids []string) bool {
	item := new(Banner)
	count, err := databases.Orm.Table("`banner`").In("banner_category_id", ids).Count(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		return true
	}
	return false
}
