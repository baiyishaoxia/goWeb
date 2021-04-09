module databases

go 1.15

replace (
	app/vendors/loger/models => ../app/vendors/loger/models
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
)

require (
	app/vendors/loger/models v0.0.0-00010101000000-000000000000
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/garyburd/redigo v1.6.2
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-ini/ini v1.62.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-xorm/core v0.6.3
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/lib/pq v1.10.0
)
