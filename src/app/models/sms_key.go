package models

import (
	"app"
)

type SmsKey struct {
	Id        int32    `json:"id"`
	SmsId     string   `json:"sms_id"` // 短信ID
	Name      string   `json:"name"`   // 名称
	Key       string   `json:"key"`    // 账号签名
	CreatedAt app.Time `xorm:"created" json:"created_at"`
	UpdatedAt app.Time `xorm:"updated" json:"-"`
}
