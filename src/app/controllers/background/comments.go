package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//region Remark:评论列表 Author:tang
func GetCommentsList(c *gin.Context) {

}

//endregion
//region Remark:意见反馈 Author:tang
func GetFeedbackList(c *gin.Context) {
	c.HTML(http.StatusOK, "comments/feedback_list", gin.H{
		"Title": "Background Index",
	})
}

//endregion
