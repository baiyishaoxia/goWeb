package blog

import (
	"app/service/common"
	"app/service/home"
	"config"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//关于本站
func GetBlogAbout(c *gin.Context) {
	data1, _ := home.BannerList("about_blog")
	data2, _ := home.BannerList("about_author")
	data3, _ := home.BannerList("about_friendship")
	data4, _ := home.BannerList("about_wall")
	c.HTML(http.StatusOK, "default/about", gin.H{
		"Title": "欢迎使用GO语言编程",
		"User":  common.ValidateLogin(c),
		"Data1": *data1[0],
		"Data2": *data2[0],
		"Data3": *data3[0],
		"Data4": *data4[0],
		"Html": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
}

//Ajax请求数据
func PostBannerList(c *gin.Context) {
	index := c.PostForm("index")
	banner, info := home.BannerList(index)
	if banner == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"data": gin.H{
				"info": info,
			},
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"data": gin.H{
				"info": banner,
			},
		})
		return
	}
}
