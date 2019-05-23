package blog

import (
	"app/service/home"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMain(c *gin.Context) {
	banner, _ := home.BannerList("pc_banner")
	c.HTML(http.StatusOK, "default/index", gin.H{
		"Title":  "欢迎使用GO语言编程",
		"Banner": banner,
	})
}
