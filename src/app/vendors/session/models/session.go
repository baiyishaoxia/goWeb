package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}
func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	ss := session.Get(key)
	return ss
}
func HasSession(c *gin.Context, key string) bool {
	session := sessions.Default(c)
	value := session.Get(key)
	session.Save()
	if value == nil {
		return false
	} else {
		return true
	}
}
func DeleteSession(c *gin.Context, key string) {
	session := sessions.Default(c)
	session.Delete(key)
	session.Save()
}
