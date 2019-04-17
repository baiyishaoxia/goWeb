package models

import (
	"app"
	"databases"
	"fmt"
	"math"
)

type Picture struct {
	Id        int64    `xorm:"pk autoincr BIGINT"`
	Title     string   `xorm:"not null unique VARCHAR(255)"`
	CateId    int64    `xorm:"bigint" json:"-"`
	AuthorId  int64    `xorm:"bigint" json:"-"`
	Sort      int      `xorm:"default 99 INTEGER"`
	IsComment bool     `xorm:"default true"`
	Img       string   `xorm:"VARCHAR(255)"`
	Source    string   `xorm:"VARCHAR(255)"`
	Keywords  string   `xorm:"VARCHAR(255)"`
	Intro     string   `xorm:"VARCHAR(255)"`
	Images    string   `xorm:"TEXT"`
	StartTime string   `xorm:"VARCHAR(255)"`
	EndTime   string   `xorm:"VARCHAR(255)"`
	Status    int      `xorm:"not null default 1 INTEGER"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
	CateName  string   `xorm:"- <- ->"`
}

type PictureAndCategory struct {
	Picture  `xorm:"extends"`
	Category `xorm:"extends"`
}

//region Remark:相册列表 Author:tang
func GetPictureList(page, limit int64, keywords string, category_id int64, author_id int64) (*[]PictureAndCategory, float64, float64, int64) {
	picture := new([]PictureAndCategory)
	err := databases.Orm.Table("picture").Join("left", "category", "picture.cate_id = category.id").Desc("picture.id")
	if keywords != "" {
		err.Where("picture.title like ?", "%"+keywords+"%")
	}
	if category_id != 0 {
		err.Where("cate_id = ?", category_id)
	}
	if author_id != 0 {
		err.Where("author_id = ?", author_id)
	}
	err1 := *err
	//记录数
	num, err2 := err.Table("picture").Count()
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	//分页
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int64(all) {
		page = int64(all) - 1
	}
	err3 := err1.Limit(int(limit), int(page*limit)).Find(picture)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return picture, float64(num), all, page + 1
}

//endregion

//region Remark:相册----根据id获取记录并格式化数据 Author:tang
func GetPictureInfo(id int64) (*Picture, map[int]map[string]interface{}) {
	picture := new(Picture)
	_, err := databases.Orm.Id(id).Select("id,title,images").Get(picture)
	if err != nil {
		fmt.Println(err.Error())
	}
	//格式化源数据
	data := app.StrSplitArray(picture.Images)
	images := make(map[int]map[string]interface{}, 0)
	for key, val := range data {
		images[key] = make(map[string]interface{})
		images[key]["Id"] = picture.Id
		images[key]["Title"] = picture.Title
		images[key]["Image"] = val
	}
	return picture, images
}

//endregion
