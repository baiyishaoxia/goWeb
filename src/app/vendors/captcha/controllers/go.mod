module controllers

replace (
	app/vendors/captcha/models => ../../../../app/vendors/captcha/models
	app/vendors/loger/models => ../../../../app/vendors/loger/models
	app/vendors/redis/models => ../../../../app/vendors/redis/models
	app/vendors/session/models => ../../../../app/vendors/session/models
	config => ../../../../config
	databases => ../../../../databases
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

go 1.15

require (
	app/vendors/captcha/models v0.0.0-00010101000000-000000000000
	app/vendors/redis/models v0.0.0-00010101000000-000000000000
	app/vendors/session/models v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/sessions v0.0.3 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/mojocn/base64Captcha v1.2.2
)
