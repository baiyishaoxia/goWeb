package background

import (
	"app/models"
	"databases"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//增
func InsertTimeLine(self *models.TimeLine) bool {
	has, err := databases.Orm.Insert(self)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if has < 1 {
		return false
	}
	return true
}

//改
func EditTimeLine(field string, self *models.TimeLine) (bool, error) {
	has, err := databases.Orm.Id(self.Id).Cols(field).Update(self)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if has < 1 {
		return false, err
	}
	return true, nil
}

//查
func FindOneTimeLine(field string, val interface{}) *models.TimeLine {
	item := new(models.TimeLine)
	_, err := databases.Orm.Where(field+" = ?", val).Get(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	return item
}
func FindAllTimeLine(field string, order string) []*models.TimeLine {
	item := make([]*models.TimeLine, 0)
	var err error
	if order == "asc" {
		err = databases.Orm.Asc(field).Find(&item)
	} else {
		err = databases.Orm.Desc(field).Find(&item)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return item
}

//根据时间段查询
func GetTimeLineByTime(year string, mouth string) *[]models.TimeLine {
	if year == "" {
		//查询不同年份的数据
		data := new([]models.TimeLine)
		err := databases.Orm.GroupBy("year").Desc("time").Find(data)
		if err != nil {
			fmt.Println(err)
		}
		return data
	} else if mouth != "" {
		//查询当前月份下数据
		data := new([]models.TimeLine)
		err := databases.Orm.Where("time like ?", year+"-"+mouth+"%").Desc("time").Find(data)
		if err != nil {
			fmt.Println(err)
		}
		return data
	} else {
		//查询当前年份下数据
		data := new([]models.TimeLine)
		err := databases.Orm.Where("time like ?", year+"%").Desc("time").Find(data)
		if err != nil {
			fmt.Println(err.Error())
		}
		return data
	}
}
func LineToLine() []map[string]interface{} {
	list := GetTimeLineByTime("", "")
	fmt.Println(*list)
	data := make([]map[string]interface{}, 0)
	item := make(map[string]interface{}) //年--月
	for _, val := range *list {
		item["year"] = val.Year
		line_time := strconv.Itoa(int(val.Year))
		son := GetTimeLineByTime(line_time, "") //当前年
		if len(*son) != 0 {
			list := make(map[string]interface{})
			pre_year,pre_month :="",""
			for _, v2 := range *son {
				line_time2 := strings.Split(v2.Time.Format("2006,01,02,15,04,05"), ",")
				//过滤相同的月份
				if line_time2[0] == pre_year && line_time2[1] == pre_month{
					continue
				}
				child := GetTimeLineByTime(line_time2[0], line_time2[1]) //当前年下的月份
				init := make([]map[string]string, len(*child))
				for k1, v3 := range *child {
					init[k1] = make(map[string]string)
					init[k1]["create_time"] = v3.Time.Format("01月02日 15:04:05")
					init[k1]["content"] = v3.Content
				}
				list[line_time2[1]] = init
				pre_year,pre_month = line_time2[0], line_time2[1]
			}
			item["month"] = list
		}
		data = append(data, item)
	}
	return data
}

//存在?
func HasTimeLine(ids []string) bool {
	item := new(models.TimeLine)
	count, err := databases.Orm.Table("`time_line`").In("id", ids).Count(item)
	if err != nil {
		fmt.Println(err.Error())
	}
	if count > 0 {
		return true
	}
	return false
}

//列表
func GetTimeLineList(page int, limit int, keywords string) (*[]models.TimeLine, float64, float64, int) {
	var data = new([]models.TimeLine)
	err := databases.Orm.Desc("id")
	if keywords != "" {
		err.Where("content like ?", "%"+keywords+"%")
	}
	err1 := *err
	num, err3 := err1.Table("time_line").Count()
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
	err2 := err.Limit(limit, page*limit).Find(data)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return data, float64(num), all, page + 1
}
