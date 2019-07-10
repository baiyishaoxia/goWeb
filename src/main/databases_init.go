package main

import (
	"app/models"
	background2 "app/service/background"
	"databases"
)

func main() {
	databases.Orm.Sync(new(background2.TimeLine))       //时光轴表
	databases.Orm.Sync(new(background2.BannerCategory)) //图片分类表
	databases.Orm.Sync(new(background2.Banner))         //图片内容表
	databases.Orm.Sync(new(background2.Message))        //留言内容表
	databases.Orm.Sync(new(models.UserQqInfo))          //QQ用户信息表
}
