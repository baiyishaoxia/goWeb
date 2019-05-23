package background

import (
	"app/models/background"
	"app/service/background"
	"config"
	"databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//时光轴 Controller

// region Remark:列表 Author:tang
func GetTimeLineList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	data, num, all, page := background.GetTimeLineList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "time_line/list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"TimeInit": func(time time.Time) string {
			return time.Format("2006-01-02 15:04:05")
		},
	})
}

//endregion

//region Remark:新增 Author:tang
func GetTimeLineCreate(c *gin.Context) {
	admin_roles := make([]models.AdminRole, 0)
	databases.Orm.Find(&admin_roles)
	//模版
	c.HTML(http.StatusOK, "time_line/create", gin.H{
		"Title": "Background Login",
	})
}
func PostTimeLineCreate(c *gin.Context) {
	if c.PostForm("content") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "内容不能为空！",
		})
		return
	}
	loc, _ := time.LoadLocation("Local")
	line_time, err := time.ParseInLocation("2006-01-02 15:04:05", c.PostForm("start_time"), loc)
	line := new(background.TimeLine)
	line.Title = c.PostForm("title")
	line.Content = c.PostForm("content")
	line.Time = line_time
	line.Year, _ = strconv.ParseInt(line_time.Format("2006"), 10, 64)
	line.IsShow = true
	_, err = databases.Orm.Insert(line)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/time_line/list",
	})
	return
}

//endregion

//region Remark:删除 Author:tang
func PostTimeLineDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	var line = new(background.TimeLine)
	_, err := databases.Orm.In("id", ids).Delete(line)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/time_line/list",
	})
	return
}

//endregion

//region Remark:修改 Author:tang
func GetTimeLineEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := background.FindOneTimeLine("id", id)
	//模版
	c.HTML(http.StatusOK, "time_line/edit", gin.H{
		"Title": "Background Login",
		"Data":  data,
		"TimeInit": func(time time.Time) string {
			return time.Format("2006-01-02 15:04:05")
		},
	})
}
func PostTimeLineEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := background.FindOneTimeLine("id", id)
	data.Content = c.PostForm("content")
	loc, _ := time.LoadLocation("Local")
	line_time, err := time.ParseInLocation("2006-01-02 15:04:05", c.PostForm("start_time"), loc)
	data.Title = c.PostForm("title")
	data.Time = line_time
	data.Year, _ = strconv.ParseInt(line_time.Format("2006"), 10, 64)
	res, err := background.EditTimeLine("content,title,time", data)
	if res == false {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/time_line/list",
	})
	return
}

//endregion
