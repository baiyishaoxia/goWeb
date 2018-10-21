package controllers

import (
	directory "app/vendors/directory/models"
	loger "app/vendors/loger/models"
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

func PostUploadWangEditorImage(c *gin.Context) {
	c.JSON(http.StatusOK, wangEditorUpload(c, "Filedata", "editorv1image", "bmp,gif,jpg,jpeg,jpe,png"))
}
func wangEditorUpload(c *gin.Context, uploadFilename string, filetype string, suffix string) gin.H {
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		//返回值
		return gin.H{
			"errno": 1,
		}
	}
	files := c.Request.MultipartForm.File[uploadFilename]
	dataPath := make([]string, len(files))
	for i, _ := range files {
		//得到上传的文件
		file, err := files[i].Open()
		if err != nil {
			return gin.H{
				"errno": 1,
			}
			break
		}
		//文件的名称
		filename := strings.Split(files[i].Filename, ".")
		filename_suffix := filename[len(filename)-1]
		uuid, _ := uuid.NewV4()
		new_filename := uuid.String() + "." + filename_suffix
		//判断文件后缀是否允许上传
		if strings.Contains(suffix, filename_suffix) == false {
			return gin.H{
				"errno": 1,
			}
			break
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
		dataPath[i] = "/" + path + new_filename
	}

	//返回值
	return gin.H{
		"data":  dataPath,
		"errno": 0,
	}
}
