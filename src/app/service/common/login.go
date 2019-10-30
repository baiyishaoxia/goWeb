package common

import (
	"app/models"
	session "app/vendors/session/models"
	"databases"
	"github.com/gin-gonic/gin"
)

func ValidateLogin(c *gin.Context) *models.Users {
	user := new(models.Users)
	if session.HasSession(c, "userid") == false {
		user.Id = 0
		return user
	}
	var userid int64 = (session.GetSession(c, "userid")).(int64)
	databases.Orm.Where("id=?", userid).Get(user)
	return user
}
