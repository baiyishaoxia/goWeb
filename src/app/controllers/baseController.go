package controllers

import (
	"app"
	"config"
	"github.com/gin-gonic/gin"
	"net/http"
)

//基础controller，用于API继承
type BaseController struct {
	c *gin.Context
	app.Responser
}

func (a *BaseController) GinContext(c *gin.Context) {
	a.c = c
}

//成功返回
func (a *BaseController) SuccessJSON(status int, msg string, data interface{}) {
	a.c.JSON(http.StatusOK, a.Success(status, msg, data))
	return
}

//失败返回
func (base *BaseController) ErrorJSON(status int, msg string, tr ...bool) {
	if gin.Mode() != gin.DebugMode && status == config.QUERY_ERROR{
		status = config.SYSTEM_ERROR
		msg = "system error"
		tr[0] = true
	}
	translate := len(tr) > 0 && tr[0]
	base.c.JSON(http.StatusOK, base.Error(status, msg, translate))
	return
}
func (base *BaseController) InvalidArgumentJSON(errors ...string) {
	base.c.JSON(http.StatusOK, base.InvalidArgument(errors...))
}

func (base *BaseController) SystemErrorJSON(errors ...string) {
	base.c.JSON(http.StatusOK, base.SystemError(errors...))
}

func (base *BaseController) QueryErrorJSON(errors ...string) {
	if gin.Mode() != gin.DebugMode {
		errors = []string{}
	}
	base.c.JSON(http.StatusOK,  base.QueryError(errors...))
}