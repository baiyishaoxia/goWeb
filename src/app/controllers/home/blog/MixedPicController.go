package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//杂七杂八
func GetBlogMixedPic(c *gin.Context) {
	c.HTML(http.StatusOK, "default/mixed_pic", gin.H{
		"Title": "欢迎使用GO语言编程",
	})
}
