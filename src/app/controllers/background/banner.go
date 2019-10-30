package background

import (
	"app"
	"app/models"
	"app/service/background"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

//region Remark: 列表 Author      tang
func GetBannerList(c *gin.Context) {
	banner_category_id, _ := strconv.ParseInt(c.Query("banner_category_id"), 10, 64)
	keywords := c.DefaultQuery("keywords", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	data, num, all, page := background.PageBannerList(keywords, banner_category_id, limit, page-1)
	if app.IsAjax(c) {
		data := background.PageBannerListAjax(keywords, banner_category_id, limit, page-1)
		c.String(http.StatusOK, data)
		return
	}
	cate := new([]models.BannerCategory)
	databases.Orm.Find(cate)
	//模版
	c.HTML(http.StatusOK, "banner/list", gin.H{
		"Title":            "Background Login",
		"BannerCategoryId": banner_category_id,
		"Category":         cate,
		"Data":             data,
		"Keywords":         keywords,
		"Num":              num,
		"DownPage":         float64(page + 1),
		"Page":             float64(page),
		"UpPage":           float64(page - 1),
		"All":              all,
		"Html": func(html string) template.HTML {
			return template.HTML(html)
		},
		"TimeInit": func(time app.Time) string {
			return time.String()
		},
	})
}

//endregion

//region Remark: 新增 Author      tang
func GetBannerCreate(c *gin.Context) {
	banner_category := new([]models.BannerCategory)
	databases.Orm.Find(banner_category)
	//模版
	c.HTML(http.StatusOK, "banner/create", gin.H{
		"Title":    "Background Login",
		"Category": banner_category,
	})
}
func PostBannerCreate(c *gin.Context) {
	title := c.PostForm("title")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	banner_category_id, _ := strconv.ParseInt(c.PostForm("banner_category_id"), 10, 64)
	url := c.PostForm("url")
	intro := c.PostForm("intro")
	abstract := c.PostForm("abstract")
	image := c.PostForm("image")
	content := c.PostForm("content")
	add := &models.Banner{Title: title, BannerCategoryId: banner_category_id, Sort: sort, Url: url, Intro: intro, Image: image, Content: content, Abstract: abstract}
	has, err := databases.Orm.Insert(add)
	if has < 1 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "新增失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "新增成功",
		"url":    "/admin/banner/list",
		"title":  "图片列表",
	})
	return
}

//endregion

//region Remark: 编辑 Author      tang
func GetBannerEdit(c *gin.Context) {
	id := c.Param("id")
	data := new(models.Banner)
	databases.Orm.Where("id=?", id).Get(data)
	if data.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "无该图片",
		})
		return
	}
	banner_category := new([]models.BannerCategory)
	databases.Orm.Find(banner_category)
	//模版
	c.HTML(http.StatusOK, "banner/edit", gin.H{
		"Title":    "Background Login",
		"Data":     data,
		"Category": banner_category,
	})
}

func PostBannerEdit(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	banner_category_id, _ := strconv.ParseInt(c.PostForm("banner_category_id"), 10, 64)
	url := c.PostForm("url")
	intro := c.PostForm("intro")
	abstract := c.PostForm("abstract")
	image := c.PostForm("image")
	content := c.PostForm("content")
	item := new(models.Banner)
	databases.Orm.Where("id=?", id).Get(item)
	if item.Title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "无标题",
		})
		return
	}
	item.Title = title
	item.Sort = sort
	item.Url = url
	item.BannerCategoryId = banner_category_id
	item.Intro = intro
	item.Abstract = abstract
	item.Image = image
	item.Content = content
	has, err := databases.Orm.Cols("title", "sort", "url", "banner_category_id", "image", "intro", "abstract", "content").ID(id).Update(item)
	fmt.Println(has, err)
	if has < 1 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/banner/list",
		"title":  "图片列表",
	})
	return
}

//endregion
//region Remark: 删除 Author; chijian
func PostBannerDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	has, err := databases.Orm.In("id", ids).Delete(new(models.Banner))
	if err != nil {
		fmt.Println(err.Error())
	}
	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/banner/list",
		"title":  "图片管理",
	})
	return
}

//endregion

//region   保存数据
func PostBannerSave(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	banner := new(models.Banner)
	data := databases.Orm.ID(id)
	banner.Sort, _ = strconv.ParseInt(c.PostForm("value"), 10, 64)
	data.Update(banner)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/banner/list",
	})
	return
}

//endregion
