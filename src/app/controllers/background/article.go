package background

import (
	"app/models/background"
	models2 "app/vendors/redis/models"
	"config"
	"databases"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	data, num, all, page := models.GetArticleList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "article/list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Category": models.CategoryTreeData(),
	})
}

//region Remark:创建 Author:tang
func GetArticleCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "article/add", gin.H{
		"Title":         "Background Index",
		"Category":      models.CategoryTreeData(),
		"ArticleType":   models.GetArticleType(),
		"ArticleAuthor": models.GetArticleAuthor(),
	})
}
func PostArticleCreate(c *gin.Context) {
	title := c.PostForm("title")
	author_id, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	cate_id, _ := strconv.ParseInt(c.PostForm("cate_id"), 10, 64)
	art_type, _ := strconv.ParseInt(c.PostForm("type"), 10, 64)
	img := c.PostForm("img")
	intro := c.PostForm("intro")
	source := c.PostForm("source")
	keywords := c.PostForm("keywords")
	content := c.PostForm("content")
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	article := models.Article{
		Title: title, AuthorId: author_id, CateId: cate_id, Img: img, Content: template.HTML(content),
		Intro: intro, Keywords: keywords, StartTime: start_time, EndTime: end_time, Source: source, Type: int(art_type),
	}
	has, _ := databases.Orm.Insert(article)
	if has >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "添加成功",
			"url":    "/admin/article/list",
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
func GetArticleEdit(c *gin.Context) {
	id := c.Param("id")
	article := new(models.Article)
	_, err := databases.Orm.Table("article").Where("id = ?", id).Get(article)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.HTML(http.StatusOK, "article/edit", gin.H{
		"Title":         "Background Index",
		"Category":      models.CategoryTreeData(),
		"Data":          article,
		"ArticleType":   models.GetArticleType(),
		"ArticleAuthor": models.GetArticleAuthor(),
	})
}
func PostArticleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	title := c.PostForm("title")
	author_id, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	cate_id, _ := strconv.ParseInt(c.PostForm("cate_id"), 10, 64)
	art_type, _ := strconv.ParseInt(c.PostForm("type"), 10, 64)
	fmt.Println(author_id, cate_id, art_type)
	img := c.PostForm("img")
	intro := c.PostForm("intro")
	source := c.PostForm("source")
	keywords := c.PostForm("keywords")
	content := c.PostForm("content")
	start_time := c.PostForm("start_time")
	end_time := c.PostForm("end_time")
	article := models.Article{
		Title: title, AuthorId: author_id, CateId: cate_id, Img: img, Content: template.HTML(content),
		Intro: intro, Keywords: keywords, StartTime: start_time, EndTime: end_time, Source: source, Type: int(art_type),
	}
	has, _ := databases.Orm.Cols("title", "author_id", "cate_id", "type", "img", "intro", "source", "keywords", "content", "start_time", "end_time").Update(article, models.Article{Id: id})
	if has >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "编辑成功",
			"url":    "/admin/article/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"info":   "编辑失败",
	})
	return
}
func PostArticleDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	_, err := databases.Orm.In("id", ids).Delete(&models.Article{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "删除成功",
			"url":    "/admin/article/list/",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"info":   "删除失败",
	})
	return
}

//endregion

//region Remark:设置资讯状态 Author:tang
func GetArticleSetStatus(c *gin.Context) {
	id := c.Query("id")
	article := new(models.Article)
	databases.Orm.Id(id).Get(article)
	if article.Status == 1 {
		article.Status = 2 //下架
	} else {
		article.Status = 1 //发布
	}
	res, _ := databases.Orm.Id(id).Update(article)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/article/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/article/list",
		})
	}
}

//endregion

//region Remark:自定义搜索 Author:tang
type ArticleAllSort struct {
	article []models.Article
	by      func(p, q *models.Article) bool
}

func (pw ArticleAllSort) Len() int { // 重写 Len() 方法
	return len(pw.article)
}
func (pw ArticleAllSort) Swap(i, j int) { // 重写 Swap() 方法
	pw.article[i], pw.article[j] = pw.article[j], pw.article[i]
}
func (pw ArticleAllSort) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.article[i], &pw.article[j])
}
func GetSearchList(c *gin.Context) {
	keywords := c.Query("keywords")
	imgHost := c.Request.Host
	//[way 1升序 2降序  type_key根据某一字段排序]
	way := c.Query("way")
	type_key := c.Query("type_key")
	//搜索数据
	limit := int64(config.Limit)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
	if page == 0 {
		page = 1
	}
	list, page, all, num := models.SearchArticleBykeys(keywords, imgHost, limit, page, way, type_key)
	if way == "" && type_key == "" {
		//默认排序
		sort.Sort(ArticleAllSort{
			list, func(p, q *models.Article) bool {
				return q.CreatedAt.String() < p.CreatedAt.String()
			}})
	}
	//其它信息
	admin_info := models.GetAdminInfo(c)
	//最后一次登录时间
	key := "admin:last_login_time:" + strconv.FormatInt(int64(admin_info.Id), 10)
	last_login_time, _ := redis.String(models2.Get(key))
	c.HTML(http.StatusOK, "search", gin.H{
		"Title":         "Background Index",
		"client_ip":     c.ClientIP(),
		"AdminInfo":     admin_info,
		"LastLoginTime": last_login_time,
		"Article":       list,
		"Keywords":      keywords,
		"DownPage":      float64(page + 1),
		"Page":          float64(page),
		"UpPage":        float64(page - 1),
		"Num":           num,
		"All":           all,
		"AuthorById": func(id int) string {
			return models.GetArticleAuthorById(id)
		},
	})
}

//endregion
