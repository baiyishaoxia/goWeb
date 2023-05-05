module blog

go 1.15

replace (
	app => ../../../../app
	app/controllers => ../../../../app/controllers
	app/models => ../../../../app/models
	app/service/background => ../../../../app/service/background
	app/service/common => ../../../../app/service/common
	app/service/home => ../../../../app/service/home
	app/vendors/loger/models => ../../../../app/vendors/loger/models
	app/vendors/redis/models => ../../../../app/vendors/redis/models
	app/vendors/session/models => ../../../../app/vendors/session/models
	app/vendors/size/models => ../../../../app/vendors/size/models
	config => ../../../../config
	databases => ../../../../databases
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

require (
	app v0.0.0-00010101000000-000000000000
	app/models v0.0.0-00010101000000-000000000000
	app/service/background v0.0.0-00010101000000-000000000000
	app/service/common v0.0.0-00010101000000-000000000000
	app/service/home v0.0.0-00010101000000-000000000000
	app/vendors/session/models v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	databases v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.9.0
	github.com/smartystreets/goconvey v1.8.0 // indirect
	github.com/zcshan/d3outh v0.0.0-20201222010721-a8e886c23105
	gopkg.in/ini.v1 v1.67.0 // indirect
)
