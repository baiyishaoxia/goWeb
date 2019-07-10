package blog

import (
	"app/service/common"
	"app/service/home"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetMain(c *gin.Context) {
	user := common.ValidateLogin(c)
	if user.Id != 0 {
		fmt.Println("登录成功", user)
	}
	banner, _ := home.BannerList("pc_banner")
	c.HTML(http.StatusOK, "default/index", gin.H{
		"Title":  "欢迎使用GO语言编程",
		"Banner": banner,
		"User":   user,
	})
}

func FileDownload(c *gin.Context) {
	name := c.Query("path")
	data := strings.Split(name, "/")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", data[len(data)-1])) //对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Accept-Ranges", "bytes")
	c.File("./" + c.Query("path"))
}
