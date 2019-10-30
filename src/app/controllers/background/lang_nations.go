package background

import (
	"app"
	"app/models"
	"config"
	"databases"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

//region Remark:列表 Author:tang
func GetLangNationsIndex(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//条数
	limit := config.Limit
	//总数
	num := models.LangNationsNum(keywords)
	//页数
	all := math.Ceil(float64(num) / float64(limit))
	if page < 1 {
		page = 1
	}
	if page > int(all) {
		page = int(all)
	}
	//渲染
	c.HTML(http.StatusOK, "lang_nations/list", gin.H{
		"Title":       "国家",
		"LangNations": models.LangNationsList(page-1, limit, keywords),
		"Keywords":    keywords,
		"Num":         num,
		"DownPage":    float64(page + 1),
		"Page":        float64(page),
		"UpPage":      float64(page - 1),
		"All":         all,
	})
}

//endregion

//region Remark:创建 Author:tang
func GetLangNationsCreate(c *gin.Context) {
	//模版c
	c.HTML(http.StatusOK, "lang_nations/create", gin.H{
		"Title": "国家",
	})

}

func PostLangNationsCreate(c *gin.Context) {
	title := c.PostForm("title")
	img_url := c.PostForm("img_url")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_open := c.PostForm("is_open")
	is_default := c.PostForm("is_default")

	var isOpen bool
	if is_open == "1" {
		isOpen = true
	} else {
		isOpen = false
	}
	var isDefault bool
	if is_default == "1" {
		isDefault = true
	} else {
		isDefault = false
	}

	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请输入国家名称",
		})
		return
	}
	if img_url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请上传国家图片",
		})
		return
	}

	add := &models.LangNations{Id: app.GetRandomSalt(32), Title: title, ImageUrl: img_url, Sort: sort, IsOpen: isOpen, IsDefault: isDefault}
	if add.AddNations() {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "保存成功",
			"url":    "/admin/nations/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "保存失败",
			"url":    "/admin/nations/create",
		})
	}
}

//endregion

//region Remark:修改 Author:tang
func GetLangNationsEdit(c *gin.Context) {
	id := c.Param("id")
	nations := models.GetNationsById(id)
	c.HTML(http.StatusOK, "lang_nations/edit", gin.H{
		"Title": "修改国家",
		"Data":  nations,
	})
}
func PostLangNationsEdit(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	img_url := c.PostForm("img_url")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_open := c.PostForm("is_open")
	is_default := c.PostForm("is_default")

	var isOpen, isDefault bool
	if is_open == "1" {
		//查询该国家下是否存在语言包
		println("-----国家启用------")
		mark := models.GetNationsInLangs(id)
		if mark != "" {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"info":   "请在语言管理中添加标识符" + mark + "的值",
			})
			return
		}
		isOpen = true

	} else {
		isOpen = false
		if models.GetNationsIsOpen() == 1 {
			if models.GetOneNationsIsOpen(id) == true {
				c.JSON(http.StatusOK, gin.H{
					"status": config.HttpError,
					"info":   "至少要启用一个国家语言",
				})
				return
			}
		}
	}
	if is_default == "1" {
		isDefault = true
		println(models.GetOneNationsIsDefault())
	} else {
		isDefault = false
		if models.GetNationsIsDefault() == 1 {
			if models.GetDefaultNations(id) == true {
				c.JSON(http.StatusOK, gin.H{
					"status": config.HttpError,
					"info":   "必须存在一个默认国家的语言",
				})
				return
			}
		}
	}
	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请输入国家名称",
		})
		return
	}
	if img_url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请上传国家图片",
		})
		return
	}

	nations := new(models.LangNations)
	nations.Title = title
	nations.ImageUrl = img_url
	nations.Sort = sort
	nations.IsOpen = isOpen
	nations.IsDefault = isDefault
	_, err := databases.Orm.Cols("image_url", "title", "sort", "is_default", "is_open").Update(nations, models.LangNations{Id: id})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "修改失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/nations/list",
	})
}

//endregion

//region Remark:删除 Author:tang
func PostLangNationsDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	nations := new(models.LangNations)
	_, err := databases.Orm.In("id", ids).Delete(nations)
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
		"url":    "/admin/nations/list",
	})
}

//endregion

//region Remark:保存 Author:tang
func PostLangNationsSave(c *gin.Context) {
	ids := c.PostFormArray("data[sort][]")
	for _, v := range ids {
		nations := new(models.LangNations)
		sort, _ := strconv.ParseInt(c.PostForm("data["+v+"][sort]"), 10, 64)
		nations.Sort = sort
		databases.Orm.Where("id = ?", v).Update(nations)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/nations/list",
	})
}

//endregion
