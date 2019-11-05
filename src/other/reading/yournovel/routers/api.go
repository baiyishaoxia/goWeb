package routers

import (
	"github.com/gin-gonic/gin"
	"other/reading/yournovel/conf"
	"other/reading/yournovel/db/redis"
	"other/reading/yournovel/http"
	"other/reading/yournovel/middleware"
)

// 初始化Get路由
func InitGetRouter() *gin.Engine{
	router := gin.Default()
	//静态资源
	//router.LoadHTMLGlob("./other/reading/yournovel/view/**/**/*")
	//router.Static("/assets", "./other/reading/yournovel/static")
	router.LoadHTMLGlob("./views/home/reading/**/*")
	router.Static("/assets", "./public/home/reading")
	router.Static("/public", "./public")
	//初始化配置文件
	conf.InitConfig()
	redis.InitRedisClient()

    //初始化get定义路由
    v1:=router.Group("/reading")
	v1.GET("/", middleware.RequestMiddlewareWrapper(http.Home, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))
	v1.GET("/chapter", middleware.RequestMiddlewareWrapper(http.NovelChapter, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))
	v1.GET("/search", middleware.RequestMiddlewareWrapper(http.NovelSearch, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))
	v1.GET("/content", middleware.RequestMiddlewareWrapper(http.NovelContent, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))
	return  router
}

// 初始化Post路由
func initPostRouter(engine *gin.Engine) {

}
