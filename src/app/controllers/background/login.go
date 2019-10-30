package background

import (
	"app"
	"app/models"
	"app/vendors/captcha/controllers"
	newredis "app/vendors/redis/models"
	session "app/vendors/session/models"
	"config"
	"databases"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title": "Background Login",
		"Host":  c.Request.Host,
	})
}

//POST 登陆处理方法
func PostLogin(c *gin.Context) {
	//验证码判断
	captcha := c.PostForm("captcha")
	if !controllers.VerifyCaptcha(c, captcha) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "验证码错误",
		})
		return
	}
	//查询数据库
	username := c.PostForm("username")
	password := app.Strmd5(c.PostForm("password"))
	admin := new(models.BlogAdmin)
	has, _ := databases.Orm.Where("username=?", username).Get(admin)

	//判断管理员是否锁定
	if admin.IsLock {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "账户已经被锁定",
		})
		return
	}
	//判断用户是否存在
	if has == false {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "用户名不存在",
		})
		return
	}
	//判断密码是否一致
	if password != admin.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "用户名密码错误",
		})
		return
	}

	//记录一次登陆次数
	admin.LoginCount = admin.LoginCount + 1

	var redis_key string = "admin:last_login_time:" + strconv.FormatInt(admin.Id, 10)
	//判断redis是否有缓存数据
	if res, _ := newredis.Exists(redis_key); res == true {
		fmt.Println("加载缓存")
		// json数据在go中是[]byte类型，所以此处用redis.Bytes转换
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		// [将json解析成map类型]
		lastLoginTime := make(map[string]interface{})
		json.Unmarshal(valueBytes, &lastLoginTime)
		fmt.Println("Unmarshal lastLoginTime", valueBytes, lastLoginTime)
	} else {
		fmt.Println("设置缓存")
		//将上次的登录时间存入redis ,保存时间 5分钟 [将数据编码成json字符串]
		value, _ := json.Marshal(admin.LastLogin.Format("2006-01-02 15:04:05"))
		newredis.Set(redis_key, value, 5*60)
	}
	//更新当前时间
	admin.LastLogin = time.Now()
	databases.Orm.ID(admin.Id).Update(admin)
	//写入session并进行页面的跳转
	session.SetSession(c, "admin_id", admin.Id)
	//缓存到redis 1小时
	newredis.Set("admin_id", admin.Id, 60*60)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "校验成功，登陆中...",
		"url":    "/admin/main",
	})
}
