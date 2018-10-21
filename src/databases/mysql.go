package databases

import (
	loger "app/vendors/loger/models"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"runtime"
	"strconv"
)

var (
	Orm *xorm.Engine
)

func init() {
	cfg, err := ini.Load("config/db.conf")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		loger.Error(file+":"+strconv.Itoa(line), "无法打开 'config/db.conf': %v", err)
	}
	key, _ := cfg.Section("db").GetKey("host")
	host := key.Value()
	key, _ = cfg.Section("db").GetKey("port")
	port := key.Value()
	key, _ = cfg.Section("db").GetKey("user")
	user := key.Value()
	key, _ = cfg.Section("db").GetKey("password")
	password := key.Value()
	key, _ = cfg.Section("db").GetKey("dbname")
	dbname := key.Value()
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, dbname)
	Orm, err = xorm.NewEngine("mysql", mysqlInfo)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		loger.Error(file+":"+strconv.Itoa(line), "MySql 数据库连接失败：%s", err)
	}
	//连接测试
	if err := Orm.Ping(); err != nil {
		_, file, line, _ := runtime.Caller(1)
		loger.Error(file+":"+strconv.Itoa(line), "MySql 数据库连接测试：%s", err)
	}
	//日志打印SQL
	//Orm.ShowSQL(true)
	//设置连接池的空闲数大小
	Orm.SetMaxIdleConns(1000000)
	//设置最大打开连接数
	Orm.SetMaxOpenConns(5000000)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	Orm.SetTableMapper(core.SnakeMapper{})
	//Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	//连接成功
	_, file, line, _ := runtime.Caller(0)
	loger.Info(file+":"+strconv.Itoa(line), "Mysql %v", "连接成功")
	//defer Orm.Close()
}
