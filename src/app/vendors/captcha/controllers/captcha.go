package controllers

import (
	captcha "app/vendors/captcha/models"
	session "app/vendors/session/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCaptcha(c *gin.Context) {
	width, _ := strconv.Atoi(c.Param("width"))
	height, _ := strconv.Atoi(c.Param("height"))
	d := make([]byte, 4)
	s := captcha.NewLen(4)
	char := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		char += strconv.FormatInt(int64(d[v]), 32)
	}
	session.SetSession(c, "captcha_value", char)
	c.Header("Content-Type", "image/png")
	captcha.NewImage(d, width, height).WriteTo(c.Writer)
}
func VerifyCaptcha(c *gin.Context, verify_value string) bool {
	value := session.GetSession(c, "captcha_value")
	if value != nil {
		return value.(string) == verify_value
	} else {
		return false
	}
}
