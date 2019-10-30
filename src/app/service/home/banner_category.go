package home

import (
	"app/models"
	"databases"
)

//根据index获取分类
func BannerCategoryByIndex(index string) *models.BannerCategory {
	baner_category := new(models.BannerCategory)
	has, err := databases.Orm.Where("`index` = ?", index).Get(baner_category)
	if !has || err != nil {
		return nil
	}
	return baner_category
}
