package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"time"
)

//Db数据库连接池
var DB *sql.DB

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "student"
)

type Time time.Time

func (c Time) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

type Users struct {
	Id            int64
	Name          string
	Email         string
	RememberToken string
	Password      int64
	CreatedAt     Time `xorm:"created"`
	UpdatedAt     Time `xorm:"updated"`
}

type User2 struct {
	id             int64
	name           string
	email          string
	remember_token string
	password       int64
	created_at     Time
	updated_at     Time
}

var engine *xorm.Engine

func main() {
	//xorm 连接
	engine, _ = xorm.NewEngine("mysql", "root:root@/student?charset=utf8")
	engine.ShowSQL(true) //在控制台打印出生成的SQL语句
	user := new([]Users)

	//user.Name = "测试"
	//has, err1 := engine.Insert(user)
	//if has < 1 {
	//	fmt.Println(err1.Error())
	//}
	//fmt.Println("---------------新增成功---------------!")

	err2 := engine.Select("id,name,email").Limit(3, 0).Find(user)
	check(err2)
	for _, val := range *user {
		fmt.Println(val.Id, val.Name, val.Email, val.CreatedAt, val.UpdatedAt)
	}

	users := new(Users)
	rows, err3 := engine.Where("id > ?", 5).Rows(users)
	check(err3)
	defer rows.Close()
	for rows.Next() {
		err3 = rows.Scan(users)
		fmt.Println(*users)
	}
	fmt.Println("---------------查询的2种方式成功---------------!")

	//开启事务
	db := engine.NewSession()
	defer db.Close()
	err := db.Begin()
	user1 := Users{Name: "小李", Email: "9521@qq.com", Password: 123456}
	_, err = db.Insert(&user1)
	if err != nil {
		db.Rollback()
		return
	}
	user2 := Users{Name: "xxx"}
	_, err = db.Where("id = ?", 2).Update(&user2)
	if err != nil {
		db.Rollback()
		return
	}
	_, err = db.Exec("delete from users where name = ?", user2.Name)
	if err != nil {
		db.Rollback()
		return
	}
	err = db.Commit()
	if err != nil {
		return
	}

	//缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1)
	engine.MapCacher(&users, cacher)
	engine.Exec("update users set name = ? where id = ?", "YYY", 9)
	engine.ClearCache(new(Users))

	//mysql驱动连接
	//var u User2
	//database := InitDB()
	//rows, err := database.Query("select * from users")
	//if err != nil {
	//	fmt.Println("查询出错了")
	//}
	//userinfo := make(map[interface{}]interface{})
	//for rows.Next() {
	//	err := rows.Scan(&u.id, &u.name, &u.email, &u.password, &u.remember_token, &u.created_at, &u.updated_at)
	//	check(err)
	//	userinfo[u.id] = u
	//}
	//for _, val := range userinfo {
	//	fmt.Println(val)
	//}
}

//注意方法名大写，就是public
func InitDB() *sql.DB {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return nil
	}
	fmt.Println("connnect success")
	return DB
}

//region Remark:输出错误信息 Author:tang
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//endregion
