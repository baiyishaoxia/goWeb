package main

import (
	"app/service/background"
	"databases"
)

func main() {
	databases.Orm.Sync(new(background.TimeLine))       //时光轴表
	databases.Orm.Sync(new(background.BannerCategory)) //图片分类表
	databases.Orm.Sync(new(background.Banner))         //图片内容表
}
