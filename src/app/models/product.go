package models

import (
	"databases"
	"fmt"
	"html/template"
	"math"
	"time"
)

type Product struct {
	Id         int64         `xorm:"pk autoincr BIGINT"`
	Title      string        `xorm:"not null unique VARCHAR(255)"` //产品标题
	CategoryId int64         `xorm:"BIGINT"`                       //分类栏目
	Sort       int64         `xorm:"default 99 INTEGER"`           //排序
	IsComment  bool          `xorm:"default true"`                 //是否允许评论
	Long       int64         `xorm:"default 99 INTEGER"`           //长
	Wide       int64         `xorm:"default 99 INTEGER"`           //宽
	High       int64         `xorm:"default 99 INTEGER"`           //高
	Address    string        `xorm:"VARCHAR(255)" json:"address"`  //产地
	Material   string        `xorm:"VARCHAR(255)" json:"address"`  //材质
	Supplier   string        `xorm:"VARCHAR(255)" json:"address"`  //供应商
	Unit       int64         `xorm:"int" json:"unit"`              //价格计算单位
	Weight     string        `xorm:"VARCHAR(255)" json:"weight"`   //重量
	Price      float64       `xorm:"double" json:"price"`          //市场价格
	Cost       float64       `xorm:"double" json:"cost"`           //成本价格
	LowPrice   float64       `xorm:"double" json:"low_price"`      //最低价
	StartTime  time.Time     `xorm:"DATETIME" json:"start_time"`   //销售开始时间
	EndTime    time.Time     `xorm:"DATETIME" json:"end_time"`     //销售结束时间
	Tags       string        `xorm:"VARCHAR(255)" json:"tags"`     //产品关键字
	Intro      string        `xorm:"VARCHAR(255)" json:"intro"`    //摘要
	Image      string        `xorm:"VARCHAR(255)" json:"image"`    //缩略图
	Content    template.HTML `xorm:"not null TEXT" json:"content"` //详细内容
	CreatedAt  time.Time     `xorm:"created"`
	UpdatedAt  time.Time     `xorm:"updated"`
}

//价格计算单位
func GetUnit() map[int]string {
	data := map[int]string{1: "件", 2: "斤", 3: "KG", 4: "吨", 5: "套"}
	return data
}

//根据ID查询产品
func GetProductById(id int64) *Product {
	data := new(Product)
	databases.Orm.Where("id=?", id).Get(data)
	return data
}

//产品管理左侧
func GetCategoryDictory() []map[string]interface{} {
	category := make([]Category, 0)
	databases.Orm.Asc("sort").Find(&category)
	category = UnlimitedForLevel(category, "", 0, 0)
	category = append(category, Category{Id: 0, TitleHtml: "无父级菜单"})
	data := make([]map[string]interface{}, len(category))
	for key, val := range category {
		data[key] = make(map[string]interface{})
		data[key]["id"] = val.Id
		data[key]["pId"] = val.ParentId
		data[key]["name"] = val.TitleHtml
		if key == 0 {
			data[key]["open"] = true
		}
	}
	return data
}

//关联分类表
type ProductAndCategory struct {
	Product  `xorm:"extends"`
	Category `xorm:"extends"`
}

//region   产品管理列表   Author:tang
func GetProductList(page int, limit int, keywords string, cate_id int64, start_time string, end_time string) (*[]ProductAndCategory, float64, float64, int) {
	var product = new([]ProductAndCategory)
	err := databases.Orm.Desc("product.id").Table("product")
	err.Join("LEFT", "category", "category_id = category.id")
	if keywords != "" {
		err.Where("product.title like ?", "%"+keywords+"%")
	}
	if cate_id != 0 {
		err.Where("product.category_id = ?", cate_id)
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
	err2 := err.Limit(limit, page*limit).Find(product)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return product, float64(num), all, page + 1
}

//endregion
