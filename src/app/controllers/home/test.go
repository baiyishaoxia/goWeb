package home

import (
	"app"
	"app/models/home"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//region Remark:数据库读取速度测试 Author:tang
func GetTest(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//每页1万条
	limit := 10000
	data, num, all, page := models.GetTestList(page-1, limit, keywords)
	c.HTML(http.StatusOK, "test/demo", gin.H{
		"Title":    "欢迎使用GO语言编程",
		"Host":     c.Request.Host,
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Substr": func(str string) string {
			return app.SubString(str, 0, 30)
		},
	})
}

//endregion
