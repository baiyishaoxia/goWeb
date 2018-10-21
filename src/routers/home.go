package routers

import (
	"app"
	"app/controllers/home"
	"app/middlewares/common"
	"app/middlewares/home"
	captcha "app/vendors/captcha/controllers"
	upload "app/vendors/upload/controllers"
	"github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func InitHomeRouter() *gin.Engine {
	router := gin.Default()
	//Session初始化
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	//前台请求日志
	router.Use(common.Web("home"))
	//静态资源
	router.Static("/public", "./public")
	router.Static("/uploads", "./uploads")
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "html/home",
		Extension:    ".html",
		Master:       "layouts/main",
		DisableCache: true,
		Funcs:        app.TemplateFunc(),
	})
	router.GET("/test", home.GetTest)
	//用户验证组
	user := router.Group("/user", middlewares.CheckLogin())
	user.GET("/index", home.GetIndex)
	//前台路由组
	v1 := router.Group("/home")
	v1.POST("/upload/image", upload.PostUpLoadImg)
	v1.POST("/upload/file", upload.PostUpLoadFile)
	v1.POST("/upload/video", upload.PostUpLoadVideo)
	v1.POST("/upload/wang_editor/image", upload.PostUploadWangEditorImage)
	v1.GET("/index/list", home.GetIndex)
	v1.GET("/index/create", home.GetIndexCreate)
	v1.POST("/index/create", home.PostIndexCreate)
	v1.GET("/index/edit/:id", home.GetIndexEdit)
	v1.POST("/index/edit/:id", home.PostIndexEdit)
	v1.POST("/index/del", home.PostLinksDel)
	v1.POST("/index/save", home.PostLinksSort)
	//自定义组
	router.GET("/", home.GetMain)
	router.GET("/captcha/:width/:height", captcha.GetCaptcha)
	return router
}
