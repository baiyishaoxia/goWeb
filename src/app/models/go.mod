module models

go 1.15

replace (
	app => ../../app
	app/models => ../models
	app/vendors/loger/models => ../vendors/loger/models
	app/vendors/redis/datasource => ../vendors/redis/datasource
	app/vendors/redis/models => ../vendors/redis/models
	app/vendors/session/models => ../vendors/session/models
	app/vendors/size/models => ../vendors/size/models
	config => ../../config
	databases => ../../databases
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

require (
	app v0.0.0-00010101000000-000000000000
	app/vendors/redis/models v0.0.0-00010101000000-000000000000
	app/vendors/session/models v0.0.0-00010101000000-000000000000
	databases v0.0.0-00010101000000-000000000000
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/sessions v0.0.3 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-xorm/xorm v0.7.9
	github.com/kr/pretty v0.3.1 // indirect
	github.com/smartystreets/goconvey v1.7.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
