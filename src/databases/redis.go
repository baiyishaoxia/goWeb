package databases

import (
	loger "app/vendors/loger/models"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-ini/ini"
	"runtime"
	"strconv"
	"time"
)

var (
	host      string
	port      string
	password  string
	dbname    int
	maxidle   int
	maxactive int
)

func RedisClient() *redis.Pool {
	//读取配置文件
	cfg, err := ini.Load("config/redis.conf")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		loger.Error(file+":"+strconv.Itoa(line), "无法打开 'config/redis.conf': %v", err)
	}
	key, _ := cfg.Section("redis").GetKey("host")
	host = key.Value()
	key, _ = cfg.Section("redis").GetKey("port")
	port = key.Value()
	key, _ = cfg.Section("redis").GetKey("password")
	password = key.Value()
	key, _ = cfg.Section("redis").GetKey("dbname")
	dbname, _ = strconv.Atoi(key.Value())
	key, _ = cfg.Section("redis").GetKey("maxidle")
	maxidle, _ = strconv.Atoi(key.Value())
	key, _ = cfg.Section("redis").GetKey("maxactive")
	maxactive, _ = strconv.Atoi(key.Value())
	//连接redis
	RedisClient := &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     maxidle,
		MaxActive:   maxactive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				fmt.Print(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				fmt.Print(err)
				return nil, err
			}
			// 选择db
			c.Do("SELECT", dbname)
			return c, nil
		},
	}
	return RedisClient
}
