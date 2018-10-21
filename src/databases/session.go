package databases

import (
	loger "app/vendors/loger/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/go-ini/ini"
	"runtime"
	"strconv"
)

func SessionStore() sessions.Store {
	//读取配置文件
	cfg, err := ini.Load("config/session.conf")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		loger.Error(file+":"+strconv.Itoa(line), "无法打开 'config/session.conf': %v", err)
	}
	key, _ := cfg.Section("session").GetKey("host")
	host = key.Value()
	key, _ = cfg.Section("session").GetKey("port")
	port = key.Value()
	key, _ = cfg.Section("session").GetKey("password")
	password = key.Value()
	key, _ = cfg.Section("session").GetKey("dbname")
	dbname, _ = strconv.Atoi(key.Value())
	key, _ = cfg.Section("session").GetKey("maxidle")
	maxidle, _ = strconv.Atoi(key.Value())
	key, _ = cfg.Section("session").GetKey("maxactive")
	maxactive, _ = strconv.Atoi(key.Value())
	store, _ := sessions.NewRedisStore(dbname, "tcp", host+":"+port, password, []byte("secret"))
	option := sessions.Options{
		MaxAge: 60 * 60 * 2, //2h
		Path:   "/",
	}
	store.Options(option)
	return store
}
