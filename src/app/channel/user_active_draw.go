package channel

import (
	"app/models/home"
	"app/vendors/redis/datasource"
	"databases"
	"errors"
)

//region Remark:活动抽奖支付订单增加 Author:tang
func AddUserAmountActiveOrder(userAmountActiveOrder *models.UserAmountActiveOrder) bool {
	db := databases.Orm.NewSession()
	db.Begin()
	//region Remark:添加支付订单
	userAmountActiveOrder.Status = 0 // 支付平台未请求，都是等于0：待付款
	res, err := db.Insert(userAmountActiveOrder)
	if res == 0 || err != nil {
		err = errors.New("活动订单插入超时")
		datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "status", -1)
		datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "msg", err)
		db.Rollback()
		return false
	}
	db.Commit()
	datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "status", 1)
	datasource.RedisPool.HSet(userAmountActiveOrder.Uuid, "msg", "提交成功")
	return true
	//endregion
}

//endregion
