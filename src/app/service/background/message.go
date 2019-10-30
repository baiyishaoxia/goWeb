package background

import (
	"app/models"
	"databases"
	"fmt"
)

//region   定义评论类型   Author:tang
func GetMessageCate() []string {
	data := []string{1: "留言墙", 2: "文章评论"}
	return data
}

//endregion

//region   添加留言   Author:tang
func InsertMessage(self *models.Message) (bool, map[string]interface{}) {
	has, err := databases.Orm.Insert(self)
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}
	if has < 1 {
		return false, nil
	}
	data := make(map[string]interface{})
	data["id"] = self.Id
	data["message_cate_id"] = self.MessageCateId
	data["users_id"] = self.UsersId
	data["parent_id"] = self.ParentId
	data["address"] = self.Address
	data["content"] = self.Content
	if self.ParentId != 0 {
		data["replay_name"] = models.GetUserById(GetMessageById(self.ParentId).UsersId).Name
	} else {
		data["replay_name"] = "My"
	}
	data["created_at"] = self.CreatedAt.Format("2006-01-02 15:04:05")
	data["updated_at"] = self.UpdatedAt.Format("2006-01-02 15:04:05")
	return true, data
}

//endregion

//region   获取首页最新留言墙   Author:tang
func GetMessageNew() []map[string]interface{} {
	message := make([]*models.Message, 0)
	databases.Orm.Where("is_show=?", true).Desc("id").Find(&message)
	data := messageInIt(message)
	return data
}

//endregion

//region   获取首页热评用户   Author:tang
func GetMessageHot(limit int) []map[string]interface{} {
	item := make([]models.Users, 0)
	databases.Orm.OrderBy("hot_count desc").Asc("id").Select("head_img,name,hot_count,level").Limit(limit).Find(&item)
	data := make([]map[string]interface{}, len(item))
	for key, val := range item {
		data[key] = make(map[string]interface{})
		data[key]["name"] = val.Name
		data[key]["head_img"] = val.HeadImg
		data[key]["hot_count"] = val.HotCount
		data[key]["level"] = val.Level
		data[key]["level_name"] = models.GetLevelUpdate(val.Level)
	}
	return data
}

//endregion

//region   获取留言墙数据、文章评论数据   Author:tang
func GetMessageListApi(wheres map[string]interface{}) []map[string]interface{} {
	message := make([]*models.Message, 0)
	if wheres["article_id"].(int64) != 0 {
		databases.Orm.Where("message_cate_id=? and article_id=?", wheres["key"], wheres["article_id"]).Where("is_show=?", true).Where("parent_id=?", 0).Find(&message)
	} else {
		databases.Orm.Where("message_cate_id=?", wheres["key"]).Where("is_show=?", true).Where("parent_id=?", 0).Find(&message)
	}
	data := messageInIt(message)
	for k, v := range data {
		item := make([]*models.Message, 0)
		databases.Orm.Where("parent_id=?", v["id"]).Where("is_show=?", true).Find(&item)
		child := messageInIt(item)
		data[k]["child"] = child
	}
	return data
}
func messageInIt(mm []*models.Message) []map[string]interface{} {
	data := make([]map[string]interface{}, len(mm))
	for key, val := range mm {
		user := models.GetUserById(val.UsersId)
		data[key] = make(map[string]interface{})
		data[key]["id"] = val.Id
		data[key]["head_img"] = user.HeadImg
		data[key]["name"] = user.Name
		if val.ParentId != 0 {
			data[key]["replay_name"] = models.GetUserById(GetMessageById(val.ParentId).UsersId).Name
		} else {
			data[key]["replay_name"] = "My"
		}
		data[key]["message_cate_id"] = val.MessageCateId
		data[key]["users_id"] = val.UsersId
		data[key]["parent_id"] = val.ParentId
		data[key]["address"] = val.Address
		data[key]["content"] = val.Content
		data[key]["is_show"] = val.IsShow
		data[key]["created_at"] = val.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return data
}
func GetMessageById(id int64) *models.Message {
	data := new(models.Message)
	databases.Orm.Where("id=?", id).Get(data)
	return data
}

//endregion
