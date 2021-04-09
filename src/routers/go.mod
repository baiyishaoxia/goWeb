module routers

go 1.15

replace (
	app => ../app
	app/models => ../app/models
	app/channel => ../app/channel
	app/channel/chat => ../app/channel/chat
	app/controllers => ../app/controllers
	app/controllers/background => ../app/controllers/background
	app/controllers/home => ../app/controllers/home
	app/controllers/home/blog => ../app/controllers/home/blog
	app/controllers/home/chat => ../app/controllers/home/chat
	app/service/common => ../app/service/common
	app/vendors/captcha/controllers => ../app/vendors/captcha/controllers
	app/vendors/size/models => ../app/vendors/size/models
	app/vendors/redis/models => ../app/vendors/redis/models
	app/vendors/redis/datasource => ../app/vendors/redis/datasource
	app/vendors/session/models => ../app/vendors/session/models
	config => ../config
	databases => ../databases
	statistical => ../statistical
	app/service/background => ../app/service/background
	app/service/home => ../app/service/home
	app/vendors/captcha/models => ../app/vendors/captcha/models
	app/vendors/loger/models => ../app/vendors/loger/models
	app/vendors/size/models => ../app/vendors/size/models
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

require (
	app v0.0.0-00010101000000-000000000000
	app/channel v0.0.0-00010101000000-000000000000 // indirect
	app/channel/chat v0.0.0-00010101000000-000000000000
	app/controllers v0.0.0-00010101000000-000000000000 // indirect
	app/controllers/background v0.0.0-00010101000000-000000000000
	app/controllers/home v0.0.0-00010101000000-000000000000
	app/controllers/home/blog v0.0.0-00010101000000-000000000000
	app/controllers/home/chat v0.0.0-00010101000000-000000000000
	app/service/common v0.0.0-00010101000000-000000000000
	app/vendors/captcha/controllers v0.0.0-00010101000000-000000000000
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/garyburd/redigo v1.6.2 // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/mojocn/base64Captcha v1.3.1 // indirect
	github.com/zcshan/d3outh v0.0.0-20201222010721-a8e886c23105 // indirect
)
