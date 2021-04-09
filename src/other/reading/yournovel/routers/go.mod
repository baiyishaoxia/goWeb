module routers

go 1.15

replace (
	other/reading/yournovel/conf => ../conf
	other/reading/yournovel/db/redis => ../db/redis
	other/reading/yournovel/http => ../http
	other/reading/yournovel/middleware => ../middleware
	other/reading/yournovel/model => ../model
	other/reading/yournovel/service/novel => ../service/novel
	other/reading/yournovel/service/searchengine => ../service/searchengine
	other/reading/yournovel/tool => ../tool
)

require (
	github.com/gin-gonic/gin v1.7.1
	other/reading/yournovel/conf v0.0.0-00010101000000-000000000000
	other/reading/yournovel/db/redis v0.0.0-00010101000000-000000000000
	other/reading/yournovel/http v0.0.0-00010101000000-000000000000
	other/reading/yournovel/middleware v0.0.0-00010101000000-000000000000
)
