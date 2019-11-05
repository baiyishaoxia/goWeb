package config

const (
	Limit       = 15  //每页条数
	HttpSuccess = 200 //成功
	HttpError   = 201 //错误
	VueSuccess  = 1   //成功
	VueError    = 0   //错误
	VueRelogin  = 2   //重新登录

	PARAMS_ERROR             = 10001 //参数错误
	SYSTEM_ERROR             = 10002 //系统错误
	TOO_FREQUENTLY           = 10003 //操作频繁
	QUERY_ERROR              = 10004 //请求错误
)
