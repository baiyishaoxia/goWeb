module src

go 1.15

replace (
	app => ./app
	app/channel => ./app/channel
	app/channel/chat => ./app/channel/chat
	app/channel/job => ./app/channel/job
	app/controllers => ./app/controllers
	app/controllers/background => ./app/controllers/background
	app/controllers/home => ./app/controllers/home
	app/controllers/home/blog => ./app/controllers/home/blog
	app/controllers/home/chat => ./app/controllers/home/chat
	app/grpc => ./app/grpc
	app/models => ./app/models
	app/service/background => ./app/service/background
	app/service/common => ./app/service/common
	app/service/home => ./app/service/home
	app/vendors/captcha/controllers => ./app/vendors/captcha/controllers
	app/vendors/captcha/models => ./app/vendors/captcha/models
	app/vendors/loger/models => ./app/vendors/loger/models
	app/vendors/redis/datasource => ./app/vendors/redis/datasource
	app/vendors/redis/models => ./app/vendors/redis/models
	app/vendors/session/models => ./app/vendors/session/models
	app/vendors/size/models => ./app/vendors/size/models
	config => ./config
	databases => ./databases
	github.com/mojocn/base64Captcha v1.3.1 => github.com/mojocn/base64Captcha v1.2.2
	other/reading/yournovel/conf => ./other/reading/yournovel/conf
	other/reading/yournovel/db/redis => ./other/reading/yournovel/db/redis
	other/reading/yournovel/fetcher => ./other/reading/yournovel/fetcher
	other/reading/yournovel/http => ./other/reading/yournovel/http
	other/reading/yournovel/middleware => ./other/reading/yournovel/middleware
	other/reading/yournovel/model => ./other/reading/yournovel/model
	other/reading/yournovel/routers => ./other/reading/yournovel/routers
	other/reading/yournovel/service/novel => ./other/reading/yournovel/service/novel
	other/reading/yournovel/service/searchengine => ./other/reading/yournovel/service/searchengine
	other/reading/yournovel/tool => ./other/reading/yournovel/tool
	routers => ./routers
	statistical => ./statistical
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

require (
	app v0.0.0-00010101000000-000000000000
	app/channel v0.0.0-00010101000000-000000000000
	app/channel/chat v0.0.0-00010101000000-000000000000
	app/channel/job v0.0.0-00010101000000-000000000000
	app/controllers/home/chat v0.0.0-00010101000000-000000000000
	app/grpc v0.0.0-00010101000000-000000000000
	app/models v0.0.0-00010101000000-000000000000
	databases v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.6.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.5 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-xorm/xorm v0.7.9
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/ini.v1 v1.62.0 // indirect
	other/reading/yournovel/db/redis v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/fetcher v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/http v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/middleware v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/model v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/routers v0.0.0-00010101000000-000000000000
	other/reading/yournovel/service/novel v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/service/searchengine v0.0.0-00010101000000-000000000000 // indirect
	other/reading/yournovel/tool v0.0.0-00010101000000-000000000000 // indirect
	routers v0.0.0-00010101000000-000000000000
	statistical v0.0.0-00010101000000-000000000000 // indirect
)
