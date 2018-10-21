package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//region Remark:折线图 Author:tang
func GetChartsZx(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/zx", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:时间轴折线图 Author:tang
func GetChartsSj(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/sj", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:区域图 Author:tang
func GetChartsQy(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/qy", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:柱状图 Author:tang
func GetChartsZz(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/zz", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:饼状图 Author:tang
func GetChartsBz(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/bz", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:3D饼状图 Author:tang
func GetCharts3Dbz(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/3Dbz", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:3D柱状图 Author:tang
func GetCharts3Dzz(c *gin.Context) {
	c.HTML(http.StatusOK, "charts/3Dzz", gin.H{
		"Title": "Background Index",
	})
}

//endregion
