package home

import (
	"app"
	"app/channel"
	"app/models/home"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//region   活动抽奖   Author:tang
func PostActiveSign(c *gin.Context) {
	db := databases.Orm.NewSession()
	err := db.Begin()
	if err != nil {
		fmt.Println("-----立即抽奖-----", err.Error())
	}
	//----------------处理逻辑代码 start----------------
	db, res1, trading, info := models.CreateActiveTrading(c, db)
	if res1 == false {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"data":   info, //错误信息
		})
		return
	}
	if res1 {
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status":   config.HttpSuccess,
			"data":     "成功",
			"order_id": trading.Id,
			"money":    0,
		})
		return
	}
	//----------------处理逻辑代码 end----------------
	res4 := false
	pay_uuid := app.Guid()
	trade_no := trading.OrderNumber
	// start-生成支付订单号
	order_no := time.Now().Format("20060102150405") + strconv.FormatInt(trading.UserId, 10)
	// db提交成功后查询数据信息
	activeId := strconv.Itoa(int(trading.ActiveId))
	activeInfo := models.GetActiveById(activeId)
	//使用管道处理数据
	channel.UserAmountActiveOrderChan <- &models.UserAmountActiveOrder{
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
		Status:       0,
		Uuid:         pay_uuid,
	}
	//获取结果
	msg := channel.HandleResult(pay_uuid)
	fmt.Println(msg)
	if msg.Status == 1 {
		res4 = true
	}
	if res1 && res4 {
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status":   config.HttpSuccess,
			"data":     "成功",
			"order_id": trading.Id,
			"payment":  trading.Payment,
			"money":    trading.Money,
		})
		return
	} else {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"data":   "网络错误",
		})
		return
	}
}
