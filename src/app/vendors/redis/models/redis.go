package models

import (
	"databases"
	"github.com/garyburd/redigo/redis"
)

/**
生成Ke
*/
func Set(key string, value interface{}, expiration int) bool {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	_, err := newredis.Do("SET", key, value) //成功返回的是OK
	if err != nil {
		return false
	}
	if expiration > 0 {
		newredis.Do("EXPIRE", key, expiration)
	}
	return true
}

/**
读取Key
*/
func Get(key string) (reply interface{}, err error) {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	value, err := newredis.Do("GET", key)
	return value, err
}

/**
判断Key是否存在
*/
func Exists(key string) (bool, error) {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	return redis.Bool(newredis.Do("EXISTS", key))
}

/**
删除该db下的所有Key
*/
func DelAll() {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	newredis.Do("flushdb")
}

/**
删除某一个Key
*/
func DelKey(key string) {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	newredis.Do("DEL", key)
}

/**
删除前缀为key_prefix的所有Key
*/
func DelKeyByPrefix(prefix string) {
	newredis := databases.RedisClient().Get()
	defer newredis.Close()
	val, _ := redis.Strings(newredis.Do("KEYS", prefix+"*"))
	newredis.Send("MULTI")
	for i, _ := range val {
		newredis.Send("DEL", val[i])
	}
	newredis.Do("EXEC")
}

/**
value, _ := json.Marshal(admin)
newredis.Set("adminid",value,120)
valueBytes, _ := redis.Bytes(newredis.Get("adminid"))
admin1 := &models.Admin{}
json.Unmarshal(valueBytes,admin1)
fmt.Println(admin1)
fmt.Println(admin1.Id)
fmt.Println(admin1.Name)
fmt.Println(admin1.Password)
fmt.Println(admin1.AdminRoleId)
fmt.Println(admin1.CreatedAt)
*/
