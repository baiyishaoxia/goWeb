package models

import (
	"app"
	"databases"
	"fmt"
	"math"
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
	HotCount    int64     `xorm:"not null default 0 INTEGER" json:"hot_count"` //热评次数
	Level       int64     `xorm:"not null default 0 INTEGER" json:"level"`     //等级(潜水,活跃等)
	CreatedAt   app.Time  `xorm:"created" json:"created_at"`
	UpdatedAt   app.Time  `xorm:"updated" json:"updated_at"`
	Nickname    string    `xorm:"- <- ->" json:"-"`
	AreaCode    string    `xorm:"- <- ->" json:"-"`
	Mobile      string    `xorm:"- <- ->" json:"-"`
	Token       string    `xorm:"- <- ->" json:"-"`
	FailureTime time.Time `xorm:"- <- ->" json:"-"` //失效时间
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

//region   验证用户信息是否失效   Author:tang
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

//region   为用户增加热评次数   Author:tang
func UserHotCountInrc(user_id int64) {
	item := new(Users)
	res, _ := databases.Orm.Where("id=?", user_id).Get(item)
	if res {
		item.HotCount = item.HotCount + 1
		databases.Orm.Cols("hot_count").Update(item, Users{Id: item.Id})
	}
}

//endregion

//region   获取所有用户   Author:tang
func GetUsersList2() []Users {
	item := make([]Users, 0)
	databases.Orm.Desc("id").Find(&item)
	return item
}

//endregion

//region   活跃度升级方式   Author:tang
func GetLevelUpdate(level int64) string {
	item := "新手"
	switch level {
	case 1:
		item = "潜水"
		break
	case 2:
		item = "活跃"
		break
	case 3:
		item = "吐槽"
		break
	case 4:
		item = "传说"
		break
	case 5:
		item = "至尊"
		break
	}
	return item
}

//endregion