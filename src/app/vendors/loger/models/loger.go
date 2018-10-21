package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func Error(longfile string, format string, v ...interface{}) {
	file := "./runtime/logs/error/" + time.Now().Format("2006_01_02_15_00_00") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if nil != err {
		panic(err)
	}
	defer logFile.Close()
	loger := log.New(logFile, "[ERROR]  ", log.Ldate|log.Ltime|log.Llongfile)
	loger.Printf(format+"["+longfile+"]", v)
	if gin.Mode() == "debug" {
		fmt.Printf(format+"\r\n", v)
	}
}
func Info(longfile string, format string, v ...interface{}) {
	file := "./runtime/logs/info/" + time.Now().Format("2006_01_02_15_00_00") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if nil != err {
		panic(err)
	}
	defer logFile.Close()
	loger := log.New(logFile, "[INFO]  ", log.Ldate|log.Ltime)
	loger.Printf(format+"  ["+longfile+"]", v)
	if gin.Mode() == "debug" {
		fmt.Printf(format+"\r\n", v)
	}
}
