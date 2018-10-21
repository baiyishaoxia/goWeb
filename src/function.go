package main

import "fmt"

func main() {
	fmt.Println(fun1(111, "白衣少侠"))
	fmt.Println(fun2())
	slice := []int{10, 9, 8, 7, 6}
	fun3(10, slice, "我", "是", "GO")
	fmt.Println("fun3运行之后:", slice)

	//函数类型的使用 (一切皆类型)
	b := fun2
	fmt.Println(b())

	//匿名函数
	c := func() {
		fmt.Println("我是匿名函数")
	}
	c()

	//闭包
	f := closure(10)
	fmt.Println(f(100))
	fmt.Println(f(200))

	//析构函数
	for i := 1; i < 4; i++ {
		defer fmt.Println("正常的析构", i)

		defer func() {
			fmt.Println("闭包的析构", i)
		}()
	}

	//异常处理
	A()
	B()
	C()

	//demo
	demo()

}

//[变量 类型] [返回值]
func fun1(a int, b string) (int, string) {
	return a, b
}

//[无参数]  [返回值]
func fun2() (a, b, c int) {
	a = 1
	b = 2
	c = a + b
	return a, b, c
}

//[变量 slice 不定长变参]  slice带来的直接就是地址!
func fun3(age int, slice []int, a ...string) {
	fmt.Print(age)
	fmt.Println(a)
	fmt.Println("原来的slice值:", slice)
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	fmt.Println("改变原来的slice值:", slice)
}

//闭包
func closure(x int) func(int) int {
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//[panic/recover]异常处理机制
func A() {
	fmt.Println("我")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("你继续吧!")
		}
	}()
	panic("要")
	fmt.Println("要怎样?能执行到这里吗?")
}
func C() {
	fmt.Println("继续执行")
}

//demo
func demo() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)                      //逆向执行,先进后出 3 2 1 0
		defer func() { fmt.Println("defer_closure i = ", i) }() //引用取i的值 4 4 4 4
		fs[i] = func() {
			fmt.Println("closure i = ", i) //写入闭包函数 4 4 4 4
		}
	}
	for _, f := range fs {
		f()
	}
}
