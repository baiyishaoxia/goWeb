package background

import (
	"app"
	"databases"
	"fmt"
	"math"
)

//banner图类别表
type BannerCategory struct {
	Id        int64    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	Title     string   `xorm:"VARCHAR(255)" json:"title"`
	Index     string   `xorm:"VARCHAR(255)" json:"index"`
	Intro     string   `xorm:"VARCHAR(255)" json:"intro"`
	CreatedAt app.Time `xorm:"created" json:"created_at"`
	UpdatedAt app.Time `xorm:"updated" json:"updated_at"`
}

//获取列表
func PagebannerCategoryList(page int, limit int, keywords string) (*[]BannerCategory, float64, float64, int) {
	data := new([]BannerCategory)
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
