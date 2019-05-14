package main

import (
	"fmt"
	"runtime"
	"time"
)

var c chan string

func PingPong() {
	i := 0
	for {
		fmt.Println(<-c) //等待直到有值写入
		c <- fmt.Sprintf("From PingPong: hi,#%d", i)
		i++
	}
}

func main() {
	c = make(chan string)
	go PingPong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main: hi,#%d", i) //格式化的字符串
		fmt.Println(<-c)                         //等待接收
	}

	//时区
	t := time.Now()
	fmt.Println(t.Format("Mon Jan 2006-01-02 15:04:05"))

	//闭包与goroutine
	runtime.GOMAXPROCS(runtime.NumCPU()) //提升性能
	mmm := make(chan string)
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
			fmt.Println(<-mmm)
		}(v)
	}
	for n := 1; n <= 3; n++ {
		select {
		case mmm <- "写入呀":
		case mmm <- "写入啊":
		case mmm <- "写入嗄":
		}
	}

}
