package models

import (
	"app"
	"time"
)

type UserAmountActiveOrder struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`        // 用户id
	ActiveUserId int64     `json:"active_user_id"` // 活动创建用户id
	ActiveId     int64     `json:"active_id"`      // 活动id
	OrderNo      string    `json:"order_no"`       // 平台交易号，平台展示使用
	OrderNumber  string    `json:"order_number"`   // 创业ruiec_active_trading（order_number）订单号
	TradeNo      string    `json:"trade_no"`       // 支付交易号，需给第三方支付接口交易号
	PayTradeNo   string    `json:"pay_trade_no"`   // 支付平台返回交易号，payment为余额支付则为空值
	Uuid         string    `json:"uuid"`           // 唯一标识
	Amount       float64   `json:"amount"`         // 订单金额
	Payment      int64     `json:"payment"`        // 支付方式(1支付宝 2微信3余额)
	PayClass     int64     `json:"pay_class"`      // 支付类型
	Status       int64     `json:"status"`         // 0：待付款；1：支付成功；2：支付失败
	Comment      string    `json:"comment"`        // 文字描述支付说明
	SuccessAt    time.Time `json:"success_at"`     // 支付成功回调时间
	NotifyAt     time.Time `json:"notify_at"`      // 支付通知的发送时间
	GmtPaymentAt time.Time `json:"gmt_payment_at"` // 该笔交易的买家付款时间
	CreatedAt    app.Time  `xorm:"created" json:"created_at"`
	UpdatedAt    app.Time  `xorm:"updated" json:"updated_at"`
}
