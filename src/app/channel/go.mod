module channel

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
	app/models v0.0.0-00010101000000-000000000000
	app/vendors/redis/datasource v0.0.0-00010101000000-000000000000
	app/vendors/redis/models v0.0.0-00010101000000-000000000000 // indirect
	app/vendors/session/models v0.0.0-00010101000000-000000000000 // indirect
	config v0.0.0-00010101000000-000000000000
	databases v0.0.0-00010101000000-000000000000
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/gin-contrib/sessions v0.0.3 // indirect
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)
