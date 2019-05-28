package main

import "fmt"

func main() {
	aa := "sdsa"
	mobile := string([]byte(aa)[:3]) + "****" + string([]byte(aa)[7:11])
	fmt.Println("----", mobile)

	c := make(chan int, 2)
	c <- 10
	c <- 100
	close(c)

	v, ok := <-c // v=10,ok=true，虽然c关闭了，但是有数据，ok依然是true
	v, ok = <-c  // v=100,ok=true，读失败了。
	v, ok = <-c  // v=0,ok=false，读失败了。
	fmt.Println(v, ok)

}
