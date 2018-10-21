package main

import (
	"fmt"
)

func main() {
	a := 1
	a++
	var p *int = &a
	fmt.Println(*p)
	//选择语句if
	if b := 3; b > a {
		fmt.Println(b)
	}
	//循环语句for
	a = 1
	for {
		a++
		if a > 3 {
			break
		}
		println("for1:", a)
	}
	println("第一种for方式结束!")
	for a <= 4 {
		a++
		fmt.Println("for2:", a)
		println("第二种for方式结束!")
	}
	k := "123456789"
	len := len(k)
	for i := 1; i <= len; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	println("第三种for方式结束!")
	//选择语句switch
	week := 6
	switch week {
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("没放假哦!")
	}
	switch week = 1; {
	case week < 6:
		fmt.Println("1-5正常上班")
	case week > 6:
		fmt.Println("周末了")
	}
	//跳转语句 goto break continue
Tang1:
	for {
		for b := 0; b <= 1000000000; b++ {
			if b > 100000000 {
				fmt.Println("通过goto跳转结束!,当前b值为:", b)
				goto Tang2
			}
			if b > 500000000 {
				fmt.Println("通过跳转结束!,当前b值为:", b)
				break Tang1
			}
		}
	}
Tang2:
}
