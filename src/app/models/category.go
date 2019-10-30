package models

import (
	"app"
	newredis "app/vendors/redis/models"
	"databases"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"html/template"
	"math"
	"strings"
)

type Category struct {
	Id        int64 `xorm:"pk autoincr BIGINT"`
	ParentId  int64
	Title     string `xorm:"not null"`
	CallIndex string `xorm:"unique VARCHAR(255)"`
	ClassList string
	Link      string
	Img       string
	Content   string
	Sort      int64
	Status    int64
	Level     int           `xorm:"- <- ->"`
	TitleHtml template.HTML `xorm:"- <- ->"`
	CreatedAt app.Time      `xorm:"created"`
	UpdatedAt app.Time      `xorm:"updated"`
}

//region Remark:获取所有父亲分类 Author:tang
func GetAllCategory() []Category {
	cate := make([]Category, 0)
	err := databases.Orm.Asc("sort").Find(&cate)
	if err != nil {
		fmt.Println("分类模块出现bug...", err)
	}
	cate = UnlimitedForLevel(cate, "|--", 0, 0)
	cate = Array2ToArray1(&cate)
	return cate
}

//endregion
type Tree struct {
	Id       int64
	ParentId int64
	Level    int
	Html     string
	Title    string
}

//region Remark:树形 Author:tang
func UnlimitedForLevel(data []Category, html string, pid int64, level int) (res []Category) {
	for _, v := range data {
		if v.ParentId == pid {
			v.Level = level + 1
			newhtml := strings.Repeat(html, v.Level)
			v.TitleHtml = template.HTML(newhtml + v.Title)
			res = append(res, v)
			res1 := UnlimitedForLevel(data, html, v.Id, level+1)
			for _, v1 := range res1 {
				res = append(res, v1)
			}
		}
	}
	return res
}

//endregion

//region Remark:数组格式化 Author:tang
func Array2ToArray1(data *[]Category) []Category {
	arr := make([]Category, len(*data))
	for key, value := range *data {
		arr[key].Id = value.Id
		arr[key].Title = value.Title
		fmt.Println("正在建立树形生成分类下拉列表-----------", value.Title)
	}
	return arr
}

//endregion

//region Remark:添加和修改的缓存 Author:tang
func CategoryTreeData() (data []Category) {
	//判断redis是否有缓存数据
	var redis_key string = "admin:category:tree"
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &data)
	} else {
		category := make([]Category, 0)
		databases.Orm.Asc("sort").Find(&category)
		data = UnlimitedForLevel(category, "|--", 0, 0)
		data = append(data, Category{Id: 0, TitleHtml: "无父级菜单"})
		//缓存到redis
		value, _ := json.Marshal(data)
		newredis.Set(redis_key, value, 60*60)
	}
	return data
}

//endregion

//region Remark:分类列表 Author:tang
func GetCategoryList(page int, limit int, keywords string) ([]Category, float64, float64, int) {
	var cate = new([]Category)
	err := databases.Orm.Asc("sort")
	if keywords != "" {
		err.Where("title like ?", "%"+keywords+"%")
	}
	err1 := *err
	//记录数
	num, err3 := err1.Table("category").Count()
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
	//分页后的记录
	err2 := err.Limit(limit, page*limit).Find(cate)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return *cate, float64(num), all, page + 1
}
//endregion

//region Remark:根据id获取分类信息 Author:tang
func GetCategoryById(id int64) *Category {
	var category = new(Category)
	has, err := databases.Orm.Where("id = ?", id).Get(category)
	if has == false {
		return nil
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return category
}

//endregion


//region   获取新闻类别 [所有分类|顶级分类]   Author:tang
func GetCategory() (*[]Category, *[]Category) {
	category, category_p := new([]Category), new([]Category)
	databases.Orm.OrderBy("id asc").Find(category)
	databases.Orm.OrderBy("id asc").Where("parent_id=?", 0).Find(category_p)
	return category, category_p
}

//endregion
