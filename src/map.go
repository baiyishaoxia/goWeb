package main

import (
	"fmt"
	"sort"
)

func main() {
	var m map[int]string = make(map[int]string)
	m[1] = "OK"
	fmt.Println("第一个map:", m)
	fmt.Println(m[1]) //取值
	delete(m, 1)      //删除

	var n map[int]map[int]string = make(map[int]map[int]string)
	//仅仅对key为1的第二层map初始化
	n[1] = make(map[int]string)
	n[1][1] = "haha"
	fmt.Println("第二个map:", n)

	//key为2的并没有初始化
	a, ok := n[2][1]
	fmt.Println("key为2没有初始化:", a, ok)
	if !ok {
		n[2] = make(map[int]string)
	}
	n[2][1] = "xixi"
	a, ok = n[2][1]
	fmt.Println("key为2已初始化:", a, ok)

	for _, val := range n {
		fmt.Println(val[1])
	}

	sm := make([]map[int]string, 5)
	for key, v := range sm {
		v = make(map[int]string, 1)
		v[key] = "我们都有一个家"
		fmt.Println("都已经初始化了:", v[key])
	}
	fmt.Println("为啥又为空了?(对sm本身没影响)", sm)

	var j int = 0
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][j] = "我们都有2个家"
		fmt.Println("都已经初始化了:", sm[i][j])
		j++
	}
	fmt.Println("这回都好了吧!", sm)

	mm := map[int]string{3: "c", 2: "b", 1: "a", 4: "d", 5: "e"}
	fmt.Println("mm的map值:", mm)
	ss := make([]int, len(mm))
	i := 0
	for k, _ := range mm {
		ss[i] = k
		i++
	}
	sort.Ints(ss) //排序 (这样一来,就可以有序的取出 mm中的每个值)
	for _, v := range ss {
		println(mm[v])
	}
	fmt.Println(mm)

	sss := make(map[string]int)
	for k, v := range mm {
		sss[v] = k
	}
	fmt.Println(sss)

}
