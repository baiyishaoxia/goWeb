package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	t := time.Now() // get current time
	more()          //业务逻辑部分
	elapsed := time.Since(t)
	fmt.Println("消耗时长: ", elapsed)
}

func more() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //使用CPU最大核心数
	//runtime.GOMAXPROCS(1)
	strCh := make(chan string, 1000000) //百万缓冲区

	//开2个线程写 2亿条数据
	go AddStr(strCh, "polaris")
	go AddStr(strCh, "studygolang")

	//开100个线程读
	for i := 0; i < 100; i++ {
		go PrintStr(strCh, i)
	}

	time.Sleep(1 * time.Minute) //给线程足够时间运行,这里先给1 Minute
}

//写入chan管道
func AddStr(ch chan<- string, str string) {
	for i := 0; i < 100000000; i++ {
		ch <- str + strconv.Itoa(i)
	}
}

//读出chan管道
func PrintStr(ch <-chan string, i int) {
	for {
		select {
		case str := <-ch:
			fmt.Println("i=", i, str)
			break
		}
	}
}
