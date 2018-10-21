package background

import (
	"app/models/background"
	"config"
	"databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetVideoList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(models.ReadConfig("sys.paginate"))
	data, num, all, page := models.VideoList(page-1, limit, keywords)
	//模版
	c.HTML(http.StatusOK, "video/list", gin.H{
		"Title":    "Background Login",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
	})
}
func GetVideoCreate(c *gin.Context) {
	//模版
	c.HTML(http.StatusOK, "video/create", gin.H{
		"Title": "Background Login",
	})
}
func PostVideoCreate(c *gin.Context) {
	title := c.PostForm("title")
	remark := c.PostForm("remark")
	url := c.PostForm("url")
	img_url := c.PostForm("img_url")

	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请输入标题",
		})
		return
	}
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请上传对应视频文件",
		})
		return
	}

	add := &models.Video{Title: title, Url: url, ImgUrl: img_url, Remark: remark}
	if add.AddVideo() {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "保存成功",
			"url":    "/admin/video/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "保存失败",
			"url":    "/admin/video/create",
		})
	}
}
func GetVideoEdit(c *gin.Context) {
	id := c.Param("id")
	video := models.GetVideoById(id)

	c.HTML(http.StatusOK, "video/edit", gin.H{
		"Title": "Background Login",
		"Data":  video,
	})
}

func PostVideoEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	title := c.PostForm("title")
	remark := c.PostForm("remark")
	url := c.PostForm("url")
	img_url := c.PostForm("img_url")

	if title == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "请输入标题",
		})
		return
	}
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "请上传对应视频文件",
		})
		return
	}

	edit := new(models.Video)
	edit.Title = title
	edit.Remark = remark
	edit.Url = url
	edit.ImgUrl = img_url

	_, err := databases.Orm.Cols("title", "url", "img_url", "remark").Update(edit, models.Video{Id: id})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/video/list",
	})

}
func PostVideoDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	video := new(models.Video)
	_, err := databases.Orm.In("id", ids).Delete(video)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/video/list",
	})
}
