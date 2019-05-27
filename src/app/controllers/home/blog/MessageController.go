package blog

import (
	"app"
	models2 "app/models/background"
	"app/models/home"
	"app/service/background"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//region   获取留言数据   Author:tang
func GetBlogMessageAjax(c *gin.Context) {
	key, _ := strconv.Atoi(c.DefaultQuery("key", "1"))
	article_id, _ := strconv.ParseInt(c.DefaultQuery("article_id", "0"), 10, 64)
	wheres := map[string]interface{}{"key": key, "article_id": article_id} //综合条件
	flag := c.DefaultQuery("type", "")
	if flag == "index" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"data":   background.GetMessageNew(),
		})
		return
	}
	data := background.GetMessageListApi(wheres)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data":   data,
	})
	return
}

//endregion

//region   获取热评用户   Author:tang
func GetHotMessageUser(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	background.GetMessageHot(limit)
}

//endregion

//Ajax提交留言数据
func PostBlogMessageCreate(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)       //用户ID
	flag, _ := strconv.ParseInt(c.PostForm("type"), 10, 64)             //类型来自于(1留言墙，2文章评论)
	article_id, _ := strconv.ParseInt(c.PostForm("article_id"), 10, 64) //类型来自于(1留言墙，2文章评论)
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)   //父级ID
	content := app.RemoveHtmlScript(c.PostForm("content"))              //评论内容
	user := models.GetUserById(user_id)
	var add *background.Message
	if flag == 2 {
		//来自于文章的评论
		item := models.GetArticleById(article_id)
		if item.IsComment == false {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"data":   "博主对该文章暂未开放评论",
			})
			return
		}
		if item.StartTime > time.Now().Format("2006-01-02 15:04:05") {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"data":   "该文章评论于" + item.StartTime + "开放",
			})
			return
		}
		if time.Now().Format("2006-01-02 15:04:05") > item.EndTime {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"data":   "该文章评论已于" + item.EndTime + "结束",
			})
			return
		}
		//当前文章评论数+1
		res, err := databases.Orm.Id(item.Id).Incr("count_num").Update(&models2.Article{Id: item.Id})
		fmt.Println("--------", res, err)
		add = &background.Message{Content: content, UsersId: user_id, ParentId: parent_id, MessageCateId: flag, IsShow: true, ArticleId: article_id} //article_id
	} else {
		add = &background.Message{Content: content, UsersId: user_id, ParentId: parent_id, MessageCateId: flag, IsShow: true}
	}
	_, message := background.InsertMessage(add)
	c.JSON(http.StatusOK, gin.H{
		"status":  config.HttpSuccess,
		"data":    "提交成功",
		"user":    user,
		"content": content,
		"message": message,
	})
	return
}
