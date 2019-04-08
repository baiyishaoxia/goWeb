package channel

import (
	"app/models/home"
	"errors"
	"github.com/go-xorm/xorm"
)

//region Remark:创业活动支付订单增加 Author:tang
func AddUserAmountActiveOrder(userAmountActiveOrder *models.UserAmountActiveOrder, db *xorm.Session) (res int64, err error) {
	// 添加支付订单
	userAmountActiveOrder.Status = 0 // 支付平台未请求，都是=0：待付款
	res, err = db.Insert(userAmountActiveOrder)
	if res == 0 || err != nil {
		err = errors.New("创业活动订单插入超时")
		return
	}
	return
}

//endregion
