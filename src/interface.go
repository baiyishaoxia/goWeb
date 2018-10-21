package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string {
	return pc.name
}
func (pc PhoneConnect) Connect() {
	fmt.Println("connect:", pc.name)
}

func main() {
	var a USB
	//只要某一个类型有该接口的所有方法签名,即算实现该接口! 称之为strictural Typing
	a = PhoneConnect{"Iphone手机已连接!"}
	a.Connect()
	Disconnect(a)

	pc := PhoneConnect{"字面值定义"}
	pl := &PhoneConnect{"字面值定义"}
	var b Connecter
	b = Connecter(pc)
	b.Connect()

	pc.name = "哈哈哈哈哈哈哈哈哈哈哈" //这里只是一个复制品,不会对原来造成影响
	b.Connect()
	pl.name = "哈哈哈哈哈哈哈哈哈哈哈" //这里只是一个地址,会对原来造成影响
	var pp Connecter
	pp = Connecter(pl)
	pp.Connect()

	var m interface{} //空接口可以实现任何类型的容器
	fmt.Println(m == nil)
	var p *int = nil
	m = p
	fmt.Println(m == nil)
}

func Disconnect(usb interface{}) {
	if pc, ok := usb.(PhoneConnect); ok {
		fmt.Println("Disconnect", pc.name)
	} else {
		fmt.Println("未知设备!")
	}

	switch v := usb.(type) {
	case PhoneConnect:
		fmt.Println("1111111111", v.name)
	default:
		fmt.Println("2222222222!")
	}

}
