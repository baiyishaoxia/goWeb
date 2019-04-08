package models

import (
	"app"
	"databases"
	"fmt"
	"time"
)

type Users struct {
	Id          int64     `json:"id"`
	Name        string    `json:"account"`
	Nickname    string    `json:"nickname"`
	Mobile      string    `json:"mobile"`
	Password    string    `json:"-"`
	Sex         int64     `json:"sex"`
	IsLock      bool      `json:"is_lock"`
	HeadUrl     string    `json:"head_url"`
	AreaCode    string    `json:"area_code"`
	LimitTime   int64     `json:"-"`
	Token       string    `json:"token"`
	FailureTime time.Time `json:"failure_time"` //失效时间
	CreatedAt   app.Time  `xorm:"created" json:"created_at"`
	UpdatedAt   app.Time  `xorm:"updated" json:"-"`
}

//region Remark: 通过id查询用户 Author: tang
func GetUserById(id int64) *Users {
	user := new(Users)
	has, err := databases.Orm.Where("id = ?", id).Get(user)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return user
}

//endregion

//region   通过id获取用户信息   Author:tang
func GetUserByToken(token string) *Users {
	user := new(Users)
	has, err := databases.Orm.Where("token = ?", token).Where("failure_time>?", time.Now().Format("2006-01-02 15:04:05")).Get(user)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return user
}

//endregion
