package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)
type Map map[string]interface{}

//UtilsAction 模块统一遵循的 对外接口---------------------------------------------------------------------------
type UtilsAction interface {
	Action(appID int64, params ...interface{}) (Map, UtilsException)
}

type UtilsMiddleware interface {
	SetRequestMiddleware(func(appID int64, params ...interface{}) interface{}) UtilsMiddleware
	SetResponMiddleware(func(appID int64, mp Map, excp UtilsException, params ...interface{}) (Map, UtilsException)) UtilsMiddleware
}

//ActionStruct 上游协议层调用的接口，并实现了UtilsAction所有方法------------------------------------------------
type ActionStruct struct { 	//应用id
	requestMiddleware func(appID int64, params ...interface{}) interface{}                                      //模块包自定义方法（处理协议层请求，并返回响应结果）
	responMiddleware  func(appID int64, mp Map, excp UtilsException, params ...interface{}) (Map, UtilsException) //模块包自定义中间件方法(对返回响应进行处理)
}

//SetRequestMiddleware 设置对协议层请求处理的句柄方法（模块包）
func (that *ActionStruct) SetRequestMiddleware(requestMiddleware func(appID int64, params ...interface{}) interface{}) UtilsMiddleware {
	that.requestMiddleware = requestMiddleware
	return that
}

//SetResponMiddleware 设置对协议层响应的中间件方法
func (that *ActionStruct) SetResponMiddleware(responMiddleware func(appID int64, mp Map, excp UtilsException, params ...interface{}) (Map, UtilsException)) UtilsMiddleware {
	that.responMiddleware = responMiddleware
	return that
}
//_parseRespon 解析响应，生成UtilsAction.Action规范的返回格式
func (that *ActionStruct) _parseRespon(rs interface{}) (mp Map, excp UtilsException) {
	switch t := rs.(type) {
	case error:
		excp = ThrowException(40003001, t)
		return
	case Exception:
		excp = ThrowException(40003001, t)
		return
	case Map:
		mp = t
		return
	case map[string]interface{}:
		mp = t
		return
	case string,int,int64,float32,float64:
        string_map := make(map[string]interface{})
		err:= json.Unmarshal([]byte(`{"return_msg":"`+ ToString(t) +`"}`), &string_map)
		if err != nil {
			excp = ThrowException(40003002, errors.New(err.Error()))
		}
		mp = string_map
		return
	default:
		_jsonByte, _err := json.Marshal(t)
		if _err != nil {
			_msg := "action marshal error: " + _err.Error()
			excp = ThrowException(40003002, errors.New(_msg))
		}
		_jsonMap := make(map[string]interface{}, 0)
		if _err := json.Unmarshal([]byte(_jsonByte), &_jsonMap); _err != nil {
			_msg := "action unmarshal error: " + _err.Error()
			excp = ThrowException(40003002, errors.New(_msg))
			return
		}
		mp = _jsonMap
		return
	}
}
//Action 上游协议层最终调用的接口方法
func (that *ActionStruct) Action(appID int64, params ...interface{}) (Map, UtilsException) {
	//模块包 处理请求
	_respon := that.requestMiddleware(appID, params...)
	_mp, _excp := that._parseRespon(_respon)
	//模块包 响应处理中间件
	if that.responMiddleware != nil {
		return that.responMiddleware(appID, _mp, _excp, params...)
	}
	return _mp, _excp
}
//UtilsException 模块异常接口，Exception实现了该接口的所有方法---------------------------------------------------------
type UtilsException interface {
	GetCode() int64
	GetMessage() string
	GetTrace() []UtilsException
	GetTraceAsString() string
}
//Exception 异常
type Exception struct {
	code  int64
	msg   string
	tract []UtilsException
}
//GetCode 获取异常错误码
func (that Exception) GetCode() int64 {
	return that.code
}

//GetMessage 获取异常消息
func (that Exception) GetMessage() string {
	return that.msg
}

//GetTrace ...
func (that Exception) GetTrace() []UtilsException {
	return that.tract
}

//GetTraceAsString ...
func (that Exception) GetTraceAsString() string {
	_traceSlice := make([]string, 0)
	for _, _e := range that.tract {
		_traceSlice = append(_traceSlice, ToString(_e.GetCode())+":"+_e.GetMessage())
	}
	return strings.Join(_traceSlice, " | ")
}
//Tools----------------------------------------------------------------------------------------------------
//ThrowException 抛出异常
func ThrowException(code int64, errSlice ...interface{}) UtilsException {
	var _excp Exception
	for _i, _err := range errSlice {
		switch t := _err.(type) {
		case string:
			_excp.msg = t
			_excp.code = code
			_excp.tract = []UtilsException{
				_excp,
			}
		case error:
			_excp.msg = t.Error()
			_excp.code = code
			_excp.tract = []UtilsException{
				_excp,
			}
		case Exception:
			if _i == 0 {
				_excp = t
			} else {
				_excp.code = code
				_excp.msg = "内部异常"
				_excp.tract = append(t.tract, _excp)
			}
		}
	}
	return _excp
}

//ToString
func ToString(d interface{}) string {
	switch n := d.(type) {
	case int64:
		return strconv.FormatInt(n, 10)
	case string:
		return n
	case int:
		return strconv.Itoa(n)
	case float32:
		return strconv.FormatFloat(float64(n), 'f', 2, 64)
	case float64:
		return strconv.FormatFloat(n, 'f', 2, 64)
	case []byte:
		return string(n)
	case error:
		return n.Error()
	case bool:
		if n {
			return "1"
		}
		return "0"
	}
	return "0"
}
//响应解析，返回统一数据格式
func parse(data interface{}) map[string]interface{} {
	var (
		_code = "200"
		_msg = "SUCCESS"
		_rsData interface{}
	)
	switch _t := data.(type) {
	case UtilsException:
		_code = ToString(_t.GetCode())
		_msg = _t.GetMessage()
		_rsData = ""
	case error:
		_code = "400"
		_msg = _t.Error()
		_rsData = ""
	default:
		_rsData = data
	}
	return map[string]interface{}{
		"code":   _code,
		"msg":    _msg,
		"result": _rsData,
	}
}
//----------------------------------------------------------------------------------------------------
//设置逻辑处理层中间件对外接口
func newUtilsActon(requestMiddleware func(appID int64) interface{}) UtilsAction {
	middleware := &ActionStruct{}
	middleware.SetRequestMiddleware(func(appID int64, params ...interface{}) interface{} {
		//sudo 这里作请求..
		fmt.Println("--------------SetRequestMiddleware--------------",appID,params)
		return requestMiddleware(appID)
	})
	middleware.SetResponMiddleware(func(appID int64, rs Map, excp UtilsException, params ...interface{}) (Map, UtilsException) {
		//sudo 这里作响应..
		fmt.Println("--------------SetResponMiddleware--------------",appID , rs , excp , params)
		if excp!=nil{
			//sudo 这里作error...
			fmt.Println("--------------Exception--------------",excp.GetTraceAsString())
		}
		return rs, excp
	})
	return middleware
}
//接管函数到设置逻辑处理层
func funcDemo()  UtilsAction{
	return newUtilsActon(func(appID int64) interface{} {
		return map[string]interface{}{                        //成功处理
			"name":"tang",
			"age" :"22",
			"city":"ShenZhen",
		}
		//return  throwException(200,"错误")  //错误处理
	})
}
//处理响应异常返回
func throwException(code int64, errSlice ...interface{}) UtilsException {
	return ThrowException(code, errSlice...)
}
//Main for Example------------------------------------------------------------------------------------------
func main() {
	//连贯高级用法
	_data, _exception:=funcDemo().Action(12306,"你好")
	fmt.Println("result:",_data, _exception)
	if _exception!=nil{
		fmt.Println("json:",parse(_exception))
	}else{
		fmt.Println("json:",parse(_data))
	}
}