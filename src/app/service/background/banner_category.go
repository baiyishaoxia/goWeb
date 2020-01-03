package background

import (
	"app/models"
	"databases"
	"fmt"
	"github.com/go-xorm/xorm"
	"math"
)

//获取列表
func PageBannerCategoryList(page int, limit int, keywords string) (*[]models.BannerCategory, float64, float64, int) {
	data := new([]models.BannerCategory)
	err := databases.Orm.Table("banner_category").Desc("id")
	err1 := *err
	num, _ := err.Count()
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	if page >= int(all) {
		page = int(all) - 1
	}
	//数据
	err2 := err1.Limit(limit, page*limit-limit).Find(data)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return data, float64(num), all, page + 1
}

//根据index获取分类
func  BannerCategoryByIndex(index string) *[]models.Banner {
	item := new(models.BannerCategory)
	has, err := databases.Orm.Where("`index` = ?", index).Get(item)
	if !has || err != nil {
		return nil
	}
	var db  *xorm.Session
	var list = new([]models.Banner)
	if item.Id == 0 {
		//所有图片
		db = databases.Orm.Asc("sort").Asc("id") //ID升序
	} else {
		//每个分类下的所有图片
		db = databases.Orm.Where("banner_category_id=?",item.Id).Asc("sort").Asc("id")
	}
	_err:=db.Find(list)
	if _err!=nil{

	}
	return list
}
