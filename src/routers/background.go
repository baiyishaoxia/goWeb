package routers

import (
	"app"
	"app/controllers/background"
	background2 "app/middlewares/background"
	"app/middlewares/common"
	captcha "app/vendors/captcha/controllers"
	upload "app/vendors/upload/controllers"
	"github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func InitBackGroundRouter() *gin.Engine {
	router := gin.Default()
	//Session初始化
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	//后台请求日志
	router.Use(common.Web("background"))
	//静态资源
	router.Static("/public", "./public")
	router.Static("/uploads", "./uploads")
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views/background",
		Extension:    ".html",
		Master:       "layouts/main",
		DisableCache: true,
		Funcs:        app.TemplateFunc(),
	})
	router.GET("/captcha/:width/:height", captcha.GetCaptcha) //图片验证码
	router.GET("/login", background.GetLogin)
	router.POST("/login", background.PostLogin)
	//登陆后的router
	v1 := router.Group("/admin", background2.Auth())
	v1.POST("/upload/image", upload.PostUpLoadImg)
	v1.GET("/upload/ueditor", upload.GetUpLoadImg)
	v1.POST("/upload/ueditor", upload.PostUEditUpload)
	v1.POST("/upload/file", upload.PostUpLoadFile)
	v1.POST("/upload/video", upload.PostUpLoadVideo)
	v1.POST("/upload/wang_editor/image", upload.PostUploadWangEditorImage)
	v1.GET("/main", background.GetIndex)
	v1.GET("/center", background.GetCenter)
	v1.GET("/exit", background.GetExit)
	v1.GET("/clear", background.GetClear)
	v1.GET("/repass", background.GetAdminPass)
	v1.POST("/repass", background.PostAdminPass)
	//资讯管理
	v1.GET("/article/list", background.GetArticleList)
	v1.GET("/article/create", background.GetArticleCreate)
	v1.POST("/article/create", background.PostArticleCreate)
	v1.GET("/article/edit/:id", background.GetArticleEdit)
	v1.POST("/article/edit/:id", background.PostArticleEdit)
	v1.POST("/article/del", background.PostArticleDel)
	v1.GET("/article/status", background.GetArticleSetStatus)
	v1.GET("/article/template", background.GetArticleTemplate)
	//图片管理
	v1.GET("/picture/list", background.GetPictureList)
	v1.GET("/picture/show/:id", background.GetPictureShow)
	v1.GET("/picture/create", background.GetPictureCreate)
	v1.POST("/picture/create", background.PostPictureCreate)
	v1.GET("/picture/edit/:id", background.GetPictureEdit)
	v1.POST("/picture/edit/:id", background.PostPictureEdit)
	v1.POST("/picture/del", background.PostPictureDel)
	v1.POST("/picture/show/del", background.PostPictureShowDel)
	v1.GET("/picture/status", background.GetPictureSetStatus)
	//视频管理
	v1.GET("/video/list", background.GetVideoList)
	v1.GET("/video/create", background.GetVideoCreate)
	v1.POST("/video/create", background.PostVideoCreate)
	v1.GET("/video/edit/:id", background.GetVideoEdit)
	v1.POST("/video/edit/:id", background.PostVideoEdit)
	v1.POST("/video/del", background.PostVideoDel)
	//产品管理
	//by 品牌
	v1.GET("/brand/list", background.GetBrandList)
	//by 分类
	v1.GET("/category/list", background.GetCategoryList)
	v1.GET("/category/create", background.GetCategoryCreate)
	v1.POST("/category/create", background.PostCategoryCreate)
	v1.GET("/category/edit/:id", background.GetCategoryEdit)
	v1.POST("/category/edit/:id", background.PostCategoryEdit)
	v1.POST("/category/del", background.PostCategoryDel)
	v1.GET("/category/status", background.GetCategorySetStatus)
	v1.POST("/category/save", background.PostCategorySave)
	//by 产品
	v1.GET("/product/list", background.GetProductList)
	v1.GET("/product/znodes", background.GetProductZnodes)
	v1.GET("/product/create", background.GetProductCreate)
	v1.POST("/product/create", background.PostProductCreate)
	v1.GET("/product/edit/:id", background.GetProductEdit)
	v1.POST("/product/edit/:id", background.PostProductEdit)
	v1.POST("/product/del", background.PostProductDel)
	//评论管理
	v1.GET("/comments/list", background.GetCommentsList)
	v1.GET("/comments/status", background.GetCommentsStatus)
	v1.POST("/comments/del", background.PostCommentDel)
	v1.GET("/feedback/list", background.GetFeedbackList)
	//管理员管理
	//by 角色管理
	v1.GET("/role/list", background.GetRoleList)
	v1.GET("/role/create", background.GetRoleCreate)
	v1.POST("/role/create", background.PostRoleCreate)
	v1.GET("/role/edit/:id", background.GetRoleEdit)
	v1.POST("/role/edit/:id", background.PostRoleEdit)
	v1.POST("/role/del", background.PostRoleDel)
	//by  导航管理
	v1.GET("/navigation", background.GetNavigation)
	//by 权限管理
	v1.GET("/permission/list", background.GetNavigationList)
	v1.GET("/permission/create/:id", background.GetNavigationCreate)
	v1.POST("/permission/create", background.PostNavigationCreate)
	v1.GET("/permission/edit/:id", background.GetNavigationEdit)
	v1.POST("/permission/edit/:id", background.PostNavigationEdit)
	v1.POST("/permission/del", background.PostNavigationDel)
	v1.POST("/permission/save", background.PostNavigationSave)
	//by 管理员
	v1.GET("/list", background.GetAdminList)
	v1.GET("/add", background.GetAdminCreate)
	v1.POST("/add", background.PostAdminCreate)
	v1.GET("/edit/:id", background.GetAdminEdit)
	v1.POST("/edit/:id", background.PostAdminEdit)
	v1.POST("/del", background.PostAdminDel)
	v1.GET("/list/status", background.GetAdminStatus)
	//会员管理
	v1.GET("/member/list", background.GetMemberList)
	v1.GET("/member/create", background.GetMemberCreate)
	v1.POST("/member/create", background.PostMemberCreate)
	v1.GET("/member/edit/:id", background.GetMemberEdit)
	v1.POST("/member/edit/:id", background.PostMemberEdit)
	v1.GET("/member/recycle", background.GetMemberDelList)
	v1.GET("/member/password/edit/:id", background.GetMemberPassword)
	v1.POST("/member/password/edit/:id", background.PostMemberPassword)
	v1.GET("/member/show/:id", background.GetMemberShow)
	v1.POST("/member/del", background.PostMemberDel)
	v1.GET("/member/status", background.GetMemberStatus)
	v1.GET("/member/csv", background.GetImportCsv)
	v1.POST("/member/csv", background.PostImportCsv)
	v1.GET("/member/csv/down/:name", background.GetImportCsvDownload)
	//系统统计
	v1.GET("/charts/zx", background.GetChartsZx)
	v1.GET("/charts/sj", background.GetChartsSj)
	v1.GET("/charts/qy", background.GetChartsQy)
	v1.GET("/charts/zz", background.GetChartsZz)
	v1.GET("/charts/bz", background.GetChartsBz)
	v1.GET("/charts/3Dbz", background.GetCharts3Dbz)
	v1.GET("/charts/3Dzz", background.GetCharts3Dzz)
	//系统设置
	v1.GET("/system/base", background.GetSystemBase)
	v1.POST("/system/base", background.PostSystemBase)
	v1.GET("/system/category", background.GetSystemCategory)
	v1.GET("/system/category/create", background.GetSystemCategoryCreate)
	v1.GET("/system/category/edit/:id", background.GetSystemCategoryEdit)
	v1.GET("/system/data", background.GetSystemData)
	v1.GET("/system/shielding", background.GetSystemShield)
	//by 后台日志
	v1.GET("/system/log", background.GetSystemLog)
	v1.GET("/system/log/show/:id", background.GetSystemLogShow)
	//当前国家语言列表
	router.GET("/nations/langs", background.GetLangApi)
	//全局搜索
	v1.GET("/search/all", background.GetSearchList)
	v1.GET("/search/export", background.GetSearchAllCsv)
	//region Remark:国家 Author:tang
	v1.GET("/nations/list", background.GetLangNationsIndex)
	v1.GET("/nations/create", background.GetLangNationsCreate)
	v1.POST("/nations/create", background.PostLangNationsCreate)
	v1.GET("/nations/edit/:id", background.GetLangNationsEdit)
	v1.POST("/nations/edit/:id", background.PostLangNationsEdit)
	v1.POST("/nations/del", background.PostLangNationsDel)
	v1.POST("/nations/save", background.PostLangNationsSave)
	//endregion

	//region Remark:语言包管理 Author:tang
	v1.GET("/lang/list", background.GetLangIndex)
	v1.GET("/lang/create", background.GetLangCreate)
	v1.POST("/lang/create", background.PostLangCreate)
	v1.GET("/lang/edit/:id", background.GetLangEdit)
	v1.POST("/lang/edit/:id", background.PostLangEdit)
	v1.POST("/lang/del", background.PostLangDel)
	//endregion

	//region Remark:数据库管理 Author:tang
	v1.GET("/databases/list", background.GetDatabaseList)
	v1.GET("/database/backup", background.GetDatabaseBackup)
	v1.GET("/database/del/:name", background.GetDatabaseDel)
	v1.POST("/database/delAll", background.GetDatabaseDelAll)
	v1.GET("/database/down/:name", background.GetDatabaseDown)
	//endregion

	//-------------------Blog子站路由-------------------
	v1.GET("/time_line/list", background.GetTimeLineList)
	v1.GET("/time_line/create", background.GetTimeLineCreate)
	v1.POST("/time_line/create", background.PostTimeLineCreate)
	v1.GET("/time_line/edit/:id", background.GetTimeLineEdit)
	v1.POST("/time_line/edit/:id", background.PostTimeLineEdit)
	v1.POST("/time_line/del", background.PostTimeLineDel)
	//banner图分类管理
	v1.GET("/banner_category/list", background.GetBannerCategoryList)
	v1.GET("/banner_category/create", background.GetBannerCategoryCreate)
	v1.POST("/banner_category/create", background.PostBannerCategoryCreate)
	v1.GET("/banner_category/edit/:id", background.GetBannerCategoryEdit)
	v1.POST("/banner_category/edit/:id", background.PostBannerCategoryEdit)
	v1.POST("/banner_category/del", background.PostBannerCategoryDel)
	//banner图内容管理
	v1.GET("/banner/list", background.GetBannerList)
	v1.GET("/banner/create", background.GetBannerCreate)
	v1.POST("/banner/create", background.PostBannerCreate)
	v1.GET("/banner/edit/:id", background.GetBannerEdit)
	v1.POST("/banner/edit/:id", background.PostBannerEdit)
	v1.POST("/banner/del", background.PostBannerDel)
	v1.POST("/banner/save", background.PostBannerSave)
	return router
}
