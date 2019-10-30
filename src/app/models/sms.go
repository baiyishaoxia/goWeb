package models

import (
	"app"
	"databases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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


//region   短信类型 [flag: 内容类型]  Author:tang
func SendSmsByType(flag int64, user_id int64, title string) {
	user := GetUserById(user_id)
	prefix := user.AreaCode
	sms_info := GetSmsInfo(prefix)
	apikey := sms_info.SmsKey.Key // 运营商秘钥
	sign := sms_info.Sms.Sign     // 短信签名
	var text string
	//-------根据flag 调用模板-------
	switch flag {
	case 1:
		text = "【" + sign + "】" + "您好"
		break
	case 2:
		text = "【" + sign + "】" + ""
		break
	case 3:
		text = "【" + sign + "】" + ""
		break
	case 4:
		text = "【" + sign + "】" + ""
		break
	case 5:
		text = "【" + sign + "】" + ""
		break
	case 6:
		text = "【" + sign + "】" + ""
		break
	}
	deal_mobile := prefix + user.Mobile
	url_send_sms := "https://sms.yunpian.com/v2/sms/single_send.json"
	data_send_sms := url.Values{"apikey": {apikey}, "mobile": {deal_mobile}, "text": {text}}
	//发送请求
	resp, err := http.PostForm(url_send_sms, data_send_sms)
	if err != nil {
		fmt.Println(err.Error()) // handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error()) // handle error
	}
	smsData := make(map[string]interface{})
	json.Unmarshal(body, &smsData)
	//回调参数发送成功还是失败
	a := smsData["code"].(float64)
	if a != 0 {
		fmt.Println("发送失败")
	} else {
		fmt.Println("发送成功")
	}
	//endregion
}
