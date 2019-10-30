package background

import (
	"app"
	"app/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
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
func PostCategoryCreate(c *gin.Context) {
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

//region Remark:产品管理 Author:tang
func GetProductList(c *gin.Context) {
	keywords := c.Query("keywords")
	cate_id, _ := strconv.ParseInt(c.Query("category_id"), 10, 64)
	start_time := c.Query("start_time")
	end_time := c.Query("end_time")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	data, num, all, page := models.GetProductList(page-1, limit, keywords, cate_id, start_time, end_time)
	c.HTML(http.StatusOK, "product/product_list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Category": models.CategoryTreeData(),
		"CateId":   cate_id,
		"Unit":     models.GetUnit(),
		"UnitTitle": func(unit int) string {
			return models.GetUnit()[unit]
		},
	})
}
func GetProductZnodes(c *gin.Context) {
	data := models.GetCategoryDictory()
	//value, _ := json.Marshal(data)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"info": data,
		},
	})
	return
}

//endregion

//region Remark:产品管理--创建 Author:tang
func GetProductCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "product/product_add", gin.H{
		"Title":    "Background Index",
		"Category": models.CategoryTreeData(),
		"Unit":     models.GetUnit(),
	})
}
func PostProductCreate(c *gin.Context) {
	title := c.PostForm("title")
	cate_id, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	long, _ := strconv.ParseInt(c.PostForm("long"), 10, 64)
	wide, _ := strconv.ParseInt(c.PostForm("wide"), 10, 64)
	high, _ := strconv.ParseInt(c.PostForm("high"), 10, 64)
	unit, _ := strconv.ParseInt(c.PostForm("unit"), 10, 64)
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	cost, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	low_price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	image := c.PostForm("image")
	intro := c.PostForm("intro")
	address := c.PostForm("address")
	material := c.PostForm("material")
	supplier := c.PostForm("supplier")
	weight := c.PostForm("weight")
	tags := c.PostForm("tags")
	content := c.PostForm("content")
	loc, _ := time.LoadLocation("Local")
	start_time, _ := time.ParseInLocation("2006-01-02", c.PostForm("start_at"), loc)
	end_time, _ := time.ParseInLocation("2006-01-02", c.PostForm("end_at"), loc)
	product := models.Product{
		Title: title, CategoryId: cate_id, Image: image, Content: template.HTML(content),
		Intro: intro, StartTime: start_time, EndTime: end_time,
		Sort: sort, Long: long, Wide: wide, High: high, Unit: unit, Price: price, Cost: cost, LowPrice: low_price, Address: address,
		Material: material, Supplier: supplier, Weight: weight, Tags: tags,
	}
	has, err := databases.Orm.Insert(product)
	fmt.Println(err)
	if has >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "添加成功",
			"url":    "/admin/product/list",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "添加失败",
		})
		return
	}
}

//endregion
//region Remark:产品管理--修改 Author:tang
func GetProductEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data := models.GetProductById(id)
	c.HTML(http.StatusOK, "product/product_edit", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Category": models.CategoryTreeData(),
		"Unit":     models.GetUnit(),
		"TimeInit": func(time time.Time) string {
			fmt.Println(time)
			return time.Format("2006-01-02 15:04:05")
		},
	})
}
func PostProductEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	title := c.PostForm("title")
	cate_id, _ := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	long, _ := strconv.ParseInt(c.PostForm("long"), 10, 64)
	wide, _ := strconv.ParseInt(c.PostForm("wide"), 10, 64)
	high, _ := strconv.ParseInt(c.PostForm("high"), 10, 64)
	unit, _ := strconv.ParseInt(c.PostForm("unit"), 10, 64)
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	cost, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	low_price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	image := c.PostForm("image")
	intro := c.PostForm("intro")
	address := c.PostForm("address")
	material := c.PostForm("material")
	supplier := c.PostForm("supplier")
	weight := c.PostForm("weight")
	tags := c.PostForm("tags")
	content := c.PostForm("content")
	loc, _ := time.LoadLocation("Local")
	start_time, _ := time.ParseInLocation("2006-01-02", c.PostForm("start_at"), loc)
	end_time, _ := time.ParseInLocation("2006-01-02", c.PostForm("end_at"), loc)
	product := &models.Product{
		Title: title, CategoryId: cate_id, Image: image, Content: template.HTML(content),
		Intro: intro, Tags: tags, StartTime: start_time, EndTime: end_time, Sort: sort,
		Long: long, Wide: wide, High: high, Unit: unit, Price: price, Cost: cost, LowPrice: low_price,
		Address: address, Material: material, Supplier: supplier, Weight: weight,
	}
	has, _ := databases.Orm.Update(product, models.Product{Id: id})
	if has >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "编辑成功",
			"url":    "/admin/product/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"info":   "编辑失败",
	})
	return
}

//endregion
//region   删除产品   Author:tang
func PostProductDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	_, err := databases.Orm.In("id", ids).Delete(&models.Product{})
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
		"url":    "/admin/product/list",
	})
	return
}

//endregion
