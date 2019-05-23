package home

import (
	"app/models/background"
	"databases"
	"github.com/go-xorm/xorm"
	"math"
)

//region   根据分类的ID获取新闻内容列表   Author:tang
func GetNewsByCategoryId(category_id int64, limit int, nowPage int) (*[]models.Article, float64, float64, int64) {
	var (
		num  int64
		page int64
		db   *xorm.Session
	)
	if limit <= 0 {
		limit = 10
	}
	if category_id == 0 {
		//所有文章
		db = databases.Orm.OrderBy("sort asc").OrderBy("id desc") //倒序
	} else {
		//每个分类下的所有文章
		db = databases.Orm.Where("cate_id=?", category_id).OrderBy("sort asc").OrderBy("id desc") //倒序
	}
	err := *db
	num, _ = db.Count(new(models.Article))
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	//数据
	news := new([]models.Article)
	err.Limit(limit, nowPage*limit).Find(news)
	return news, float64(num), all, page + 1
}

//endregion

//region   返回推荐+浏览量的文章   Author:tang
func GetNewsByRight(limit int) ([]*models.Article, []*models.Article) {
	red_data, click_data := make([]*models.Article, 0), make([]*models.Article, 0)
	databases.Orm.Asc("sort").Where("is_red=?", true).Limit(limit).Find(&red_data) //推荐
	databases.Orm.Asc("sort").Desc("click_num").Find(&click_data)                  //浏览量
	return red_data, click_data
}

//endregion

//region   根据id获取文章信息   Author:tang
func GetArticleById(id int64) *models.Article {
	item := new(models.Article)
	databases.Orm.Where("id=?", id).Get(item)
	if item != nil {
		item.AuthorName = models.GetArticleAuthorById(int(item.AuthorId))
		item.CateName = models.GetCategoryById(item.CateId).Title
		item.ClickNum = item.ClickNum + 1
		databases.Orm.Cols("click_num").Update(item, models.Article{Id: item.Id})
	}
	return item
}

//endregion