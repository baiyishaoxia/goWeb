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
	var d1,d2,d3,d4 interface{}
	if len(data1)>0{
		d1 = data1[0]
	}
	if len(data2)>0{
		d2 = data2[0]
	}
	if len(data3)>0{
		d3 = data3[0]
	}
	if len(data4)>0{
		d4 = data4[0]
	}
	c.HTML(http.StatusOK, "default/about", gin.H{
		"Title": "欢迎使用GO语言编程",
		"User":  common.ValidateLogin(c),
		"Data1": d1,
		"Data2": d2,
		"Data3": d3,
		"Data4": d4,
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
