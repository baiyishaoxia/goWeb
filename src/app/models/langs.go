//region Remark:语言表 Author:tang
package models

import (
	"app"
	"databases"
	"fmt"
	"math"
)

type Langs struct {
	Id            int64
	LangNationsId string
	Mark          string
	Remark        string
	Title         string
	Sort          int64
	CreatedAt     app.Time `xorm:"created"`
	UpdatedAt     app.Time `xorm:"updated"`
}

//region Remark:获取所有国家 Author:tang
func GetLangsNations() *[]LangNations {
	var nations = new([]LangNations)
	databases.Orm.Select("id,title").Asc("sort").Find(nations)
	return nations
}

//endregion

//region Remark:添加语言 Author:tang
func (a *Langs) AddLangs() bool {
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

//region Remark:获取所有语言记录数 Author:tang
func LangNum(keywords string) int64 {
	if keywords != "" {
		num, _ := databases.Orm.Table("blog_langs").Where("mark=?", keywords).Count()
		return num
	} else {
		num, _ := databases.Orm.Table("blog_langs").Count()
		return num
	}
}

//endregion

//region Remark:列表 Author:tang
func LangList(page int, limit int, keywords string, nations_id string) (*[]Langs, float64, float64, int) {
	var langs = new([]Langs)
	err := databases.Orm.Asc("sort")
	if nations_id != "all" && nations_id != "" {
		err.Where("lang_nations_id= ?", nations_id)
	}
	if keywords != "" {
		err.Where("mark = ? or title like ?", keywords, "%"+keywords+"%")
	}
	err1 := *err
	//记录数
	num, err3 := err1.Table("langs").Count()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	//分页
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	err2 := err.Limit(limit, page*limit).Asc("sort").Find(langs)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return langs, float64(num), all, page + 1
}

//endregion

type Lang struct {
	Langs       `xorm:"extends"`
	LangNations `xorm:"extends"`
}

//region Remark:判定id是否存在 Author:tang
func GetLangInId(id string, mark string) bool {
	lang := new(Langs)
	has, err := databases.Orm.Where("lang_nations_id = ?", id).Where("mark = ?", mark).Get(lang)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if has {
		return true
	}
	return false
}

//endregion

//region Remark:国家与语言之间的对应 Author:tang
func GetNationsByLangs(id string) []string {
	lang := new(Langs)
	has, err := databases.Orm.Where("id = ?", id).Get(lang)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	langs := new([]Langs)
	println("make为:", lang.Mark)
	err2 := databases.Orm.Where("mark= ?", lang.Mark).Select("lang_nations_id,title").Asc("id").Find(langs)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	datas := *langs
	println("总计:", len(datas))
	info := make([]string, len(datas))

	for key, val := range datas {
		info[key] = val.Title
		println("-----title------")
		println("国家语言值:", info[key])
	}
	return info
}

//endregion

//region Remark:通过id查询 Author:tang
func GetLangById(id string) *Langs {
	lang := new(Langs)
	has, err := databases.Orm.Where("id = ?", id).Get(lang)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return lang
}

//endregion

//region Remark:获取国家 Author:tang
func GetNations() *[]LangNations {
	nations := new([]LangNations)
	err := databases.Orm.Select("id,title").Asc("sort").Find(nations)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nations
}

//endregion

//region Remark:获取当前国家的语言包 Author:tang
func GetNowNationsLangs(nat string) map[string]string {
	nations := new(LangNations)
	has, err := databases.Orm.Where("id = ?", nat).Get(nations)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	langs := new([]Langs)
	err2 := databases.Orm.Select("id,mark,title").Where("lang_nations_id = ?", nations.Id).Asc("sort").Find(langs)
	if err2 != nil {
		fmt.Println(err2.Error())
		return nil
	}
	var lang_data map[string]string = make(map[string]string, len(*langs))
	for _, val := range *langs {
		lang_data[val.Mark] = val.Title
	}
	return lang_data
}

//endregion
