package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPictureList(c *gin.Context) {
	c.HTML(http.StatusOK, "picture/list", gin.H{
		"Title": "Background Index",
	})
}
func GetPictureShow(c *gin.Context) {
	c.HTML(http.StatusOK, "picture/show", gin.H{
		"Title": "Background Index",
	})
}

//region Remark:新增图片 Author:tang
func GetPictureCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "picture/add", gin.H{
		"Title": "Background Index",
	})
}

//endregion
//region Remark:编辑 Author:tang
func GetPictureEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "picture/edit", gin.H{
		"Title": "Background Index",
	})
}

//endregion
