package home

import (
	"app"
	"app/channel"
	"app/models"
	"config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//region   活动抽奖   Author:tang
func PostActiveSign(c *gin.Context) {
	var (
		token        = c.Request.Header.Get("token")
		user_id      = models.GetUserByToken(token).Id
		active_id, _ = strconv.ParseInt(c.PostForm("active_id"), 10, 64)
		ticket_id, _ = strconv.ParseInt(c.PostForm("ticket_id"), 10, 64)
		uuid         = app.Uuid()
	)
	//写入管道
	channel.UserAmountActiveOrderChan <- &channel.ActiveRewardStruck{
		UserId:   user_id,
		ActiveId: active_id,
		TicketId: ticket_id,
		Uuid:     uuid,
	}
	//获取结果
	msg := channel.HandleResult(uuid)
	if msg.Status == 1 {
		c.String(http.StatusOK, msg.Message)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"data":   msg.Message,
		})
		return
	}
}
