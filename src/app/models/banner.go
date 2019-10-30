package models

import "app"

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
