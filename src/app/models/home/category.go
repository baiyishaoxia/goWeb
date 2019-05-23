package models

import (
	"app/models/background"
	"databases"
)

//根据Id 获取 分类  Author:tang
func GetCategoryById(id int64) *models.Category {
	item := new(models.Category)
	databases.Orm.Where("id=?", id).Get(item)
	return item
}

//region   获取新闻类别 [所有分类|顶级分类]   Author:tang
func GetCategory() (*[]models.Category, *[]models.Category) {
	category, category_p := new([]models.Category), new([]models.Category)
	databases.Orm.OrderBy("id asc").Find(category)
	databases.Orm.OrderBy("id asc").Where("parent_id=?", 0).Find(category_p)
	return category, category_p
}

//endregion
