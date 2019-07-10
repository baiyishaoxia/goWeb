package blog

import (
	"app/service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBlogSiguVideo(c *gin.Context) {
	c.HTML(http.StatusOK, "default/sigu_video", gin.H{
		"Title": "欢迎使用GO语言编程",
		"User":  common.ValidateLogin(c),
	})
}
