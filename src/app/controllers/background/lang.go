package background

import (
	"app/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//region Remark:列表 Author:tang
func GetLangIndex(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit := config.Limit
	nations_id := c.Query("nations_id")
	data, num, all, page := models.LangList(page-1, limit, keywords, nations_id)
	c.HTML(http.StatusOK, "lang/list", gin.H{
		"Title":     "语言管理",
		"Langs":     data,
		"Keywords":  keywords,
		"Num":       num,
		"DownPage":  float64(page + 1),
		"Page":      float64(page),
		"UpPage":    float64(page - 1),
		"All":       all,
		"Nations":   models.GetNations(),
		"NationsId": nations_id,
	})
}

//endregion

//region Remark:创建 Author:tang
func GetLangCreate(c *gin.Context) {
	//提取所有国家信息

	c.HTML(http.StatusOK, "lang/create", gin.H{
		"LangNations": models.GetLangsNations(),
		"Title":       "新增语言",
	})
}
func PostLangCreate(c *gin.Context) {
	nations_titles := c.PostFormArray("nations[]")
	leng := len(nations_titles)
	mark := c.PostForm("mark")
	remark := c.PostForm("remark")
	//获取国家列表 id字段
	nations := models.GetLangsNations()
	//开启事务
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var count int = 0
	for key, val := range *nations {
		id := val.Id
		title := nations_titles[key]
		add := &models.Langs{LangNationsId: id, Title: title, Mark: mark, Remark: remark, Sort: 99}
		res, err := db.Insert(add)
		if err != nil {
			fmt.Println(err.Error())
			db.Rollback()
			return
		}
		if res < 1 {
			fmt.Println(res)
			db.Rollback()
			return
		}
		count++
	}

	if count == leng {
		//事务提交
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "添加成功",
			"url":    "/admin/lang/list",
		})
	} else {
		//事务回滚
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "添加失败",
			"url":    "/admin/lang/create",
		})
	}
}

//endregion

//region Remark:修改 Author:tang
func GetLangEdit(c *gin.Context) {
	id := c.Param("id")
	langs := models.GetLangById(id)
	langandnations := models.GetLangAndNationViews(langs.Mark)

	c.HTML(http.StatusOK, "lang/edit", gin.H{
		"Title":   "修改语言",
		"Data":    langandnations,
		"Info":    models.GetLangById(id),
		"Nations": models.GetNations(),
	})
}
func PostLangEdit(c *gin.Context) {
	nations_titles := c.PostFormArray("nations[]")
	lang_nations_id := c.PostFormArray("id[]")
	mark := c.PostForm("mark")
	remark := c.PostForm("remark")
	length := len(nations_titles)
	langs := new(models.Langs)
	langs.Mark = mark
	langs.Remark = remark
	var count int = 0
	//开启事务
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < length; i++ {
		langs.Title = nations_titles[i]
		if models.GetLangInId(lang_nations_id[i], mark) {
			_, err := db.Cols("mark", "remark", "title").Where("mark=?", langs.Mark).Update(langs, models.Langs{LangNationsId: lang_nations_id[i]})
			if err != nil {
				fmt.Println(err.Error())
				db.Rollback()
				return
			}
			count++
		} else {
			println("国家的记录语言包中没有", lang_nations_id[i])
			langs.LangNationsId = lang_nations_id[i]
			_, err = db.Insert(langs)
			if err != nil {
				fmt.Println(err.Error())
				db.Rollback()
				return
			}
			count++
		}
	}
	if count == length {
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "修改成功",
			"url":    "/admin/lang/list",
		})
		return
	} else {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "修改失败",
			"url":    "/admin/lang/list",
		})
		return
	}
}

//endregion

//region Remark:删除 Author:tang
func PostLangDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	//单个id也删除同为 标识符 的 所有数据
	length := len(ids)
	var k bool
	k = true
	for i := 0; i < length; i++ {
		lang := models.GetLangById(ids[i])
		if lang != nil {
			langs := new(models.Langs)
			_, err := databases.Orm.Where("mark = ?", lang.Mark).Delete(langs)
			if err != nil {
				fmt.Println(err.Error())
				k = false
				break
			}
		}
	}
	if k == true {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "删除成功",
			"url":    "/admin/lang/list",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败",
			"url":    "/admin/lang/list",
		})
		return
	}
}

//endregion

//region Remark:前台国家语言包接口调用 Author:tang
func GetLangApi(c *gin.Context) {
	var nat_id = c.Query("nations_id")
	if nat_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"data":   "国家ID不能为空!",
			"status": 201,
		})
		return
	}
	data := models.GetNowNationsLangs(nat_id)
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"data":   "未找到该国家的相关语言包!",
			"status": 200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":   data,
			"status": 200,
		})
		return
	}
}

//endregion
