package controllers

import (
	captcha "app/vendors/captcha/models"
	redisModel "app/vendors/redis/models"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAppCaptcha(c *gin.Context) {
	width, _ := strconv.Atoi(c.Query("width"))
	height, _ := strconv.Atoi(c.Query("height"))
	d := make([]byte, 4)
	s := captcha.NewLen(4)
	char := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		char += strconv.FormatInt(int64(d[v]), 32)
	}
	redisModel.Set("app_captcha_value", char, 60)
	c.Header("Content-Type", "image/png")
	captcha.NewImage(d, width, height).WriteTo(c.Writer)
}
func VerifyAppCaptcha(c *gin.Context, verify_value string) bool {
	value, _ := redis.String(redisModel.Get("app_captcha_value"))
	if value != "" {
		return value == verify_value
	} else {
		return false
	}
}
