package models

import (
	"app"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"strconv"
	"time"
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

func CreateActiveTrading(c *gin.Context, db *xorm.Session) (*xorm.Session, bool, *ActiveTrading, string) {
	token := c.Request.Header.Get("token")
	user_id := GetUserByToken(token).Id
	active_id, _ := strconv.ParseInt(c.PostForm("active_id"), 10, 64)
	ticket_id, _ := strconv.ParseInt(c.PostForm("ticket_id"), 10, 64)
	order_number := time.Now().Format("20060102150405") + strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	add := &ActiveTrading{
		UserId: user_id, ActiveId: active_id, TicketId: ticket_id, Number: 1, Money: 0.01, Price: 0.01,
		Check: 1, Status: 1, Payment: 1, OrderNumber: order_number}
	res, err := db.Insert(add)
	if res < 1 {
		return db, false, add, "订单数据生成失败"
	}
	if err != nil {
		fmt.Println(err.Error())
		return db, false, add, "订单数据生成失败"
	}
	return db, true, add, "订单数据生成成功"
}
