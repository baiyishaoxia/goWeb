package home

import (
	"app/service/background"
	"databases"
)

func BannerCategoryByIndex(index string) *background.BannerCategory {
	baner_category := new(background.BannerCategory)
	has, err := databases.Orm.Where("`index` = ?", index).Get(baner_category)
	if !has || err != nil {
		return nil
	}
	return baner_category
}
