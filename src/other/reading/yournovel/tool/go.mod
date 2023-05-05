module tool

go 1.15

replace other/reading/yournovel/conf => ../conf

require (
	github.com/gin-gonic/gin v1.9.0
	other/reading/yournovel/conf v0.0.0-00010101000000-000000000000
)
