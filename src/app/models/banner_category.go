package models

import "app"

//banner图类别表
type BannerCategory struct {
	Id        int64    `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	Title     string   `xorm:"VARCHAR(255)" json:"title"`
	Index     string   `xorm:"VARCHAR(255)" json:"index"`
	Intro     string   `xorm:"VARCHAR(255)" json:"intro"`
	CreatedAt app.Time `xorm:"created" json:"created_at"`
	UpdatedAt app.Time `xorm:"updated" json:"updated_at"`
}
