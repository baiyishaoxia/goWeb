package job

import (
	"app/models"
	"databases"
	"fmt"
)

var UsereLevelChan = make(chan string, 10000)

func HandleUsereLevel() {
	//用户活跃度自动升级
	data := models.GetUsersList2()
	for key, val := range data {
		if val.HotCount >= 10 && val.HotCount <= 100 {
			val.Level = 1
		}
		if val.HotCount > 100 && val.HotCount <= 200 {
			val.Level = 2
		}
		if val.HotCount > 200 && val.HotCount <= 500 {
			val.Level = 3
		}
		if val.HotCount > 500 && val.HotCount <= 1000 {
			val.Level = 4
		}
		if val.HotCount > 1000 {
			val.Level = 5
		}
		if val.Level > 0 {
			res, err := databases.Orm.Cols("level").Update(val, models.Users{Id: val.Id})
			fmt.Println(key, "----------------", val.Id, res, err)
		}
		fmt.Println(key, "--------暂无任务--------")
	}
}
