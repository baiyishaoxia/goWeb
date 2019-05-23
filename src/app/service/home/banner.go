package home

import (
	"app/service/background"
	"databases"
)

func BannerList(index string) ([]*background.Banner, string) {
	banner_cate := BannerCategoryByIndex(index)
	if banner_cate == nil {
		return nil, "无该分类"
	}
	banner := make([]*background.Banner, 0)
	err := databases.Orm.Where("banner_category_id=?", banner_cate.Id).Asc("sort").Asc("id").Find(&banner)
	if err != nil {
		return nil, ""
	}
	if len(banner) < 1 {
		return nil, ""
	}
	return banner, ""
}
