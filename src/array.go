package main

import "fmt"

func main() {
	//最后一个数据为1
	a := [20]int{19: 1}
	//隐式长度为 3
	b := [...]int{0: 1, 1: 2, 2: 3}
	c := [...]int{2: 3}
	fmt.Println("数组a:", a)
	//指针数组
	var p *[3]int = &b
	fmt.Println("数组b:", b)
	fmt.Println("指向数组b:", p)
	//比较2数组
	fmt.Println("数组b与数组c比较:", c == b)
	//普通数组赋值
	i := [10]int{}
	i[1] = 2
	fmt.Println("普通数组赋值:", i)
	//返回一个指向数组的指针并给索引赋值
	q := new([10]int)
	q[1] = 2
	fmt.Println("返回一个指向数组的指针:", q)
	//多维数组
	m := [2][4]int{
		{1, 1, 1, 3: 9},
		{2, 2, 2, 3: 9}}
	fmt.Println("多维数组:", m)
	//冒泡排序
	num := len(b)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if b[i] < b[j] {
				t := b[i]
				b[i] = b[j]
				b[j] = t
			}
		}
	}
	fmt.Println("对b数组冒泡排序之后:", b)
}
