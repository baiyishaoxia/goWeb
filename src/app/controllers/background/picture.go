package background

import (
	"app"
	"app/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//图片列表
func GetPictureList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	if page == 0 {
		page = 1
	}
	cate_id, _ := strconv.ParseInt(c.PostForm("cate_id"), 10, 64)
	author_id, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	limit, _ := strconv.ParseInt(models.ReadConfig("sys.paginate"), 10, 64)
	data, num, all, page := models.GetPictureList(page-1, limit, keywords, cate_id, author_id)
	c.HTML(http.StatusOK, "picture/list", gin.H{
		"Title":      "Background Index",
		"All":        all,
		"UpPage":     float64(page - 1),
		"Page":       float64(page),
		"DownPage":   float64(page + 1),
		"Data":       data,
		"Num":        num,
		"Keywords":   keywords,
		"CategoryId": cate_id,
		"Category":   models.CategoryTreeData(),
		"AuthorName": func(author_id int) string {
			return models.GetArticleAuthorById(author_id)
		},
	})
}

//查看相册
func GetPictureShow(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, images := models.GetPictureInfo(id)
	c.HTML(http.StatusOK, "picture/show", gin.H{
		"Title": "Background Index",
		"Data":  images,
		"Count": len(images),
	})
}

//region Remark:新增图片 Author:tang
func GetPictureCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "picture/add", gin.H{
		"Title":         "Background Index",
		"Category":      models.CategoryTreeData(),
		"ArticleAuthor": models.GetArticleAuthor(),
	})
}
func PostPictureCreate(c *gin.Context) {
	title := c.PostForm("title")
	cate_id, _ := strconv.ParseInt(c.PostForm("cate_id"), 10, 64)
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	author_id, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	source := c.PostForm("source")
	keywords := c.PostForm("keywords")
	intro := c.PostForm("intro")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	img := c.PostForm("img")
	images := c.PostFormArray("images_list[]")
	images_init := ""
	for key, val := range images {
		if key != len(images)-1 {
			images_init += val + ","
		} else {
			images_init += val
		}
	}
	picture := models.Picture{
		Title: title, AuthorId: author_id, CateId: cate_id, Img: img, Images: images_init,
		Intro: intro, Keywords: keywords, StartTime: start_time, EndTime: end_time, Source: source, Sort: sort,
	}
	has, _ := databases.Orm.Insert(picture)
	if has >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "添加成功",
			"url":    "/admin/picture/list",
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
//region Remark:编辑 Author:tang
func GetPictureEdit(c *gin.Context) {
	id := c.Param("id")
	picture := new(models.Picture)
	has, err := databases.Orm.Where("id = ?", id).Get(picture)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !has {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "该相册不存在",
		})
		return
	}
	//格式化批量图片
	images := app.StrSplitArray(picture.Images)
	fmt.Println(images)
	c.HTML(http.StatusOK, "picture/edit", gin.H{
		"Title":         "Background Index",
		"Data":          picture,
		"Images":        images,
		"Category":      models.CategoryTreeData(),
		"ArticleAuthor": models.GetArticleAuthor(),
	})
}
func PostPictureEdit(c *gin.Context) {
	images := c.PostFormArray("images_list[]")
	images_init := ""
	for key, val := range images {
		if key != len(images)-1 {
			images_init += val + ","
		} else {
			images_init += val
		}
	}
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	title := c.PostForm("title")
	cate_id, _ := strconv.ParseInt(c.PostForm("cate_id"), 10, 64)
	author_id, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	source := c.PostForm("source")
	keywords := c.PostForm("keywords")
	intro := c.PostForm("intro")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	is_comment := c.PostForm("is_comment")
	var comment bool
	if is_comment == "on" {
		comment = true
	} else {
		comment = false
	}
	img := c.PostForm("img")
	if strings.Count(title, "")-1 > 30 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "标题最多30个字符",
		})
		return
	}
	fmt.Println()
	picture := new(models.Picture)
	picture.Title = title
	picture.CateId = cate_id
	picture.AuthorId = author_id
	picture.Images = images_init
	picture.Source = source
	picture.Keywords = keywords
	picture.Intro = intro
	picture.Sort = sort
	picture.IsComment = comment
	picture.Img = img
	picture.StartTime = start_time
	picture.EndTime = end_time
	res, err := databases.Orm.Cols("title", "cate_id", "author_id", "source", "keywords", "images", "intro", "sort", "is_comment", "img", "start_time", "end_time").Update(picture, models.Picture{Id: id})
	if err != nil || res < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "修改失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "修改成功",
			"url":    "/admin/picture/list",
		})
		return
	}
}

//endregion

//region Remark:列表删除 Author:tang
func PostPictureDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	var data = new(models.Picture)
	_, err := databases.Orm.In("id", ids).Delete(data)
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
		"url":    "/admin/picture/list",
	})
	return
}

//endregion

//region Remark:详情删除 Author:tang
func PostPictureShowDel(c *gin.Context) {
	ids := c.PostFormArray("ids[]")
	images := c.PostFormArray("images[]")
	var picture = new(models.Picture)
	var count int = 0
	var new_img string = ""
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for key, val := range ids {
		db.Id(val).Get(picture)
		old_img := picture.Images
		//处理字段中的最后一个图片
		if (strings.Index(old_img, images[key]) != -1) && (strings.Index(old_img, images[key]+",") == -1) {
			//包括图片并删除前面的逗号
			new_img = strings.Replace(","+old_img, images[key], "", -1)
		} else {
			//包括图片并删除后面的逗号
			new_img = strings.Replace(old_img, images[key]+",", "", -1)
		}
		picture.Images = new_img
		res, _ := db.Id(val).Cols("images").Update(picture)
		if res > 0 {
			count++
		}
	}
	if count == len(ids) {
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "删除成功",
			"url":    "/admin/picture/list",
		})
		return
	} else {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败",
			"url":    "/admin/picture/list",
		})
		return
	}
}

//endregion

//region Remark:设置相册状态 Author:tang
func GetPictureSetStatus(c *gin.Context) {
	id := c.Query("id")
	picture := new(models.Picture)
	databases.Orm.Id(id).Get(picture)
	if picture.Status == 1 {
		picture.Status = 2 //下架
	} else {
		picture.Status = 1 //发布
	}
	res, _ := databases.Orm.Id(id).Cols("status").Update(picture)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/picture/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/picture/list",
		})
	}
}

//endregion
