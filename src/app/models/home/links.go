package models

import (
	"databases"
	"fmt"
	"math"
)

type BlogLinks struct {
	LinkId    int64  `xorm:"pk autoincr BIGINT"`
	LinkName  string `json:"name"`
	LinkTitle string `json:"title"`
	LinkUrl   string `json:"url"`
	LinkOrder int64  `json:"sort"`
	LinkIsdel int    `json:"isdel"`
}

func GetLinksList(page int, limit int, keywords string) (*[]BlogLinks, float64, float64, int) {
	links := new([]BlogLinks)
	err := databases.Orm.Desc("link_order").Where("link_isdel = ?", 0)
	if keywords != "" {
		err.Where("link_name like ?", "%"+keywords+"%")
	}
	err1 := *err
	//记录数
	num, err3 := err1.Table("blog_links").Count()
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
	err2 := err.Limit(limit, page*limit).Find(links)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return links, float64(num), all, page + 1
}

//region Remark:提取当前id信息 Author:tang
func GetLinksById(id int64) *BlogLinks {
	var links = new(BlogLinks)
	has, err := databases.Orm.Where("link_id = ?", id).Get(links)
	if has == false {
		return nil
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return links
}

//endregion
