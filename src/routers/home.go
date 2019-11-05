package routers

import (
	"app"
	"app/controllers/home"
	"app/controllers/home/blog"
	"app/middlewares/common"
	"app/middlewares/home"
	fun "app/service/common"
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
		Root:         "views/home",
		Extension:    ".html",
		Master:       "layouts/default/main",
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
	router.GET("/captcha/:width/:height", captcha.GetCaptcha)
	//自定义组
	router.GET("/", blog.GetMain)                                      //默认首页
	router.GET("/article", blog.GetBlogArticle)                        //文章专栏
	router.POST("/article/ajax", blog.PostArticleList)                 // 文章专栏(Ajax)
	router.POST("/category/ajax", blog.PostArticleCatgory)             // 文章分类(Ajax)
	router.GET("/article/right", blog.GetArticleRight)                 // 文章推荐(Ajax)
	router.GET("/hot/user", blog.GetHotMessageUser)                    // 热评用户(Ajax)
	router.GET("/article/detail/:id", blog.GetBlogArticleDetail)       //文章详情
	router.GET("/mixed/pic", blog.GetBlogMixedPic)                     //杂七杂八
	router.POST("/mixed/pic/ajax", blog.PostMixedPic)                  //杂七杂八(Ajax)
	router.GET("/mixed/pic/detail/:id", blog.GetBlogMixedPicDetail)    //相册详情
	router.POST("/mixed/pic/detail/ajax", blog.PostMixedPicDetailAjax) //相册详情(Ajax)
	router.GET("/time/line", blog.GetBlogTimeLine)                     //点点滴滴
	router.GET("/time/line/ajax", blog.GetBlogTimeLineAjax)            //点点滴滴(Ajax)
	router.GET("/about", blog.GetBlogAbout)                            //关于本站
	router.POST("/about/ajax", blog.PostBannerList)                    //关于本站、关键字获取(Ajax)
	router.GET("/sigu/video", blog.GetBlogSiguVideo)                   //思古视频
	router.GET("/qq/login", blog.GetBlogQQLogin)                       //QQ互联登录
	router.GET("/logout", blog.GetBlogExit)                            //退出
	router.GET("/message/ajax", blog.GetBlogMessageAjax)               //获取留言(Ajax)
	router.POST("/message/create", blog.PostBlogMessageCreate)         //提交文章评论、留言(Ajax)
	router.GET("/download", blog.FileDownload)                         //文件下载
	router.GET("/site/count", fun.SiteStatistic)                       //站点统计
	//---------------------------------API---------------------------------
	v2 := router.Group("/v2")
	v2.GET("/test", fun.GetTest)
	return router
}
