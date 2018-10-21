package models

import (
	"app"
	newredis "app/vendors/redis/models"
	"databases"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type Config struct {
	Name      string   `xorm:"not null pk VARCHAR(100)"`
	Value     string   `xorm:"TEXT"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
}

//region Remark:读取配置 Author:tang
func ReadConfig(name string) string {
	//判断redis是否有缓存数据
	var redis_key string = "admin:config:" + name
	if res, _ := newredis.Exists(redis_key); res == true {
		value, _ := redis.String(newredis.Get(redis_key))
		return value
	}
	//查询数据
	config := new(Config)
	databases.Orm.Where("name=?", name).Get(config)
	//缓存到redis
	newredis.Set(redis_key, config.Value, 60*60)
	return config.Value
}

//endregion
//region Remark:保存配置 Author:tang
func SetConfig(configs []Config) {
	for _, v := range configs {
		config := new(Config)
		result, _ := databases.Orm.Where("name=?", v.Name).Exist(config)
		if result == false {
			config.Name = v.Name
			config.Value = v.Value
			databases.Orm.Insert(config)
		} else {
			config.Value = v.Value
			databases.Orm.Where("name=?", v.Name).Update(config)
		}
	}
	newredis.DelKeyByPrefix("admin:config")
}

//endregion
func readConfig() []Config {
	var redis_key string = "admin:config"
	//判断redis是否有缓存数据
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		var configs []Config
		//反序列化json格式
		json.Unmarshal(valueBytes, &configs)
		return configs
	}
	//查询数据
	configs := make([]Config, 0)
	databases.Orm.Find(&configs)
	//缓存到redis [序列化将数据编码成json字符串]
	value, _ := json.Marshal(configs)
	newredis.Set(redis_key, value, 60*60)
	return configs
}
