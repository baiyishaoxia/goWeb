package main

import (
	"fmt"
	"math"
	"strconv"
)

//类型别名
type (
	牛逼 string
)

//全局变量
var k bool = true

func main() {
	var a [2]byte
	var b 牛逼 = "中文类型名"
	fmt.Println(a)
	fmt.Println(math.MaxInt8)
	fmt.Println(b)

	c := "我是系统推断的简写"
	fmt.Println(c)

	fmt.Println(k)
	//多变量赋值
	var t1, t2, t3 int = 1, 2, 3
	t1, t2, t3 = 4, 5, 6
	fmt.Println(t1, t2, t3)

	//忽略变量 (常用于函数有多个返回值下使用)
	r1, _, r3, r4 := 6, 7, 8, 9
	fmt.Println(r1, r3, r4)

	//强制类型转换
	var p float32 = 100.1
	fmt.Println(p)
	u := int(p)
	fmt.Println(u)

	//整型转字符型会变成ASCII码
	var i int = 65
	//转ASCII
	o1 := string(i)
	//int 转 string
	o2 := strconv.Itoa(i)
	//string 转 int
	i, _ = strconv.Atoi(o2)
	fmt.Println(o1)
	fmt.Println(o2)
	fmt.Println(i)
}
