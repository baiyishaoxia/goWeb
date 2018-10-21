package models

import (
	"app"
	newredis "app/vendors/redis/models"
	session "app/vendors/session/models"
	"databases"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type BlogAdmin struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	AdminRoleId int64  `xorm:"BIGINT"`
	Username    string `xorm:"not null VARCHAR(255)"`
	Password    string `xorm:"not null VARCHAR(500)"`
	Email       string `xorm:"unique VARCHAR(255)"`
	Mobile      string `xorm:"not null VARCHAR(255)"`
	IsLock      bool   `xorm:"not null default false BOOL"`
	LoginCount  int64  `xorm:"not null default 0 INTEGER"`
	LastLogin   time.Time
	CreatedAt   app.Time   `xorm:"created"`
	UpdatedAt   app.Time   `xorm:"updated"`
	Role        *AdminRole `xorm:"- <- ->"`
}

//region Remark:获取当前管理员登录信息 Author:tang
func GetAdminInfo(c *gin.Context) *BlogAdmin {
	var admin *BlogAdmin
	var admin_id int64 = (session.GetSession(c, "admin_id")).(int64)
	//缓存到redis
	var redis_key string = "admin:info:" + strconv.FormatInt(admin_id, 10)
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &admin)
	} else {
		admin = new(BlogAdmin)
		databases.Orm.ID(admin_id).Get(admin)
		//获取当前管理员的角色
		admin_role := new(AdminRole)
		databases.Orm.Id(admin.AdminRoleId).Get(admin_role)
		admin.Role = admin_role
		//缓存到redis
		value, _ := json.Marshal(admin)
		newredis.Set(redis_key, value, 60*60)
	}
	return admin
}

//endregion

func AdminNowRoleNodes(c *gin.Context, handler_name string) bool {
	fmt.Println("验证当前路由是否有权访问！")
	var data []AdminNavigationNode
	var role_id int64 = GetAdminInfo(c).Role.Id
	var admin_id int64 = GetAdminInfo(c).Id
	//判断redis是否有缓存数据
	var redis_key string = "admin:power:" + strconv.FormatInt(admin_id, 10)
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &data)
	} else {
		admin_navigation_nodes := make([]AdminNavigationNode, 0)
		databases.Orm.Cols("route_action").Where("id in(select admin_navigation_node_id from admin_role_node_routes where admin_role_id=?)", role_id).Find(&admin_navigation_nodes)
		//缓存到redis
		value, _ := json.Marshal(admin_navigation_nodes)
		newredis.Set(redis_key, value, 60*60)
		data = admin_navigation_nodes
	}
	for _, v := range data {
		//fmt.Println(v.RouteAction,handler_name,v.RouteAction==handler_name)
		if v.RouteAction == handler_name {
			return true
		}
	}
	return false
}
