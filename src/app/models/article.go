package models

import (
	"app"
	"databases"
	"fmt"
	"html/template"
	"math"
)

type Article struct {
	Id         int64         `xorm:"pk autoincr BIGINT" json:"id"`
	Title      string        `xorm:"not null unique VARCHAR(255)" json:"title"`
	Img        string        `xorm:"VARCHAR(255)" json:"img"`
	Content    template.HTML `xorm:"TEXT" json:"content"`
	CateId     int64         `xorm:"BIGINT" json:"cate_id"`
	Type       int           `xorm:"default 0 INTEGER" json:"type"`
	Sort       int           `xorm:"default 99 INTEGER" json:"sort"`
	CountNum   int64         `xorm:"not null default 0 INTEGER" json:"count_num"` //评论量
	ClickNum   int64         `xorm:"not null default 0 INTEGER" json:"click_num"` //浏览量
	Status     int           `xorm:"not null default 1 INTEGER" json:"status"`    //状态(1发布2下架)
	Keywords   string        `xorm:"VARCHAR(255)" json:"keywords"`
	Intro      string        `xorm:"VARCHAR(255)" json:"intro"`
	AuthorId   int64         `xorm:"BIGINT" json:"author_id"`
	Source     string        `xorm:"VARCHAR(255)" json:"source"`
	IsComment  bool          `xorm:"bool default true" json:"is_comment"` //是否评论
	IsRed      bool          `xorm:"bool default true" json:"is_red"`     //是否推荐
	IsTop      bool          `xorm:"bool default true" json:"is_top"`     //是否置顶
	StartTime  string        `xorm:"VARCHAR(255)" json:"start_time"`
	EndTime    string        `xorm:"VARCHAR(255)" json:"end_time"`
	CreatedAt  app.Time      `xorm:"created" json:"created_at"`
	UpdatedAt  app.Time      `xorm:"updated" json:"updated_at"`
	CateName   string        `xorm:"- <- ->"`
	AuthorName string        `xorm:"- <- ->"`
}

//关联分类表
type ArticleAndCategory struct {
	Article  `xorm:"extends"`
	Category `xorm:"extends"`
}

//自定义常量
const (
	UserStatue        = 0
	UserStatueWaiting = 1
	UserStatueChecked = 2
)

//region Remark:文章列表 Author:tang
func GetArticleList(page int, limit int, keywords string, cate_id int64, start_time string, end_time string) (*[]ArticleAndCategory, float64, float64, int) {
	var art = new([]ArticleAndCategory)
	err := databases.Orm.Desc("article.id").Table("article")
	err.Join("LEFT", "category", "cate_id = category.id")
	if keywords != "" {
		err.Where("article.title like ?", "%"+keywords+"%")
	}
	if cate_id != 0 {
		err.Where("article.cate_id = ?", cate_id)
	}
	if start_time != "" && end_time != "" {
		err.Where("start_time < ?", start_time).Where("end_time > ?", end_time)
	}
	err1 := *err
	num, err3 := err1.Count()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	err2 := err.Limit(limit, page*limit).Find(art)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return art, float64(num), all, page + 1
}

//endregion
//region Remark:自定义文章类型 Author:tang
func GetArticleType() map[int]string {
	article_type := map[int]string{1: "帮助说明", 2: "新闻资讯", 3: "军事政策", 4: "娱乐圈"}
	return article_type
}
func GetArticleTypeById(id int) string {
	switch id {
	case 1:
		return "帮助说明"
	case 2:
		return "新闻资讯"
	case 3:
		return "军事政策"
	case 4:
		return "娱乐圈"
	default:
		return "--"
	}
}

//endregion
//region Remark:自定义文章作者 Author:tang
func GetArticleAuthor() map[int]string {
	article_type := map[int]string{1: "白衣少侠", 2: "张三", 3: "李四", 4: "王五"}
	return article_type
}
func GetArticleAuthorById(id int) string {
	switch id {
	case 1:
		return "白衣少侠"
	case 2:
		return "张三"
	case 3:
		return "李四"
	case 4:
		return "王五"
	default:
		return "--"
	}
}

//endregion

//region Remark:自定义搜索文章方法 [status为1已发布，2已下架]  [way 1升序 2降序  type_key根据某一字段排序(count)] Author:tang
func SearchArticleBykeys(keywords string, imgHost string, limit int64, page int64, way string, type_key string) ([]Article, int64, float64, int64) {
	article := new([]ArticleAndCategory)
	//关联分类表
	err := databases.Orm.Table("article").Join("LEFT", "category", "cate_id = category.id")
	if way == "1" {
		if type_key == "count" {
			err.Asc("count")
		}
	} else {
		if type_key == "count" {
			err.Desc("count")
		}
	}
	err1 := err.Where("article.status = ?", 1).Where("article.title like ?", "%"+keywords+"%").Find(article)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	length := len(*article)
	data := make([]Article, length)
	for key, val := range *article {
		data[key].Id = val.Article.Id
		data[key].Title = val.Article.Title
		data[key].Intro = val.Intro
		data[key].Img = imgHost + "/" + val.Article.Img
		data[key].Source = val.Source
		data[key].CreatedAt = val.Article.CreatedAt
		data[key].AuthorId = val.AuthorId
		if GetCategoryById(val.CateId) == nil {
			data[key].CateName = "--"
		} else {
			data[key].CateName = val.Category.Title
		}
		data[key].ClickNum = val.ClickNum
	}
	//总记录数
	num := float64(len(data))
	//防止超出切片范围
	if limit > int64(num) {
		limit = int64(num) - ((page - 1) * limit)
	}
	//总页数
	if page < 0 {
		page = 1
	}
	all := math.Ceil(float64(length) / float64(limit))
	if float64(page) > all {
		page = int64(all)
		return data[0:0], page, all, int64(num)
	}
	if float64(page*limit) > num {
		return data[(page-1)*limit:], page, all, int64(num)
	}
	return data[(page-1)*limit : page*limit], page, all, int64(num)
}

//endregion
func GetArticleById(id int64) *Article {
	data := new(Article)
	databases.Orm.Where("id=?", id).Get(data)
	return data
}