package background

import (
	"app/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//region Remark:系统设置 Author:tang
func GetSystemBase(c *gin.Context) {
	c.HTML(http.StatusOK, "system/base", gin.H{
		"title":          "Background Login",
		"sys_name":       models.ReadConfig("sys.name"),
		"sys_keywords":   models.ReadConfig("sys.keywords"),
		"sys_abstract":   models.ReadConfig("sys.abstract"),
		"sys_cache_time": models.ReadConfig("sys.cache_time"),
		"sys_paginate":   models.ReadConfig("sys.paginate"),

		"static_style":     models.ReadConfig("static_style"),
		"upload_directory": models.ReadConfig("upload_directory"),
		"bottom_copyright": models.ReadConfig("bottom_copyright"),
		"for_the_record":   models.ReadConfig("for_the_record"),
		"statistical_code": models.ReadConfig("statistical_code"),

		"ip":               models.ReadConfig("ip"),
		"login_file_count": models.ReadConfig("login_file_count"),

		"email_model":   models.ReadConfig("email_model"),
		"smtp":          models.ReadConfig("smtp"),
		"port":          models.ReadConfig("port"),
		"email_name":    models.ReadConfig("email_name"),
		"email_address": models.ReadConfig("email_address"),

		"qq_appid":  models.ReadConfig("qq_appid"),
		"qq_appkey": models.ReadConfig("qq_appkey"),
	})
}

//region Remark:编辑 Author:tang
func PostSystemBase(c *gin.Context) {
	//定义数组结构体
	data := []models.Config{
		models.Config{
			Name:  "sys.name",
			Value: c.DefaultPostForm("sys.name", ""),
		},
		models.Config{
			Name:  "sys.keywords",
			Value: c.DefaultPostForm("sys.keywords", ""),
		},
		models.Config{
			Name:  "sys.abstract",
			Value: c.DefaultPostForm("sys.abstract", ""),
		},
		models.Config{
			Name:  "sys.cache_time",
			Value: c.DefaultPostForm("sys.cache_time", ""),
		},
		models.Config{
			Name:  "sys.paginate",
			Value: c.DefaultPostForm("sys.paginate", ""),
		},
		models.Config{
			Name:  "ip",
			Value: c.DefaultPostForm("ip", ""),
		},
		models.Config{
			Name:  "qq_appid",
			Value: c.DefaultPostForm("qq_appid", ""),
		},
		models.Config{
			Name:  "qq_appkey",
			Value: c.DefaultPostForm("qq_appkey", ""),
		},
	}
	//保存修改
	models.SetConfig(data)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/system/base",
	})
}

//endregion
//endregion

func GetSystemCategory(c *gin.Context) {
	c.HTML(http.StatusOK, "system/category", gin.H{
		"Title": "Background Index",
	})
}
func GetSystemCategoryCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "system/category_add", gin.H{
		"Title": "Background Index",
	})
}
func GetSystemCategoryEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "system/category_edit", gin.H{
		"Title": "Background Index",
	})
}
func GetSystemData(c *gin.Context) {
	c.HTML(http.StatusOK, "system/data", gin.H{
		"Title": "Background Index",
	})
}
func GetSystemShield(c *gin.Context) {
	c.HTML(http.StatusOK, "system/shielding", gin.H{
		"Title": "Background Index",
	})
}

//日志列表
func GetSystemLog(c *gin.Context) {
	keywords := c.Query("keywords")
	limit, _ := strconv.Atoi(models.ReadConfig("sys.paginate"))
	page, _ := strconv.Atoi(c.Query("page"))
	data, num, all, page := models.GetAdminLogList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "system/log", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Num":      num,
		"Keywords": keywords,
		"UpPage":   float64(page - 1),
		"Page":     float64(page),
		"DownPage": float64(page + 1),
		"All":      all,
		"TimeInit": func(time time.Time) string {
			return time.Format("2006-01-02 15:04:05")
		},
	})
}

//endregion

//region Remark:查看日志 Author:tang
func GetSystemLogShow(c *gin.Context) {
	id := c.Query("id")
	log := new(models.AdminLog)
	has, err := databases.Orm.Where("id", id).Get(log)
	if has {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"info":   "哈哈",
		})
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

//endregion
