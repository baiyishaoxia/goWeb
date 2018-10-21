package main

import "fmt"

const (
	e = iota
	d
)

func main() {
	//一元运算符(取反再减一)
	fmt.Println(^7)
	//取反
	fmt.Println(!false)
	//左移10位(1000000000) 2^10
	fmt.Println(1 << 10)
	//右移1位 111 (011)  3
	fmt.Println(7 >> 1)
	fmt.Println(6 &^ 11)
	/*
		  6 : 0110
		 11 : 1011
		------------
		  &   0010  = 2
		  |   1111  =15
		  ^   1101  =13
		  &^  0100  =4  (第二行为1的强制改为0)
	*/
	fmt.Println(e, d)
	if d > 0 && (10/d) > 1 {
		println("ok")
	}
}
