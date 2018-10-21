package background

import (
	"app/vendors/zip/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//region Remark:备份管理 Author:tang
func GetDatabaseList(c *gin.Context) {
	files, _ := ioutil.ReadDir("./uploads/backup/")
	//模版
	c.HTML(http.StatusOK, "databases/list", gin.H{
		"Title": "Background Login",
		"Data":  files,
		"Count": len(files),
	})
}

//endregion

//region Remark:创建备份 Author:tang
func GetDatabaseBackup(c *gin.Context) {
	uuid, _ := uuid.NewV4()
	file := "./uploads/backup/" + uuid.String() + "_" + time.Now().Format("20060102153748") + ".sql"
	//转储数据库的所有表结构和数据到一个文件
	databases.Orm.DumpAllToFile(file)
	f1, err := os.Open(file)
	fmt.Println(err)
	var files = []*os.File{f1}
	filezip := "./uploads/backup/" + uuid.String() + "_" + time.Now().Format("20060102153748") + ".zip"
	models.Zip(files, filezip)
	os.Remove(file)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "压缩成功,请尽快下载到本地并删除服务器的备份文件",
		"url":    "/admin/databases/list",
	})
}

//endregion
//region Remark:下载 && 删除 Author:tang
func GetDatabaseDown(c *gin.Context) {
	zipName := "./uploads/backup/" + c.Param("name")
	c.File(zipName)
}
func GetDatabaseDel(c *gin.Context) {
	file := "./uploads/backup/" + c.Param("name")
	os.Remove(file)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/databases/list",
	})
}

//批量删除
func GetDatabaseDelAll(c *gin.Context) {
	fileArray := c.PostFormArray("name[]")
	for i := 0; i < len(fileArray); i++ {
		file := "./uploads/backup/" + fileArray[i]
		os.Remove(file)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/databases/list",
	})
}

//endregion
