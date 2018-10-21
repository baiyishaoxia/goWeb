package controllers

import (
	directory "app/vendors/directory/models"
	loger "app/vendors/loger/models"
	size "app/vendors/size/models"
	"config"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func PostUpLoadImg(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "images", "bmp,gif,jpg,jpeg,jpe,png"))
}
func PostUpLoadFile(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "file", "zip,rar,pdf,apk"))
}
func PostUpLoadVideo(c *gin.Context) {
	c.JSON(http.StatusOK, upload(c, "video", "mp4"))
}
func upload(c *gin.Context, filetype string, suffix string) gin.H {
	//得到上传的文件
	file, header, err := c.Request.FormFile("Filedata") //image这个是uplaodify参数定义中的   'fileObjName':'image'
	if err != nil {
		return gin.H{
			"status": config.VueError,
			"name":   header.Filename,
			"msg":    "上传出现错误",
			"size":   size.SizeFormat(header.Size),
			"data":   "/",
			"url":    "",
		}
	}
	//文件的名称
	filename := strings.Split(header.Filename, ".")
	filename_suffix := filename[len(filename)-1]
	uuid, _ := uuid.NewV4()
	new_filename := uuid.String() + "." + filename_suffix
	//判断文件后缀是否允许上传
	if strings.Contains(suffix, filename_suffix) == false {
		return gin.H{
			//"status": strconv.Itoa(config.HttpError),
			"status": config.VueError,
			"name":   header.Filename,
			"msg":    "上传格式不允许，只允许上传上传：" + suffix,
			"size":   size.SizeFormat(header.Size),
			"data":   "/",
			"url":    "",
		}
	}
	//创建文件夹
	path := "uploads/" + filetype + "/" + time.Now().Format("2006/0102/")
	directory.DirectoryMkdir(path)
	//创建文件
	out, err := os.Create(path + new_filename)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		loger.Error(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		loger.Error(file+":"+strconv.Itoa(line), "上传错误：%s", err)
	}
	imgHost := "http://" + c.Request.Host
	//返回值
	return gin.H{
		"status": config.VueSuccess,
		"name":   header.Filename,
		"msg":    "上传成功",
		"size":   size.SizeFormat(header.Size),
		"data":   "/" + path + new_filename,
		"url":    imgHost + "/" + path + new_filename,
		"path":   "/" + path + new_filename,
	}
}
