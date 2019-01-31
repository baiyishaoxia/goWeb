package background

import (
	"app"
	"app/models/background"
	models2 "app/vendors/redis/models"
	newredis "app/vendors/redis/models"
	session "app/vendors/session/models"
	"app/vendors/windows/controller"
	"config"
	"databases"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"runtime"
	"statistical"
	"strconv"
	"time"
)

func GetIndex(c *gin.Context) {
	admin_info := models.GetAdminInfo(c)
	c.HTML(http.StatusOK, "index", gin.H{
		"Title":      "Background Index",
		"admin_info": template.HTML("【" + admin_info.Username + "】"),
		"AdminInfo":  admin_info,
	})
}
func GetCenter(c *gin.Context) {
	admin_info := models.GetAdminInfo(c)
	//最后一次登录时间
	key := "admin:last_login_time:" + strconv.FormatInt(int64(admin_info.Id), 10)
	last_login_time, _ := redis.String(models2.Get(key))
	//获取windows操作系统信息
	var info []string
	var liunx bool = false //是否在liunx下运行
	if liunx == false {
		info, _, _ = controller.WindowsInfo()
	} else {
		info = []string{"暂无", "LinuxOs", runtime.GOOS, "暂无", "暂无", "暂无", "暂无", "暂无", "暂无", time.Now().Format("2006-01-02 15:04:05")}
	}

	c.HTML(http.StatusOK, "center", gin.H{
		"Title":         "Background Center",
		"client_ip":     c.ClientIP(),
		"version":       gin.Version,
		"os":            runtime.GOOS + " " + runtime.GOARCH + "(" + strconv.Itoa(runtime.GOMAXPROCS(0)) + ")",
		"server_ip":     app.PulicIP(),
		"AdminInfo":     admin_info,
		"LastLoginTime": last_login_time,
		"time_now":      time.Now().Year(),
		"WindowsInfo": func(i int) string {
			return info[i]
		},
		"cpu_num": runtime.NumCPU(),
		"DataCount": func(key string, i int) int64 {
			return statistical.GetDataCountSum()[key][i]
		},
	})
}

//region Remark:退出 Author:tang
func GetExit(c *gin.Context) {
	session.DeleteSession(c, "admin_id")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "退出成功",
		"url":    "/login",
	})
}

//endregion

//region Remark:清除缓存 Author:tang
func GetClear(c *gin.Context) {
	models2.DelAll()
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "清空缓存成功",
		"url":    "/admin/main",
	})
}

//endregion

//region Remark:修改密码 Author:tang
func GetAdminPass(c *gin.Context) {
	//模版
	c.HTML(http.StatusOK, "admin/repass", gin.H{
		"Title": "Background Login",
	})
}
func PostAdminPass(c *gin.Context) {
	loginAdmin := models.GetAdminInfo(c)
	if app.Strmd5(c.PostForm("old_password")) != loginAdmin.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "原密码错误",
		})
		return
	}
	if c.PostForm("password") == "" || c.PostForm("password_confirmation") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码不能为空",
		})
		return
	}
	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码重复错误",
		})
		return
	}
	admin := new(models.BlogAdmin)
	admin.Password = app.Strmd5(c.PostForm("password"))
	_, err := databases.Orm.ID(loginAdmin.Id).Update(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	newredis.DelKeyByPrefix("admin:info:" + strconv.FormatInt(loginAdmin.Id, 10))
	newredis.DelKey("admin_id")
	session.DeleteSession(c, "admin_id")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功，请重新登陆",
		"url":    "/login",
	})
}

//endregion
