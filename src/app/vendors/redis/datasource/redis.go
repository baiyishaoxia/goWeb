package datasource

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	RedisPool     *redis.Client
	SessionsStore sessions.Store
)

type redisDbConf struct {
	Redis struct {
		Host      string `yaml:host`
		Port      string `yaml:port`
		Password  string `yaml:password`
		Database  int    `yaml:database`
		Maxidle   int    `yaml:maxidle`
		Maxactive int    `yaml:maxactive`
	}
	Session struct {
		Host      string `yaml:host`
		Port      string `yaml:port`
		Password  string `yaml:password`
		Database  int    `yaml:database`
		Maxidle   int    `yaml:maxidle`
		Maxactive int    `yaml:maxactive`
	}
}

func init2() {
	var redisDbConfData = redisDbConf{}
	yamlFile, err := ioutil.ReadFile("./config/redis.yml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &redisDbConfData)
	if err != nil {
		fmt.Println(err)
	}
	RedisPool = redis.NewClient(&redis.Options{
		Addr:     redisDbConfData.Redis.Host + ":" + redisDbConfData.Redis.Port,
		Password: redisDbConfData.Redis.Password, // no password set
		DB:       redisDbConfData.Redis.Database, // use default DB
	})
	SessionsStore, _ = sessions.NewRedisStore(redisDbConfData.Session.Database, "tcp", redisDbConfData.Session.Host+":"+redisDbConfData.Session.Port, redisDbConfData.Session.Password, []byte("secret"))
	option := sessions.Options{
		MaxAge: 60 * 60 * 2, //2h
		Path:   "/",
	}
	SessionsStore.Options(option)
	return
}
