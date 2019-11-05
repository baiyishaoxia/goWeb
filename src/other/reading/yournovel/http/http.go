package http

import (
	"github.com/gin-gonic/gin"
	"other/reading/yournovel/conf"
	"other/reading/yournovel/db/redis"
	"other/reading/yournovel/tool"
)

func Init()  error{
	g := gin.Default()
	initRouter(g)
	err := g.Run(tool.GetHost())
	if err != nil {
		panic(err)
	}
	return err
}

// 初始化路由
func initRouter(engine *gin.Engine) {
	conf.InitConfig()
	redis.InitRedisClient()
}