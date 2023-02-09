module app

go 1.15

replace (
	app/vendors/size/models => ./vendors/size/models
	config => ../config
)

require (
	app/vendors/size/models v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	github.com/StackExchange/wmi v1.2.1
	github.com/garyburd/redigo v1.6.4
	github.com/gin-gonic/gin v1.7.7
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
)
