package models

import (
	"app"
	"databases"
	"fmt"
	"math"
)

type Video struct {
	Id        int64    `xorm:"pk autoincr BIGINT"`
	Title     string   `xorm:"not null VARCHAR(255)" json:"title"`
	Remark    string   `xorm:"VARCHAR(255)" json:"remark"`
	Url       string   `xorm:"VARCHAR(255)" json:"url"`
	ImgUrl    string   `xorm:"VARCHAR(255)" json:"img_url"`
	CreatedAt app.Time `xorm:"created" json:"created_at"`
	UpdatedAt app.Time `xorm:"updated" json:"-"`
}

//region Remark:列表 Author:tang
func VideoList(page int, limit int, keywords string) (*[]Video, float64, float64, int) {
	var video = new([]Video)
	err := databases.Orm.Desc("id")
	if keywords != "" {
		err.Where("remark = ? or title like ?", keywords, "%"+keywords+"%")
	}
	err1 := *err
	//记录数
	num, err3 := err1.Table("video").Count()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	//分页
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	err2 := err.Limit(limit, page*limit).Find(video)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return video, float64(num), all, page + 1
}

//endregion
func (a *Video) AddVideo() bool {
	ok, err := databases.Orm.Insert(a)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if ok < 1 {
		fmt.Println("添加失败")
		return false
	}
	return true
}

//通过id查询视频
func GetVideoById(id string) *Video {
	video := new(Video)
	has, err := databases.Orm.Where("id = ?", id).Get(video)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return video
}
