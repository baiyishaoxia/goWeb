package models

import (
	"databases"
	"fmt"
	"math"
	"time"
)

type AdminLog struct {
	Id        int64     `xorm:"pk autoincr BIGINT"`
	AdminId   int64     `xorm:"BIGINT"`
	Ip        string    `xorm:"not null VARCHAR(32)"`
	Area      string    `xorm:"not null VARCHAR(64)"`
	Url       string    `xorm:"not null VARCHAR(255)"`
	Type      string    `xorm:"not null VARCHAR(10)"`
	Request   string    `xorm:"not null TEXT"`
	Remark    string    `xorm:"TEXT"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
type Logs struct {
	AdminLog  `xorm:"extends"`
	BlogAdmin `xorm:"extends"`
}

//region Remark:后台日志记录列表 Author:tang
func GetAdminLogList(page int, limit int, keywords string) (*[]Logs, float64, float64, int) {
	admin_log := new([]Logs)
	err := databases.Orm.Table("admin_log").
		Join("LEFT", "blog_admin", "admin_log.admin_id = blog_admin.id")
	if keywords != "" {
		err.Where("blog_admin.username = ?", keywords)
	}
	err1 := *err
	//记录数
	num, err2 := err.Count()
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	//分页
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	err3 := err1.Desc("admin_log.id").Limit(limit, page*limit).Find(admin_log)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return admin_log, float64(num), all, page + 1
}

//endregion
