package models

import (
	"app"
	"databases"
	"fmt"
)

type Sms struct {
	Id        int32    `json:"id"`
	Name      string   `json:"name"`                        // 短信类型
	Sign      string   `json:"sign"`                        // 短信签名
	IsEnable  bool     `json:"not null default false BOOL"` // 是否开启true 开启，false关闭
	Key       string   `json:"key"`                         // 唯一标识分类
	CreatedAt app.Time `xorm:"created" json:"created_at"`
	UpdatedAt app.Time `xorm:"updated" json:"-"`
}

type SmsInfo struct {
	Sms    `xorm:"extends"`
	SmsKey `xorm:"extends"`
}

//region   获取模板签名、秘钥  [prefix:手机号前缀  YunPian：中文签名  YunPianEng：英文签名]  Author:tang
func GetSmsInfo(prefix string) *SmsInfo {
	sms_info := new(SmsInfo)
	if prefix == "+86" {
		has, err := databases.Orm.Table("sms").Join("LEFT", "sms_key", "sms.id = sms_key.sms_id").
			Where("sms.key = ?", "YunPian").Get(sms_info)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		if has == false {
			fmt.Println("获取YunPian失败")
			return nil
		}
	} else {
		has, err := databases.Orm.Table("sms").Join("LEFT", "sms_key", "sms.id = sms_key.sms_id").
			Where("sms.key = ?", "YunPianEng").Get(sms_info)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		if has == false {
			fmt.Println("获取YunPianEng失败")
			return nil
		}
	}
	return sms_info
}

//endregion
