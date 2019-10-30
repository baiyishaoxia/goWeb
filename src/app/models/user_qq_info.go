package models

import (
	"databases"
	"encoding/json"
	"fmt"
	"time"
)

//QQ登录后的用户信息表
type UserQqInfo struct {
	Id            int64     `xorm:"pk autoincr BIGINT" json:"id"`
	UserId        int64     `xorm:"not null BIGINT" json:"user_id"`
	NickName      string    `xorm:"not null VARCHAR(255)" json:"nick_name"` //用户在QQ空间的昵称
	Openid        string    `xorm:"not null VARCHAR(255)" json:"openid"`    //用户唯一标识
	Gender        string    `xorm:"not null VARCHAR(255)" json:"gender"`    //性别
	Image         string    `xorm:"not null VARCHAR(255)" json:"image"`     //头像
	Province      string    `xorm:"VARCHAR(255)" json:"province"`           //省份
	City          string    `xorm:"VARCHAR(255)" json:"city"`               //城市
	Constellation string    `xorm:"VARCHAR(255)" json:"constellation"`      //星座
	Year          string    `xorm:"VARCHAR(255)" json:"year"`               //出生年
	Vip           string    `xorm:"VARCHAR(2)" json:"vip"`                  //VIP用户
	CreatedAt     time.Time `xorm:"created" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated" json:"updated_at"`
}

//根据openid查询用户数据
func GetUserInfoByOpenId(openid string) *UserQqInfo {
	item := new(UserQqInfo)
	res, err := databases.Orm.Where("openid=?", openid).Get(item)
	if err != nil || res == false {
		return nil
	}
	return item
}

//新增QQ用户
func AddQQUser(data string, openid string) (bool, error, int64) {
	var info = map[string]interface{}{}
	err := json.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Println("--------1-------------", err)
		return false, err, 0 //json解析失败
	}
	if res := GetUserInfoByOpenId(openid); res != nil {
		fmt.Println("--------2-------------", err)
		return true, nil, res.UserId //用户已存在
	}
	user := new(Users)
	user.Name = info["nickname"].(string) //QQ昵称

	if info["gender"].(string) == "男" { //性别
		user.Sex = 1
	} else if info["gender"].(string) == "女" {
		user.Sex = 2
	} else {
		user.Sex = 3
	}
	user.City = info["province"].(string) + info["city"].(string)
	user.HeadImg = info["figureurl_qq_2"].(string) //头像 100*100
	user.LoginCount = 1
	db := databases.Orm.NewSession()
	res, err := db.Insert(user)
	qq := new(UserQqInfo) //QQ信息表
	qq.UserId = user.Id
	qq.Openid = openid
	qq.NickName = info["nickname"].(string)
	qq.Gender = info["gender"].(string)
	qq.City = info["city"].(string)
	qq.Province = info["province"].(string)
	qq.Year = info["year"].(string)
	qq.Constellation = info["constellation"].(string)
	qq.Image = info["figureurl_qq_2"].(string)
	qq.Vip = info["vip"].(string)
	res2, err2 := db.Insert(qq)
	if res < 1 || res2 < 1 {
		db.Rollback()
		fmt.Println("--------3-------------", err, err2)
		return false, err, 0
	}
	db.Commit()
	fmt.Println("--------4-------------", err, err2)
	return true, err2, user.Id
}
