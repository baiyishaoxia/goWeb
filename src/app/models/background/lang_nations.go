//region Remark:国家表 Author:tang

package models

import (
	"app"
	"databases"
	"fmt"
)

type BlogLangNations struct {
	Id        string   `json:"id"`
	ImageUrl  string   `json:"-"`
	Title     string   `json:"-"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
	Sort      int64
	IsDefault bool
	IsOpen    bool
}

//region Remark:国家列表 Author:tang
func LangNationsList(page int, limit int, keywords string) *[]BlogLangNations {
	var nations = new([]BlogLangNations)
	if keywords != "" {
		err := databases.Orm.Where("title like ?", keywords+"%").Limit(limit, page*limit).Asc("sort").Find(nations)
		if err != nil {
			fmt.Println(err.Error())
		}
		return nations
	} else {
		err := databases.Orm.Limit(limit, page*limit).Asc("sort").Find(nations)
		if err != nil {
			fmt.Println(err.Error())
		}
		return nations
	}
}

//endregion

//region Remark:总记录数 Author:tang
func LangNationsNum(keywords string) int64 {
	if keywords != "" {
		num, _ := databases.Orm.Table("blog_lang_nations").Where("title=?", keywords).Count()
		return num

	} else {
		num, _ := databases.Orm.Table("blog_lang_nations").Count()
		return num
	}
}

//endregion

//region Remark:添加国家 Author:tang
func (a *BlogLangNations) AddNations() bool {
	ok, err := databases.Orm.Insert(a)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if ok < 1 {
		fmt.Println("添加失败")
		return false
	}
	return true
}

//endregion

//region Remark:通过id查询数据 Author:tang
func GetNationsById(id string) *BlogLangNations {
	nations := new(BlogLangNations)
	has, err := databases.Orm.Where("id = ?", id).Get(nations)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return nations
}

//endregion

//region Remark:至少有一个为启用 Author:tang
func GetNationsIsOpen() int64 {
	num, _ := databases.Orm.Table("blog_lang_nations").Where("is_open = ?", true).Count()
	return num
}

//endregion

//region Remark:至少有一个为启用且不包括本身 Author:tang
func GetOneNationsIsOpen(id string) bool {
	nations := new(BlogLangNations)
	has, err := databases.Orm.Where("is_open = ? and id =?", true, id).Get(nations)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if has == false {
		return false
	}
	return true
}

//endregion

//region Remark:至少有一个为默认国家语言 Author:tang
func GetNationsIsDefault() int64 {
	num, _ := databases.Orm.Table("blog_lang_nations").Where("is_default = ?", true).Count()
	return num
}
func GetDefaultNations(id string) bool {
	nations := new(BlogLangNations)
	has, err := databases.Orm.Where("is_default = ? and id =?", true, id).Get(nations)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if has == false {
		return false
	}
	return true
}

//endregion

//region Remark:有且只有一个为默认 Author:tang
func GetOneNationsIsDefault() bool {
	nations := new(BlogLangNations)
	nations.IsDefault = false
	_, err := databases.Orm.Cols("is_default").Update(nations, BlogLangNations{IsDefault: true})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

//endregion

type LangAndNation struct {
	BlogLangNations `xorm:"extends"`
	BlogLangs       `xorm:"extends"`
}

//region Remark:判定id的标识符集合(联表) Author:tang
func GetLangAndNationViews(mark string) *[]LangAndNation {
	langss := new([]LangAndNation)
	err2 := databases.Orm.Table("blog_lang_nations").Asc("blog_lang_nations.sort").
		Join("LEFT", "blog_langs", "blog_lang_nations.id = blog_langs.lang_nations_id and blog_langs.mark = ?", mark).
		Find(langss)
	if err2 != nil {
		fmt.Println(err2.Error())
		return nil
	}
	return langss
}

//region Remark:查询该国家下是否存在语言包 Author:tang
func GetNationsInLangs(id string) string {
	langs := new([]BlogLangs)
	err := databases.Orm.Select("id,mark,title").Where("lang_nations_id = ?", id).Asc("sort").Find(langs)
	if err != nil {
		fmt.Println(err.Error())
	}
	var count int = 0
	println("-----国家启用中-----")
	count = len(*langs)

	println("语言包个数", count)
	langs2 := new(BlogLangs)
	if count == 0 {
		has, _ := databases.Orm.Limit(1).Get(langs2)
		if has == false {
			println("没有找到标识符")
		}
		return langs2.Mark
	} else {
		limit := make([]string, count)
		for key, val := range *langs {
			println("已存在的标识符:", val.Mark)
			limit[key] = val.Mark
		}
		langs3 := new(BlogLangs)
		has, err := databases.Orm.NotIn("mark", limit).Asc("sort").Get(langs3)
		if err != nil {
			fmt.Println(err.Error())
		}
		if has == false {
			return ""
		}
		return langs3.Mark
	}
	return ""
}

//endregion
//region Remark:根据标识符返回当前国家所需的提示信息 Author:tang
func GetNationsLangsByMark(nations_id string, mark string) string {
	langs := new(BlogLangs)
	has, err := databases.Orm.Where("lang_nations_id = ? and mark = ?", nations_id, mark).Select("title").Get(langs)
	if err != nil {
		fmt.Println(err.Error())
	}
	if has == false {
		return ""
	}
	return langs.Title
}

//endregion
//region Remark:获取已开启的语言包国家列表 Author: tang
func GetOpenNationsList() *[]BlogLangNations {
	lang := new([]BlogLangNations)
	err := databases.Orm.Asc("sort").Where("is_open = ?", true).
		Select("id,title,is_default,is_open").Find(lang)
	if err != nil {
		fmt.Println(err.Error())
	}
	return lang
}

//endregion
