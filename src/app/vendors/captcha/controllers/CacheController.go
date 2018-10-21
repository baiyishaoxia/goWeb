package controllers

import (
	"config"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"strings"
)

//数字验证码配置
var configD = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}

//声音验证码配置
var configA = base64Captcha.ConfigAudio{
	CaptchaLen: 6,
	Language:   "zh",
}

//字符,公式,验证码配置
var configC = base64Captcha.ConfigCharacter{
	Height: 50,
	Width:  100,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeNumberAlphabet,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         4,
}

// base64Captcha create http handler
func GenerateCaptchaHandler(c *gin.Context) {
	key := c.Query("tmp") //随机码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//captchaId, captcaInterfaceInstance := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	_, captcaInterfaceInstance := base64Captcha.GenerateCaptcha(key, configC)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)
	image := strings.Split(base64blob, ",")[1]
	c.JSON(http.StatusOK, gin.H{
		"status": config.VueSuccess, "data": image, "msg": "success", /* "captchaId": captchaId,*/
	})
}

// base64Captcha verify http handler
func CaptchaVerifyHandle(id string, code string) bool {
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(id, code)
	return verifyResult
}
