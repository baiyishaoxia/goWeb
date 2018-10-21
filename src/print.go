//当前程序的包名
package main
//引入包并设置别名
import system "fmt"

const PI  = 3.1415926
//全局变量的声明与赋值
var name  =  "golang"
//一般类型的声明(别名)
type agee int

//结构体的声明
type str struct {

}

//接口的声明
type golang interface {

}

func main()  {
	age := 18
	system.Println(age)
	system.Println("Hello Word!你好,世界!")

	var a agee = 10000
	system.Println(a)

}

