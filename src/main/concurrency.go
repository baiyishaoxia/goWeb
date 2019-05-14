package main

//并发
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//无缓存是同步的,有缓存是异步的  (channel是goroutine沟通的桥梁)
	c := make(chan bool)
	go func() {
		fmt.Println("go to school!")
		<-c
		close(c) //关闭
	}()
	c <- true

	//for v := range c {
	//	fmt.Println(v)
	//}
	//(goroutine原则是通过通信来共享内存,而不是共享内存来通信)
	//go Go()
	//time.Sleep(2 * time.Second) //暂停2秒钟

	//runtime.GOMAXPROCS(runtime.NumCPU()) //提升性能
	//ch := make(chan bool, 10)            //使用缓存10次 (异步)
	//for i := 0; i < 10; i++ {
	//	go Sum(ch, i)
	//}
	//for i := 0; i < 10; i++ {
	//	<-ch //使用通过循环取10次
	//}

	runtime.GOMAXPROCS(runtime.NumCPU()) //提升性能
	wg := sync.WaitGroup{}               //使用WaitGroup 同步
	wg.Add(10)                           //10个任务
	for i := 0; i < 10; i++ {
		go Sum_tb(&wg, i)
	}
	wg.Wait()

	//select的使用
	m := make(chan int)
	go func() {
		for v := range m {
			fmt.Println("输出:", v)
		}
	}()
	for n := 1; n < 10; n++ {
		select {
		case m <- 0:
		case m <- 1:
		}
	}

	//超时
	select {
	case v := <-m:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("已超时!")

	}

}

func Go() {
	fmt.Println("GO GO GO!")
}

func Sum(ch chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//if index == 9 { //在分配任务不一定是按部就班,不能直接写 index == 9 就结束
	//	ch <- true
	//}
	ch <- true
}

//同步
func Sum_tb(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done() //每完成一次任务就done
}
