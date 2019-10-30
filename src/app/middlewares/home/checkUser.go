package middlewares

import (
	"app/models"
	models2 "app/vendors/redis/models"
	"config"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := strconv.ParseInt(c.PostForm("keyid"), 10, 64)
		token := c.PostForm("token")
		//判断id对应的用户是否存在
		user := models.GetUserById(uid)
		if user == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"status": config.VueRelogin,
				"msg":    "用户不存在",
				"data":   "",
			})
		} else {
			key := user.Phone + strconv.FormatInt(int64(user.Id), 10)
			code, _ := redis.String(models2.Get(key))
			if code != token {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"status": config.VueRelogin,
					"msg":    "用户不存在",
					"data":   "",
				})
			}
		}
		c.Next()
	}
}
