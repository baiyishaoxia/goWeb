module routers

go 1.15

replace (
	app => ../app
	app/channel => ../app/channel
	app/channel/chat => ../app/channel/chat
	app/controllers => ../app/controllers
	app/controllers/background => ../app/controllers/background
	app/controllers/home => ../app/controllers/home
	app/controllers/home/blog => ../app/controllers/home/blog
	app/controllers/home/chat => ../app/controllers/home/chat
	app/models => ../app/models
	app/service/background => ../app/service/background
	app/service/common => ../app/service/common
	app/service/home => ../app/service/home
	app/vendors/captcha/controllers => ../app/vendors/captcha/controllers
	app/vendors/captcha/models => ../app/vendors/captcha/models
	app/vendors/loger/models => ../app/vendors/loger/models
	app/vendors/redis/datasource => ../app/vendors/redis/datasource
	app/vendors/redis/models => ../app/vendors/redis/models
	app/vendors/session/models => ../app/vendors/session/models
	app/vendors/size/models => ../app/vendors/size/models
	config => ../config
	databases => ../databases
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
	statistical => ../statistical
)

require (
	app v0.0.0-00010101000000-000000000000
	app/channel v0.0.0-00010101000000-000000000000 // indirect
	app/channel/chat v0.0.0-00010101000000-000000000000
	app/controllers/background v0.0.0-00010101000000-000000000000
	app/controllers/home v0.0.0-00010101000000-000000000000
	app/controllers/home/blog v0.0.0-00010101000000-000000000000
	app/controllers/home/chat v0.0.0-00010101000000-000000000000
	app/service/common v0.0.0-00010101000000-000000000000
	app/vendors/captcha/controllers v0.0.0-00010101000000-000000000000
	github.com/foolin/gin-template v0.0.0-20190415034731-41efedfb393b
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.9.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/mojocn/base64Captcha v1.3.1 // indirect
	github.com/smartystreets/goconvey v1.8.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
