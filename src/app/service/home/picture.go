package home

import (
	"app"
	"app/models"
	"databases"
	"fmt"
	"github.com/go-xorm/xorm"
	"math"
	"strconv"
)

//region    获取相册封面列表  Author:tang
func GetPictureList(wheres map[string]interface{}) ([]map[string]string, float64, float64, int) {
	var (
		db    *xorm.Session
		item  = make([]models.Picture, 0)
		num   int64
		limit = wheres["limit"].(int)
		page  = wheres["page"].(int)
	)
	if limit <= 0 {
		limit = 10
	}
	db = databases.Orm.Where("status=?", 1).Select("id,img,title,intro")
	err := *db
	num, _ = db.Count(new(models.Picture))
	all := math.Ceil(float64(num) / float64(limit))
	if page < 0 {
		page = 0
	}
	//数据
	err.Limit(limit, (page-1)*limit).Find(&item)
	data := InitPicture(item)
	return data, float64(num), all, page
}

//endregion

//region   获取相册下图片 【返回参数：数据，起始位置，总条数，总页数，当前页】   Author:tang
func GetPictureDetail(wheres map[string]interface{}) ([]map[string]interface{}, map[string]int, float64, float64, int) {
	var (
		item  = new(models.Picture)
		limit = wheres["limit"].(int)
		page  = wheres["page"].(int)
		id    = wheres["id"].(int)
	)
	if limit <= 0 {
		limit = 10
	}
	_, err := databases.Orm.Id(id).Where("status=?", 1).Select("id,img,title,intro,images").Get(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	//格式化源数据
	data := app.StrSplitArray(item.Images)
	images := make([]map[string]interface{}, len(data))
	for key, val := range data {
		images[key] = make(map[string]interface{})
		images[key]["id"] = key + 1
		images[key]["src"] = val
		images[key]["alt"] = item.Intro
	}
	all := math.Ceil(float64(len(data)) / float64(limit))
	if page < 0 {
		page = 0
	}
	start := (page - 1) * limit
	end := (page * limit)
	if start > len(data) {
		start = 0
	}
	if end > len(data) {
		end = len(data)
	}
	line := map[string]int{"start": start, "end": end}
	fmt.Println(line)
	return images[start:end], line, float64(len(data)), all, page
}

//endregion

//region   格式化数据   Author:tang
func InitPicture(item []models.Picture) []map[string]string {
	data := make([]map[string]string, len(item))
	for key, val := range item {
		data[key] = make(map[string]string)
		data[key]["id"] = strconv.Itoa(int(val.Id))
		data[key]["src"] = val.Img
		data[key]["intro"] = val.Intro
		data[key]["alt"] = val.Title
	}
	return data
}

//endregion
