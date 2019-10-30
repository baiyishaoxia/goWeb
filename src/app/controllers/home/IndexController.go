package home

import (
	"app"
	"app/models"
	models2 "app/vendors/session/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": config.VueSuccess,
		"msg":    "欢迎学习GO语言!",
		"data":   nil,
		"num":    0,
	})
	return
}

func GetIndex(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	links, num, all, page := models.GetLinksList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "index/index", gin.H{
		"Title":    "欢迎使用GO语言编程",
		"Host":     c.Request.Host,
		"Data":     links,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Substr": func(str string) string {
			return app.SubString(str, 0, 30)
		},
	})
}

//region Remark:增 Author:tang
func GetIndexCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "index/create", gin.H{
		"Title":   "增",
		"captcha": models2.GetSession(c, "captcha_value"),
	})
}
func PostIndexCreate(c *gin.Context) {
	name := c.PostForm("name")
	title := c.PostForm("title")
	url := c.PostForm("url")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请填写url",
		})
		return
	}
	add := models.BlogLinks{LinkName: name, LinkTitle: title, LinkUrl: url, LinkOrder: sort}
	res, err := databases.Orm.Insert(add)
	if err != nil || res < 1 {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "添加失败",
			"url":    "/home/index/list/create",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "新增成功",
		"url":    "/home/index/list",
	})
	return
}

//endregion

//region Remark:改 Author:tang
func GetIndexEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := models.GetLinksById(id)
	c.HTML(http.StatusOK, "index/edit", gin.H{
		"Title": "改",
		"Data":  data,
	})
}
func PostIndexEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	name := c.PostForm("name")
	title := c.PostForm("title")
	url := c.PostForm("url")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请填写url",
		})
		return
	}
	var links = new(models.BlogLinks)
	links.LinkName = name
	links.LinkTitle = title
	links.LinkUrl = url
	links.LinkOrder = sort
	fmt.Println(links)
	_, err := databases.Orm.Cols("link_name", "link_title", "link_url", "link_order").Where("link_id=?", id).Update(links)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "修改失败",
			"url":    "/home/index/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/home/index/list",
	})
	return
}

//endregion

//region Remark:删除 Author:tang
func PostLinksDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	var links = new(models.BlogLinks)
	_, err := databases.Orm.In("link_id", ids).Delete(links)
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
		"url":    "/home/index/list",
	})
	return
}

//endregion

//region Remark:排序 Author:tang
func PostLinksSort(c *gin.Context) {
	ids := c.PostFormArray("data[sort][]")
	for _, v := range ids {
		id, _ := strconv.ParseInt(v, 10, 64)
		links := new(models.BlogLinks)
		sort, _ := strconv.ParseInt(c.PostForm("data["+v+"][sort]"), 10, 64)
		links.LinkOrder = sort
		databases.Orm.ID(id).Update(links)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/home/index/list",
	})
}

//endregion
