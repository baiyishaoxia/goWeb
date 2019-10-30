package models

import "time"

//时光轴表
type TimeLine struct {
	Id        int64     `xorm:"pk autoincr BIGINT"`
	Title     string    `xorm:"VARCHAR(255)"`
	Content   string    `xorm:"TEXT"`              //内容
	Time      time.Time `xorm:"DATETIME"`          //时间
	IsShow    bool      `xorm:"bool default true"` //是否显示
	Year      int64     `xorm:"BIGINT"`            //当前时间年份
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
