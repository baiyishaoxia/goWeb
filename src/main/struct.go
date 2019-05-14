package main

import "fmt"

type user struct {
	Name string
	Age  int
}

//匿名结构
type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

//匿名字段
type man struct {
	string
	int
}

func main() {
	//字面值初始化
	a := user{
		Name: "tang",
		Age:  21,
	}
	fmt.Println("字面值初始化", a)
	//单个赋值
	a.Name = "白衣少侠"
	a.Age = 18
	fmt.Println("单个赋值", a)
	//函数值传递 struct
	fun_user1(a)
	//指针地址传递 struct
	fun_user2(&a)
	//原来的值发生改变
	fmt.Println("执行fun_user2后:", a)

	// 这样一来,b就是指向某个结构的指针
	b := &user{
		Name: "tang",
		Age:  21,
	}
	fun_user3(b)
	fmt.Println("执行fun_user3后:", b)

	//匿名结构
	c := &struct {
		Name string
		Age  int
	}{
		Name: "hony",
		Age:  10,
	}
	fmt.Println(c)

	//匿名结构
	m := person{Name: "wang", Age: 15}
	m.Contact.Phone = "123456789"
	m.Contact.City = "深圳"
	fmt.Println(m)

	//匿名字段 [保证类型顺序严格一致]
	n := man{"字符", 11}
	fmt.Println(n)
}

func fun_user1(a user) {
	a.Name = "浮云"
	a.Age = 100
	fmt.Println("fun_user1 的 struct:", a)
}

func fun_user2(a *user) {
	a.Name = "神马"
	a.Age = 1000
	fmt.Println("fun_user2 的 struct:", a)
}

func fun_user3(a *user) {
	a.Name = "哈哈"
	fmt.Println("fun_user3 的 struct:", a)
}

//region Remark:注意 Author:tang
/*
   type man1 struct {
	string
	int
}
type man2 struct {
	string
	int
}
n1 := man1{"字符", 11}
n2 := man2{"字符", 11}
//这里n1与n2没有可比性
n3 := man1{"字符", 11}
n4 := man1{"字符", 11}
fmt.Println(n3 == n4)
//这就可以比较
*/
//endregion
