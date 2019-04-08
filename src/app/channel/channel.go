package channel

import (
	"app/models/home"
	"app/vendors/redis/datasource"
	"databases"
)

var (
	// 活动订单生成管道
	UserAmountActiveOrderChan = make(chan *models.UserAmountActiveOrder, 1000000)
)

//region Remark:处理并发问题 Author:tang
func HandleConcurrent() {
	for {
		select {
		case userAmountActiveOrder := <-UserAmountActiveOrderChan:
			//region Remark:添加活动支付订单
			db := databases.Orm.NewSession()
			db.Begin()
			res, err := AddUserAmountActiveOrder(userAmountActiveOrder, db)
			if res == 0 || err != nil {
				datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "status", -1)
				datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "msg", err)
				db.Rollback()
				break
			}
			datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "status", 1)
			datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "msg", "成功")
			db.Commit()
			//endregion
		}
	}
}

//endregion

//region Remark:处理结果的返回
type HandleResultStruck struct {
	Status  int64
	Message string
}

func HandleResult(uuid string) (handleResult HandleResultStruck) {
	for {
		statusRedis := datasource.RedisPool.HGet(uuid, "status")
		i, _ := statusRedis.Int64()
		msgRedis, _ := datasource.RedisPool.HGet(uuid, "msg").Bytes()
		msg := string(msgRedis)
		if i == -1 {
			handleResult.Status = i
			handleResult.Message = msg
			datasource.RedisPool.HDel(uuid, "status", "msg")
			return
		}
		if i == 1 {
			handleResult.Status = i
			handleResult.Message = msg
			datasource.RedisPool.HDel(uuid, "status", "msg")
			return
		}
	}
}

//endregion
