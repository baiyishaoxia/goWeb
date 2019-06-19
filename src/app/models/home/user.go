package models

import (
	"app"
	"databases"
	"fmt"
	"time"
)

type Users struct {
	Id          int64     `xorm:"pk autoincr BIGINT" json:"id"`
	Name        string    `xorm:"not null unique VARCHAR(255)" json:"name"`
	Sex         int64     `xorm:"not null default 1 INTEGER" json:"sex"`
	Phone       string    `xorm:"VARCHAR(255)" json:"phone"`
	Email       string    `xorm:"VARCHAR(255)" json:"email"`
	City        string    `xorm:"VARCHAR(255)" json:"city"`
	HeadImg     string    `xorm:"VARCHAR(255)" json:"head_img"`
	Password    string    `xorm:"VARCHAR(255)" json:"password"`
	IsLock      bool      `xorm:"default true" json:"is_lock"`
	LoginCount  int64     `xorm:"not null default 0 INTEGER" json:"login_count"`
	LastLogin   time.Time `json:"last_login"`
	CreatedAt   app.Time  `xorm:"created" json:"created_at"`
	UpdatedAt   app.Time  `xorm:"updated" json:"updated_at"`
	Nickname    string    `xorm:"- <- ->" json:"-"`
	AreaCode    string    `xorm:"- <- ->" json:"-"`
	Mobile      string    `xorm:"- <- ->" json:"-"`
	Token       string    `xorm:"- <- ->" json:"-"`
	FailureTime time.Time `xorm:"- <- ->" json:"-"` //失效时间
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
