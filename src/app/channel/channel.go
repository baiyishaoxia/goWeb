package channel

import (
	"app/vendors/redis/datasource"
)

//region Remark:处理并发问题 Author:tang
func HandleConcurrent() {
	for {
		select {
		case userAmountActiveOrder := <-UserAmountActiveOrderChan:
			//添加活动抽奖支付订单
			if AddUserAmountActiveOrder(userAmountActiveOrder) == false {
				break
			}
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
