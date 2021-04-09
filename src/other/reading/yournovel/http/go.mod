module http

go 1.15

replace (
	other/reading/yournovel/conf => ../../yournovel/conf
	other/reading/yournovel/db/redis => ../../yournovel/db/redis
	other/reading/yournovel/model => ../../yournovel/model
	other/reading/yournovel/service/novel => ../../yournovel/service/novel
	other/reading/yournovel/service/searchengine => ../../yournovel/service/searchengine
	other/reading/yournovel/tool => ../../yournovel/tool
)

require (
	github.com/PuerkitoBio/goquery v1.6.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.5 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	other/reading/yournovel/conf v0.0.0-00010101000000-000000000000
	other/reading/yournovel/db/redis v0.0.0-00010101000000-000000000000
	other/reading/yournovel/model v0.0.0-00010101000000-000000000000
	other/reading/yournovel/service/novel v0.0.0-00010101000000-000000000000
	other/reading/yournovel/service/searchengine v0.0.0-00010101000000-000000000000
	other/reading/yournovel/tool v0.0.0-00010101000000-000000000000
)
