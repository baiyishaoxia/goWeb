package models

import (
	"app"
	"databases"
	"fmt"
	"math"
	"time"
)

type Users struct {
	Id         int64  `xorm:"pk autoincr BIGINT"`
	Name       string `xorm:"not null unique VARCHAR(255)"`
	Sex        int64  `xorm:"not null default 1 INTEGER"`
	Phone      string `xorm:"VARCHAR(255)"`
	Email      string `xorm:"VARCHAR(255)"`
	City       string `xorm:"VARCHAR(255)"`
	HeadImg    string `xorm:"VARCHAR(255)"`
	Password   string `xorm:"VARCHAR(255)"`
	IsLock     bool   `xorm:"default true"`
	LoginCount int64  `xorm:"not null default 0 INTEGER"`
	LastLogin  time.Time
	CreatedAt  app.Time `xorm:"created"`
	UpdatedAt  app.Time `xorm:"updated"`
}

//region Remark:列表 Author:tang
func GetUsersList(page int, limit int, keywords string) (*[]Users, float64, float64, int) {
	var users = new([]Users)
	err := databases.Orm.Desc("id").Table("users")
	if keywords != "" {
		err.Where("name like ?", "%"+keywords+"%")
	}
	err1 := *err
	num, err3 := err1.Count()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	err2 := err.Limit(limit, page*limit).Find(users)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return users, float64(num), all, page + 1
}

//endregion

//region Remark:查询用户是否已经存在 Author:tang
func GetUserExits(name string, id int64) bool {
	user := new(Users)
	has, _ := databases.Orm.Where("name =? and id !=?", name, id).Get(user)
	if has == true {
		return true
	}
	return false
}

//endregion

//region Remark:根据id获取用户数据 Author:tang
func GetUserInfoById(user_id int64) *Users {
	user := new(Users)
	databases.Orm.Id(user_id).Get(user)
	if user.HeadImg == "" {
		user.HeadImg = "/public/background/static/h-ui/images/ucnter/avatar-default.jpg"
	}
	return user
}

//endregion
