module common

go 1.15

replace (
	app => ../../../app
	app/controllers => ../../../app/controllers
	app/vendors/session/models => ../../../app/vendors/session/models
	app/vendors/size/models => ../../../app/vendors/size/models
	config => ../../../config
	databases => ../../../databases
)

require (
	app v0.0.0-00010101000000-000000000000
	app/controllers v0.0.0-00010101000000-000000000000
	app/vendors/session/models v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/sessions v0.0.3 // indirect
	github.com/gin-gonic/gin v1.7.1
)
