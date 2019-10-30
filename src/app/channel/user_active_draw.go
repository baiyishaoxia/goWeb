package channel

import (
	"app"
	"app/models"
	"app/vendors/redis/datasource"
	"config"
	"databases"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var (
	// 活动抽奖订单生成管道
	UserAmountActiveOrderChan = make(chan *ActiveRewardStruck, 1000000)
)

type ActiveRewardStruck struct {
	UserId   int64
	Uuid     string
	ActiveId int64
	TicketId int64
}

//region Remark:活动抽奖 Author:tang
func AddUserAmountActiveOrder(data *ActiveRewardStruck) bool {
	var (
		userAmountActiveOrder = new(models.UserAmountActiveOrder)
	)
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println("-----立即抽奖-----", err.Error())
	}
	//----------------处理逻辑代码 start----------------
	user_id := data.UserId
	active_id := data.ActiveId
	ticket_id := data.TicketId
	order_number := time.Now().Format("20060102150405") + strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	add := &models.ActiveTrading{
		UserId: user_id, ActiveId: active_id, TicketId: ticket_id, Number: 1, Money: 0.01, Price: 0.01,
		Check: 1, Status: 1, Payment: 1, OrderNumber: order_number}
	res, err := db.Insert(add)
	if res < 1 || err != nil {
		db.Rollback()
		datasource.RedisPool.HSet(data.Uuid, "status", -1)
		datasource.RedisPool.HSet(data.Uuid, "msg", "订单数据生成失败")
		return false
	}

	trading := add //第一条确认

	//----------------处理逻辑代码 end----------------
	pay_uuid := app.Guid()
	trade_no := trading.OrderNumber
	// start-生成支付订单号
	order_no := time.Now().Format("20060102150405") + strconv.FormatInt(trading.UserId, 10)
	// db提交成功后查询数据信息
	activeId := strconv.Itoa(int(trading.ActiveId))
	activeInfo := models.GetActiveById(activeId)
	//支付订单增加---添加支付订单
	userAmountActiveOrder = &models.UserAmountActiveOrder{
		UserId:       trading.UserId,
		ActiveUserId: activeInfo.UserId,
		ActiveId:     trading.ActiveId,
		OrderNo:      order_no,
		OrderNumber:  trading.OrderNumber,
		TradeNo:      trade_no,
		PayClass:     1,
		Payment:      trading.Payment, // 支付方式(1支付宝 2微信3余额)
		Comment:      "填写留言信息",
		Amount:       trading.Money,
		Status:       0, //支付平台未请求，都是等于0：待付款
		Uuid:         pay_uuid,
	}
	res, err = db.Insert(userAmountActiveOrder)
	if err != nil || res < 1 {
		db.Rollback()
		datasource.RedisPool.HSet(data.Uuid, "status", -1)
		datasource.RedisPool.HSet(data.Uuid, "msg", "抽奖失败")
		return false
	} else {
		msg := gin.H{
			"status":  config.HttpSuccess,
			"data":    "已中奖",
			"orderid": trading.Id,
			"money":   0,
		}
		redisRes := datasource.RedisPool.HSet(data.Uuid, "status", 1)
		redisResData, _ := redisRes.Result()
		if redisResData == false {
			db.Rollback()
			datasource.RedisPool.HSet(data.Uuid, "status", -1)
			datasource.RedisPool.HSet(data.Uuid, "msg", "提交失败")
			return false
		}
		msgStr, _ := json.Marshal(msg)
		redisRes1 := datasource.RedisPool.HSet(data.Uuid, "msg", string(msgStr))
		redisResData1, _ := redisRes1.Result()
		if redisResData1 == false {
			db.Rollback()
			datasource.RedisPool.HSet(data.Uuid, "status", -1)
			datasource.RedisPool.HSet(data.Uuid, "msg", "提交失败")
			return false
		}
		info, _ := datasource.RedisPool.HGet(data.Uuid, "msg").Result()
		if info == "" {
			db.Rollback()
			datasource.RedisPool.HSet(data.Uuid, "status", -1)
			datasource.RedisPool.HSet(data.Uuid, "msg", "提交失败")
			return false
		}
		db.Commit()
		return true
	}
}

//endregion
