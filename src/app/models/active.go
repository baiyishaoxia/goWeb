package models

import (
	"app"
	"databases"
	"fmt"
	"html/template"
	"time"
)

type Active struct {
	Id             int64         `json:"id"`
	UserId         int64         `json:"user_id"`                     //主办方用户id
	Title          string        `json:"title"`                       //活动标题
	ImgUrl         string        `json:"img_url"`                     //活动图片
	StartTime      time.Time     `json:"start_time"`                  //活动开始时间
	StartSignTime  time.Time     `json:"start_sign_time"`             //报名开始时间
	EndTime        time.Time     `json:"end_time"`                    //活动结束时间
	EndSignTime    time.Time     `json:"end_sign_time"`               //报名结束时间
	Address        string        `json:"address"`                     //活动地点
	Provinces      string        `json:"provinces"`                   //省份
	City           string        `json:"city"`                        //城市
	PeopleNum      string        `json:"people_num"`                  //限额报名人数
	Company        string        `json:"company"`                     //举办方
	ActiveAbstract template.HTML `json:"active_abstract"`             //活动亮点
	ActiveContent  template.HTML `json:"active_content"`              //活动详情
	ClickNum       int64         `json:"click_num"`                   //浏览数
	GoodNum        int64         `json:"good_num"`                    //点赞数
	CollectNum     int64         `json:"collect_num"`                 //收藏数
	IsShow         string        `json:"is_show"`                     //是否显示
	SignNum        int64         `json:"sign_num"`                    //当前报名人数
	Sort           int64         `json:"sort"`                        //排序
	ActiveImagesId int64         `json:"active_images_id"`            //活动海报
	BannerType     int64         `json:"banner_type"`                 //banner类型(1图库 2自定义)
	Type           int64         `json:"type1"`                       //所属类型
	IsOpen         bool          `json:"not null default true BOOL"`  //是否公开
	IsOver         string        `json:"is_over"`                     //活动是否结束
	IsSubmit       int64         `json:"is_submit"`                   //提交状态(1草稿 2已发布 3已取消)
	CheckStatus    int64         `json:"check_status"`                //审核状态
	InputPeopleNum int64         `json:"input_people_num"`            //录入人数
	FreeNum        int64         `json:"free_num"`                    //免费票张数
	ChargeNum      int64         `json:"charge_num"`                  //付费票张数
	Status         int64         `json:"status"`                      //活动状态
	CheckTime      time.Time     `json:"check_time"`                  //活动状态
	IsRed          bool          `json:"not null default false BOOL"` //是否推荐  false不推荐  true推荐
	CancelComment  string        `json:"cancel_comment"`              //取消原因
	CreatedAt      app.Time      `xorm:"created" json:"created_at"`
	UpdatedAt      app.Time      `xorm:"updated" json:"-"`
}

//region Remark:按照活动id获取一条记录  Author:ld
func GetActiveById(id string) *Active {
	active := new(Active)
	has, err := databases.Orm.Where("id=?", id).Get(active)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return active
}

//endregion
