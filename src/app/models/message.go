package models

import (
	"databases"
	"fmt"
	"math"
	"time"
)

type Message struct {
	Id            int64     `xorm:"not null pk autoincr unique INTEGER" json:"id"`
	MessageCateId int64     `xorm:"BIGINT"       json:"message_cate_id"`       //类型(1留言，2评论)
	UsersId       int64     `xorm:"BIGINT"       json:"users_id"`              //用户ID
	ParentId      int64     `xorm:"BIGINT"       json:"parent_id"`             //回复留言ID
	ArticleId     int64     `xorm:"BIGINT"      json:"article_id"`             //文章ID
	Address       string    `xorm:"VARCHAR(255)" json:"address"`               //地点
	Content       string    `xorm:"TEXT"    json:"content"`                    //回复内容
	IsShow        bool      `xorm:"not null default true BOOL" json:"is_show"` //是否显示
	CreatedAt     time.Time `xorm:"created" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated" json:"updated_at"`
}
type MessageUser struct {
	Message `xorm:"extends"`
	Users   `xorm:"extends"`
}

//评论列表
func GetMessageList(page int, limit int, keywords string) ([]*MessageUser, float64, float64, int) {
	var item = make([]*MessageUser, 0)
	err := databases.Orm.Table("message").Alias("mm").
		Join("LEFT", []string{"users", "u"}, "mm.users_id = u.id").Desc("mm.id")
	if keywords != "" {
		err.Where("mm.content like ?", "%"+keywords+"%")
	}
	err1 := *err
	//记录数
	num, err3 := err1.Table("message").Count()
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
	err2 := err.Limit(limit, page*limit).Find(&item)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return item, float64(num), all, page + 1
}
