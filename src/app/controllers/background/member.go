package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//region Remark:会员管理 Author:tang
func GetMemberList(c *gin.Context) {
	c.HTML(http.StatusOK, "member/list", gin.H{
		"Title": "Background Index",
	})
}
func GetMemberCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "member/add", gin.H{
		"Title": "Background Index",
	})
}
func GetMemberEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "member/edit", gin.H{
		"Title": "Background Index",
	})
}
func GetMemberDelList(c *gin.Context) {
	c.HTML(http.StatusOK, "member/del", gin.H{
		"Title": "Background Index",
	})
}

//endregion
//region Remark:修改密码 Author:tang
func GetMemberPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "member/change_password", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:查看详情 Author:tang
func GetMemberShow(c *gin.Context) {
	c.HTML(http.StatusOK, "member/show", gin.H{
		"Title": "Background Index",
	})
}

//endregion
