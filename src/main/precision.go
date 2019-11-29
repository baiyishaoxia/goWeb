package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

//精度计算
func main()  {
	var sss int64 =1011911131432200196
	ss:= struct {
		OrderNo           interface{} `json:"order_no"`           //订单号（要保证唯一）
		TransactionNo     interface{} `json:"transaction_no"`     //交易号
		PaymentWay        interface{} `json:"payment_way"`        //支付方式：6微信，7支付宝
		Price             interface{} `json:"price"`              //到帐金额，单位：分
		TransactionStatus interface{} `json:"transaction_status"` //交易状态
		TransactionCode   interface{} `json:"transaction_code"`   //交易错误码
		TransactionMsg    interface{} `json:"transaction_msg"`    //交易提示
	}{
		OrderNo:           sss,

	}
	fmt.Println(ss)


	var (
		price = 1000.02  //价格
		num = 20   //数量
	)
    //设置精度为2位, 四舍五入的精度
	decimal.DivisionPrecision = 2

	//相加 float和int相加
	d1 := decimal.NewFromFloat(price).Add(decimal.NewFromFloat(float64(num)))
	fmt.Println(d1) // output: "1020.02"

	//相减
	d2 := decimal.NewFromFloat(price).Sub(decimal.NewFromFloat(float64(num)))
	fmt.Println(d2) // output: "980.02"

    //相乘
	d3 := decimal.NewFromFloat(price).Mul(decimal.NewFromFloat(float64(num)))
	fmt.Println(d3) // output: "20000.4"

	//相除
	d4 := decimal.NewFromFloat(float64(price)).Div(decimal.NewFromFloat(float64(num)))
	fmt.Println(d4) // output: "50"
}
