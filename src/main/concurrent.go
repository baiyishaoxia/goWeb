package main

import (
	"fmt"
	"time"
)

//定义全局管道变量
var Ch = make(chan int, 10000000)

func testConcurrent() {
	for {
		select {
		//读取管道数据
		case a := <-Ch:
			if fun(a) == false {
				break
			} else {
				fmt.Println(a, ":ok")
			}
		}
	}
}

//处理逻辑代码
func fun(a int) bool {
	if a == 999 {
		return true
	} else {
		return false
	}
}
func main() {
	//监听管道
	go testConcurrent()
	for i := 0; i < 1000000; i++ {
		//写入管道
		Ch <- i
	}
	//防止main主线程先终止
	time.Sleep(1 * time.Second)
}
