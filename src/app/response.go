package app

import (
	"config"
)

//数据data的基本属性
type Response struct {
	Errors []string    `json:"errors,omitempty"`
	Time   string      `json:"time"`
	List   interface{} `json:"list"`
}

//返回json的统一格式
type JSONS struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   Response `json:"data"`
}

//响应体
type Responser struct {
	JSONS
}

//状态码200、提示信息，数据
func (r *Responser) Success(status int, msg string, data interface{}) JSONS {
	r.Status = status
	r.Msg = msg
	r.Data.List = data
	r.Data.Time = DateFormat("Y-m-d H:i:s")
	return r.JSONS
}

//状态码201，提示信息
func (r *Responser) Error(err int, message string, tr ...bool) JSONS {
	r.Status = err
	if len(tr) > 0 && tr[0] {
		r.Msg = message
	} else {
		r.Msg = message
	}
	r.Data.Time = DateFormat("Y-m-d H:i:s")
	r.Data.List = new(struct{})
	return r.JSONS
}

//抛出异常
func (res *JSONS) parseErrors(errors ...string) {
	if len(errors) > 0 {
		res.Data.Errors = errors
	}
}

func (r *Responser) InvalidArgument(errors ...string) JSONS {
	res := r.Error(config.PARAMS_ERROR, "params error", true)
	res.parseErrors(errors...)
	return res
}

func (r *Responser) SystemError(errors ...string) JSONS {
	res := r.Error(config.SYSTEM_ERROR, "system error", true)
	res.parseErrors(errors...)
	return res
}

func (r *Responser) QueryError(errors ...string) JSONS {
	res := r.Error(config.QUERY_ERROR, "query error", true)
	res.parseErrors(errors...)
	return res
}
