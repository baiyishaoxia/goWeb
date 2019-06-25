package models

import (
	"app"
)

type ActiveTrading struct {
	Id           int64    `json:"id"`
	UserId       int64    `json:"user_id"`       //用户id
	ActiveId     int64    `json:"active_id"`     //活动id
	TicketId     int64    `json:"ticket_id"`     //票种id
	Number       int64    `json:"number"`        //单次数量
	Price        float64  `json:"price"`         //单价
	Money        float64  `json:"money"`         //总计金额
	Status       int64    `json:"status"`        //支付状态(1未支付 2已支付3申请退款4已退款5已取消6已拒绝)
	Check        int64    `json:"check"`         //审核状态(1未审核 2已拒绝 3已通过)
	Payment      int64    `json:"payment"`       //支付方式(1支付宝 2微信 3余额)
	Comment      string   `json:"comment"`       //设置备注信息
	CheckComment string   `json:"check_comment"` //审核原因
	OrderNumber  string   `json:"order_number"`  //订单号
	CreatedAt    app.Time `xorm:"created" json:"created_at"`
	UpdatedAt    app.Time `xorm:"updated" json:"updated_at"`
}
