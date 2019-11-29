package main

import "fmt"

type aaa struct {
	name string
}
type bbb struct {
	ss  *aaa
	//ss  aaa
}

func (that *bbb)clone() *bbb{
	var n = *that
	return &n
}

//todo test clone ...
func main()  {
	var item = new(aaa)  //new aaa ...
	item.name = "aaa"
	var item2 =new(bbb)  //new bbb ...
	item2.ss = item
	//item2.ss = *item
	item2.ss.name = "bbb"
	var item3 = item2.clone() //clone ...
	item3.ss.name = "ccc"
	var item4 = item2.clone() //clone ...
	item4.ss.name = "ddd"
	fmt.Println("----------ok----------",item,item2.ss,item3.ss,item4.ss)
}
