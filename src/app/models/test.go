package models

import (
	"databases"
	"fmt"
	"github.com/go-xorm/xorm"
	"math"
	"time"
)

type BlogToolsContent struct {
	Id        int64
	Title     string
	CreatedAt time.Time
}

type BlogToolsContentAttache struct {
	Id        string
	ContentId string
	Filepath  string
	Filesize  string
	DownCount int64
}

//region Remark:连表测试 Author:tang
type ContentAndAttache struct {
	BlogToolsContent        `xorm:"extends"`
	BlogToolsContentAttache `xorm:"extends"`
}

//endregion

func GetTestList(page int, limit int, keywords string) (*[]ContentAndAttache, float64, float64, int) {
	data := new([]ContentAndAttache)
	err := databases.Orm.Table("blog_tools_content").Join("LEFT", "blog_tools_content_attache", "blog_tools_content_attache.content_id = blog_tools_content.id").
		Desc("blog_tools_content.id")
	num, all, page := Getpage(*err, page, limit)
	err2 := err.Limit(limit, page*limit).Find(data)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return data, float64(num), all, page + 1
}

//region Remark:数据分页 Author:tang
func Getpage(err xorm.Session, page int, limit int) (float64, float64, int) {
	//记录数
	num, err3 := err.Table("blog_tools_content").Count()
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
	return float64(num), all, page
}

//endregion
