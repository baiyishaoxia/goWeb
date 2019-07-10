package blog

import (
	"app/models/background"
	session "app/vendors/session/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcshan/d3outh"
	"net/http"
)

var (
	AppId  = models.ReadConfig("qq_appid")  //appid
	AppKey = models.ReadConfig("qq_appkey") //appkey
	RUrl   = "http://go.afurun.xyz/article" //回调地址
)

//region   QQ互联登录   Author:tang
func GetBlogQQLogin(c *gin.Context) {
	qqconf := &d3auth.Auth_conf{Appid: AppId, Appkey: AppKey, Rurl: RUrl}
	qqouth := d3auth.NewAuth_qq(qqconf)
	fmt.Println(qqouth.Get_Rurl("state")) //获取第三方登录地址
	c.Redirect(http.StatusMovedPermanently, qqouth.Get_Rurl("state"))
}

//endregion

//region   退出   Author:tang
func GetBlogExit(c *gin.Context) {
	session.DeleteSession(c, "userid")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"info":   "退出成功",
		"url":    "/",
	})
}

//endregion
