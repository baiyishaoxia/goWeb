module background

go 1.15

replace (
	app => ../../../app
	app/models => ../../../app/models
	app/service/background => ../../../app/service/background
	app/vendors/loger/models => ../../../app/vendors/loger/models
	app/vendors/redis/models => ../../../app/vendors/redis/models
	app/vendors/session/models => ../../../app/vendors/session/models
	app/vendors/size/models => ../../../app/vendors/size/models
	config => ../../../config
	databases => ../../../databases
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
	statistical => ../../../statistical
)

require (
	app v0.0.0-00010101000000-000000000000
	app/models v0.0.0-00010101000000-000000000000
	app/service/background v0.0.0-00010101000000-000000000000
	app/vendors/redis/models v0.0.0-00010101000000-000000000000
	app/vendors/session/models v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	databases v0.0.0-00010101000000-000000000000
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/gin v1.7.1
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	statistical v0.0.0-00010101000000-000000000000
)
