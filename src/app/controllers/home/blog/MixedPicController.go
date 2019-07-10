package blog

import (
	"app/service/common"
	"app/service/home"
	"config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//杂七杂八
func GetBlogMixedPic(c *gin.Context) {
	user := common.ValidateLogin(c)
	c.HTML(http.StatusOK, "default/mixed_pic", gin.H{
		"Title": "欢迎使用GO语言编程",
		"User":  user,
	})
}

func GetBlogMixedPicDetail(c *gin.Context) {
	user := common.ValidateLogin(c)
	id := c.Param("id")
	c.HTML(http.StatusOK, "default/mixed_pic_detail", gin.H{
		"Title": "欢迎使用GO语言编程",
		"Id":    id,
		"User":  user,
	})
}

//Ajax请求列表相册
func PostMixedPic(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultPostForm("limit", "10"))
	wheres := map[string]interface{}{
		"page":  page,
		"limit": limit,
	}
	data, num, all, page := home.GetPictureList(wheres)
	item := map[string]interface{}{}
	item["data"] = data
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"item": item,
			"num":  num,
			"all":  all,
			"page": page,
		},
	})
	return
}

//Ajax请求列表相册详情
func PostMixedPicDetailAjax(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultPostForm("limit", "10"))
	id, _ := strconv.Atoi(c.DefaultPostForm("id", "0"))
	wheres := map[string]interface{}{
		"page":  page,  //默认第1页
		"limit": limit, //每页条数
		"id":    id,    //当前相册ID
	}
	data, line, num, all, page := home.GetPictureDetail(wheres)
	if page > 1 && line["start"] == 0 {
		data = []map[string]interface{}{}
	}
	item := map[string]interface{}{}
	item["data"] = data
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data": gin.H{
			"item": item,
			"num":  num,
			"all":  all,
			"page": page,
		},
	})
	return
}
