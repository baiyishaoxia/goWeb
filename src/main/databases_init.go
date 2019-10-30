package main

import (
	"app/models"
	"databases"
)

//初始化数据表
func main() {
	//语言相关
	databases.Orm.Sync(new(models.LangNations))
	databases.Orm.Sync(new(models.Langs))

	//管理员相关
	databases.Orm.Sync(new(models.BlogAdmin))
	databases.Orm.Sync(new(models.AdminRole))
	databases.Orm.Sync(new(models.AdminNavigation))
	databases.Orm.Sync(new(models.AdminNavigationNode))
	databases.Orm.Sync(new(models.AdminRoleNode))
	databases.Orm.Sync(new(models.AdminRoleNodeRoutes))
	databases.Orm.Sync(new(models.AdminLog))

	//banner相关
	databases.Orm.Sync(new(models.BannerCategory)) //图片分类表
	databases.Orm.Sync(new(models.Banner))         //图片内容表

	//文章相关
	databases.Orm.Sync(new(models.Category))       //文章类别
	databases.Orm.Sync(new(models.Article))        //文章内容
	databases.Orm.Sync(new(models.Message))        //留言内容表

	//图片相关
	databases.Orm.Sync(new(models.Picture))

	//产品相关
	databases.Orm.Sync(new(models.Product))

	//视频相关
	databases.Orm.Sync(new(models.Video))

    //其他相关
	databases.Orm.Sync(new(models.Users))          //用户表
	databases.Orm.Sync(new(models.UserQqInfo))     //QQ用户信息表
	databases.Orm.Sync(new(models.TimeLine))       //时光轴表
	databases.Orm.Sync(new(models.Config))         //系统配置
}