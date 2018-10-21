package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello ", name, ",my name is ", u.Name)
}

func Info(o interface{}) {
	typ := reflect.TypeOf(o)
	fmt.Println("Type:", typ.Name())
	val := reflect.ValueOf(o)
	fmt.Println("Fields:")

	if k := typ.Kind(); k != reflect.Struct {
		fmt.Println("XXXXXXXXXXXX")
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		attribute := typ.Field(i)
		value := val.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", attribute.Name, attribute.Type, value)
	}

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("%6s:%v\n", method.Name, method.Type)
	}
}

func main() {
	user := User{1, "白衣少侠", 18}
	Info(user) //以值拷贝的方式传递到info当中

	m := Manage{User: User{2, "Ok", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println("修改值:", x)

	u := User{3, "NO", 16}
	fmt.Println("u的值为:", u)
	Set(&u)
	fmt.Println("通过反射对值的修改:", u)

	v = reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("萧炎")}
	fmt.Print("通过反射对方法的调用: ")
	mv.Call(args)
}

type Manage struct {
	User
	title string
}

func Set(o interface{}) {
	v := reflect.ValueOf(o) //取值
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("kind()判断v的类型 && 不能够被修改")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("不存在该字段")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("萧薰儿")
	}

}
