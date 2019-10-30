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

//region Remark:评论列表 Author:tang
func GetCommentsList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(models.ReadConfig("sys.paginate"))
	data, num, all, page := models.GetMessageList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "comments/list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Type": func(m_cate int64) string {
			data := []string{1: "留言墙", 2: "文章评论"}
			return data[m_cate]
		},
	})
}

//endregion

//region Remark:设置资讯状态 Author:tang
func GetCommentsStatus(c *gin.Context) {
	id := c.Query("id")
	item := new(models.Message)
	databases.Orm.Id(id).Get(item)
	if item.IsShow == true {
		item.IsShow = false //取消显示
	} else {
		item.IsShow = true //显示
	}
	res, _ := databases.Orm.Id(id).Cols("is_show").Update(item)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/comments/list",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/comments/list",
		})
		return
	}
}

//endregion

//region   批量删除   Author:tang
func PostCommentDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	_, err := databases.Orm.In("id", ids).Delete(&models.Message{})
	fmt.Println(ids, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/comments/list/",
	})
	return
}

//endregion

//region Remark:意见反馈 Author:tang
func GetFeedbackList(c *gin.Context) {
	c.HTML(http.StatusOK, "comments/feedback_list", gin.H{
		"Title": "Background Index",
	})
}

//endregion
