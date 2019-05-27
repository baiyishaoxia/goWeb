package models

import (
	"app/models/background"
	"databases"
)

func GetArticleById(id int64) *models.Article {
	data := new(models.Article)
	databases.Orm.Where("id=?", id).Get(data)
	return data
}
