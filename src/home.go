package main

import (
	"app/channel"
	"github.com/gin-gonic/gin"
	"routers"
)

func main() {
	//数据库初始化测试
	//databases.Init()
	//路由加载
	go channel.HandleConcurrent()
	gin.SetMode(gin.DebugMode)
	router := routers.InitHomeRouter()
	router.Run(":9090")
}
