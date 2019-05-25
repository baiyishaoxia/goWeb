package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcshan/d3outh"
	"time"
)

//region   QQ互联登录   Author:tang
func GetBlogQQLogin(c *gin.Context) {
	qqconf := &d3auth.Auth_conf{Appid: "xxx", Appkey: "xxx", Rurl: "http://go.afurun.xyz"}
	qqouth := d3auth.NewAuth_qq(qqconf)
	fmt.Print(qqouth.Get_Rurl("state"))    //获取第三方登录地址
	token, err := qqouth.Get_Token("code") //回调页收的code 获取token
	fmt.Println("---------error---------", err)
	me, err := qqouth.Get_Me(token) //获取第三方id
	fmt.Println("---------token---------", token, "---------openid---------", err)
	time.Sleep(1 * time.Hour)
	userinfo, _ := qqouth.Get_User_Info(token, me.OpenID) //获取用户信息 userinfo 是一个json字符串返回
	fmt.Println("---------info---------", userinfo)
}

//endregion
