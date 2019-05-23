package background

import (
	"app"
	"app/service/background"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"unicode/utf8"
)

//region Remark: 列表 Author      tang
func GetBannerCategoryList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	data, num, all, page := background.PagebannerCategoryList(limit, page-1, keywords)
	//模版
	c.HTML(http.StatusOK, "banner_category/list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"TimeInit": func(time app.Time) string {
			return time.String()
		},
	})
}

//endregion

//region Remark: 新增 Author    tang
func GetBannerCategoryCreate(c *gin.Context) {
	//模版
	c.HTML(http.StatusOK, "banner_category/create", gin.H{
		"Title": "Background Login",
	})
}
func PostBannerCategoryCreate(c *gin.Context) {
	title := c.PostForm("title")
	index := c.PostForm("index")
	intro := c.PostForm("intro")
	if utf8.RuneCountInString(intro) > 255 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "简介内容过长",
		})
		return
	}

	banner_category := new(background.BannerCategory)
	banner_category.Title = title
	banner_category.Index = index
	banner_category.Intro = intro
	has, err := databases.Orm.Insert(banner_category)
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
		"url":    "/admin/banner_category/list",
		"title":  "图片分类列表",
	})
	return
}

//endregion

//region Remark: 编辑 Author      tang
func GetBannerCategoryEdit(c *gin.Context) {
	cate := new(background.BannerCategory)
	databases.Orm.ID(c.Param("id")).Get(cate)
	//模版
	if cate.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "无该图片分类",
		})
		return
	}
	//模版
	c.HTML(http.StatusOK, "banner_category/edit", gin.H{
		"Title": "Background Login",
		"Data":  cate,
	})
}
func PostBannerCategoryEdit(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	index := c.PostForm("index")
	intro := c.PostForm("intro")
	if utf8.RuneCountInString(intro) > 255 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "简介内容过长",
		})
		return
	}

	item := new(background.BannerCategory)
	databases.Orm.ID(id).Get(item)
	if item.Title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "无该图片分类",
		})
		return
	}
	item.Title = title
	item.Index = index
	item.Intro = intro
	has, err := databases.Orm.Cols("title", "index", "intro").ID(id).Update(item)
	fmt.Println(err)
	if has < 1 || err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/banner_category/list",
		"title":  "图片分类列表",
	})
	return
}

//region Remark: 删除 Author      tang
func PostBannerCategoryDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	if background.HasBanners(ids) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "该分类下有图片, 请先删除该分类下的图片",
		})
		return
	}
	_, err := databases.Orm.In("id", ids).Delete(new(background.BannerCategory))
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
		"url":    "/admin/banner_category/list",
	})
	return
}

//endregion
