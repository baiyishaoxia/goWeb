package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type C int

type Inc int

func main() {
	a := &A{Name: "嗯嗯"}
	a.fun()
	b := &B{Name: "额额"}
	b.fun()
	//method value方式
	var c C
	c.fun()
	//method expression方式
	(*C).fun(&c)

	//调用
	var inc Inc
	inc.fun(100)
	fmt.Println("为inc加值:", inc)
}
func (a *A) fun() {
	a.Name = "哈哈"
	fmt.Println(a.Name)
}

func (b *B) fun() {
	fmt.Println(b.Name)
}

func (c *C) fun() {
	fmt.Println("这样厉害了")
}

func (inc *Inc) fun(num int) {
	*inc += Inc(num)
}
