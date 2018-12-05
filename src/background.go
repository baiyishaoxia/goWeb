package main

import (
	"databases"
	"github.com/gin-gonic/gin"
	"routers"
)

func main() {
	//数据迁移表
	//databases.Orm.Sync(new(models.Article))
	//databases.Orm.Sync(new(models.BlogAdmin))
	//databases.Orm.Sync(new(models.Config))
	//databases.Orm.Sync(new(models.Video))
	//databases.Orm.Sync(new(models.AdminLog))
	//databases.Orm.Sync(new(models.Picture))
	//databases.Orm.Sync(new(models.Users))
	//路由加载
	gin.SetMode(gin.DebugMode)
	router := routers.InitBackGroundRouter()
	defer databases.Orm.Close()
	router.Run(":9091")
}
