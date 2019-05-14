package background

import (
	"app/controllers/background"
	"app/models/background"
	newredis "app/vendors/redis/models"
	session "app/vendors/session/models"
	"config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//如果redis有管理员信息，将重新写入session
		if res, _ := newredis.Exists("admin_id"); res == true {
			val, _ := redis.String(newredis.Get("admin_id"))
			admin_id, _ := strconv.ParseInt(val, 10, 64)
			session.SetSession(c, "admin_id", admin_id)
		}
		//判断session里面是否存在admin_id，如果不存在跳转到登陆界面
		if session.HasSession(c, "admin_id") == false {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		//验证ip访问权限
		client_type := "admin"
		client_ip := c.ClientIP()
		switch client_type {
		case "admin":
			//读取ip列表
			Ips := models.ReadConfig("ip")
			result := strings.Contains(Ips, client_ip)
			if result {
				c.HTML(http.StatusOK, "login.html", gin.H{
					"status": config.HttpError,
					"info":   "您的IP已被限制访问，请勿非法操作！",
				})
				c.Abort()
			}
			break
		case "home":
			break
		}

		//验证权限
		//fmt.Println("验证开始")
		nowLoginManagerInfo := models.GetAdminInfo(c)
		//上一次请求路由
		handler_name := c.HandlerName()
		fmt.Println("请求路由", handler_name)
		//region 判断当前用户是否有权限访问该资源---如果不是超级管理员，则进行权限节点的校验
		if nowLoginManagerInfo.Role.IsSuper == false {
			is_power := models.AdminNowRoleNodes(c, handler_name)
			if is_power == false {
				if c.Request.Method == "GET" {
					c.HTML(http.StatusOK, "layouts/no_power", gin.H{
						"Title": "Background Login",
						"Data":  "您没有管理该页面的权限，请勿非法进入！",
						"Route": handler_name,
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": config.HttpError,
						"info":   "您没有操作该功能的权限，请勿非法操作！",
					})
				}
				c.Abort()
			}
		}
		//endregion
		//fmt.Println("权限操作验证结束!")
		//region 判断数据库里面是有有这个权限节点，没有则需要进行添加
		//admin_navigation_node := new(models.AdminNavigationNode)
		//res, _ := databases.Orm.Where("route_action=?", handler_name).Exist(admin_navigation_node)
		//if res == false {
		//	if c.Request.Method == "GET" {
		//		c.HTML(http.StatusOK, "layouts/no_power", gin.H{
		//			"Title": "Background Login",
		//			"Data":  handler_name + "没有在数据库中存在，请添加后再操作",
		//		})
		//	} else {
		//		c.JSON(http.StatusOK, gin.H{
		//			"status": config.HttpError,
		//			"info":   handler_name + "没有在数据库中存在，请添加后再操作",
		//		})
		//	}
		//	c.Abort()
		//}
		//endregion
		//fmt.Println("权限标识验证结束!")
		//生成请求日志
		if background.RequestAllToLog(c) == false {
			c.Abort()
		}
		c.Next() //处理请求
	}
}
