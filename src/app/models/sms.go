package models

import (
	models2 "app/models/background"
	"app/models/home"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//region   短信类型 [flag: 内容类型]  Author:tang
func SendSmsByType(flag int64, user_id int64, title string) {
	user := models.GetUserById(user_id)
	prefix := user.AreaCode
	sms_info := models2.GetSmsInfo(prefix)
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
