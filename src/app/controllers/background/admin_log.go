package background

import (
	"app/models/background"
	session "app/vendors/session/models"
	"databases"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

//region Remark:请求日志[记录后台用户post请求的每一条数据] Author:tang
func RequestAllToLog(c *gin.Context) bool {
	ip := c.ClientIP()
	request, _ := json.Marshal(c.Request.ParseForm())
	admin_id := session.GetSession(c, "admin_id").(int64)
	if admin_id == 0 {
		admin_id = 0
	}
	log := &models.AdminLog{
		AdminId: admin_id,
		Ip:      ip,
		Url:     c.Request.RequestURI,
		Type:    c.Request.Method,
		Request: string(request),
		Area:    "",
	}
	if c.Request.Method == "POST" {
		res, err := databases.Orm.Insert(log)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		if res < 1 {
			fmt.Println(res)
			return false
		}
	}
	return true
}

//endregion
