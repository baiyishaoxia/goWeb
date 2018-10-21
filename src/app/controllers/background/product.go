package background

import (
	"app"
	"app/models/background"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//region Remark:品牌管理 Author:tang
func GetBrandList(c *gin.Context) {
	c.HTML(http.StatusOK, "product/brand_list", gin.H{
		"Title": "Background Index",
	})
}

//endregion
//region Remark:分类管理 Author:tang
func GetCategoryList(c *gin.Context) {
	var cates []models.Category
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	data, num, all, page := models.GetCategoryList(page-1, limit, keywords)
	for key, val := range data {
		data[key].Title = "<span class='folder-open'></span>" + val.Title
	}
	//格式化数据
	cates = models.UnlimitedForLevel(data, "<span class='folder-line'></span>", 0, 0)
	c.HTML(http.StatusOK, "product/category_list", gin.H{
		"Title":    "Background Index",
		"Data":     cates,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"SubString": func(str string) string {
			return app.SubString(str, 0, 10)
		},
	})
}

//添加分类
func GetCategoryCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "product/category_add", gin.H{
		"Title":    "Background Index",
		"Category": models.CategoryTreeData(),
	})
}
func PosCategoryCreate(c *gin.Context) {
	title := c.PostForm("title")
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	content := c.PostForm("content")
	if strings.Count(title, "")-1 > 30 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "分类名称最长为30个字!",
		})
		return
	}
	add := models.Category{Title: title, ParentId: parent_id, Content: content}
	res, err := databases.Orm.Insert(add)
	if err != nil || res < 1 {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "添加失败",
			"url":    "/admin/category/create",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "新增成功",
		"url":    "/admin/category/list",
	})
	return
}

//region Remark:修改 Author:tang
func GetCategoryEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := models.GetCategoryById(id)
	c.HTML(http.StatusOK, "product/category_edit", gin.H{
		"Title":    "Background Login",
		"Data":     data,
		"Category": models.CategoryTreeData(),
	})
}
func PostCategoryEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	title := c.PostForm("title")
	content := c.PostForm("content")
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	if strings.Count(title, "")-1 > 30 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "分类名称最长为30个字!",
		})
		return
	}
	category := models.Category{
		ParentId: parent_id,
		Title:    title,
		Content:  content,
		Sort:     sort,
	}
	_, err := databases.Orm.Cols("parent_id", "title", "content", "sort").Where("id=?", id).Update(category)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "修改失败",
			"url":    "/admin/category/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/category/list",
	})
	return
}

//endregion

//region Remark:删除 Author:tang
func PostCategoryDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	var category = new(models.Category)
	_, err := databases.Orm.In("id", ids).Delete(category)
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
		"url":    "/admin/category/list",
	})
	return
}

//endregion
//region Remark:设置分类状态 Author:tang
func GetCategorySetStatus(c *gin.Context) {
	id := c.Query("id")
	category := new(models.Category)
	databases.Orm.Id(id).Get(category)
	if category.Status == 1 {
		category.Status = 2 //未激活
	} else {
		category.Status = 1 //已激活
	}
	res, _ := databases.Orm.Id(id).Update(category)
	fmt.Println("--------", res)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/category/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/category/list",
		})
	}
}

//endregion
//region Remark:保存分类排序 Author:tang
func PostCategorySave(c *gin.Context) {
	ids := c.PostFormArray("data[sort][]")
	for _, v := range ids {
		id, _ := strconv.ParseInt(v, 10, 64)
		category := new(models.Category)
		sort, _ := strconv.ParseInt(c.PostForm("data["+v+"][sort]"), 10, 64)
		category.Sort = sort
		databases.Orm.Where("id = ?", id).Update(category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/category/list",
	})
}

//endregion
//endregion
//region Remark:产品管理 Author:tang
func GetProductList(c *gin.Context) {
	c.HTML(http.StatusOK, "product/product_list", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:产品管理--创建 Author:tang
func GetProductCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "product/product_add", gin.H{
		"Title": "Background Index",
	})
}

//endregion
//region Remark:产品管理--修改 Author:tang
func GetProductEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "product/product_edit", gin.H{
		"Title": "Background Index",
	})
}

//endregion
