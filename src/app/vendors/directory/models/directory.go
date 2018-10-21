package models

import (
	loger "app/vendors/loger/models"
	"os"
	"runtime"
	"strconv"
)

/**
创建文件夹
*/
func DirectoryMkdir(path string) {
	if res, _ := DirectoryExists(path); res == false {
		err := os.MkdirAll(path, os.ModePerm)
		_, file, line, _ := runtime.Caller(0) //获取错误文件和错误行
		loger.Error(file+":"+strconv.Itoa(line), "创建文件夹错误：%s", err)
	}
}

/**
判断文件夹是否存在
*/
func DirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
