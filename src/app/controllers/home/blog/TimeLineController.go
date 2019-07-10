package blog

import (
	"app/service/background"
	"app/service/common"
	"config"
	"github.com/gin-gonic/gin"
	"net/http"
)

//点点滴滴
func GetBlogTimeLine(c *gin.Context) {
	c.HTML(http.StatusOK, "default/timeline", gin.H{
		"Title": "欢迎使用GO语言编程",
		"User":  common.ValidateLogin(c),
	})
}

//Ajax请求数据
func GetBlogTimeLineAjax(c *gin.Context) {
	data := background.LineToLine()
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"data":   data,
	})
	return
}
