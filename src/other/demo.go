package main

import (
	"fmt"
	"os"
)

type Noder struct {
	data int
	Next *Noder
}

func initList() *Noder {
	pHead := new(Noder)
	pHead.Next = pHead
	return pHead //返回头指针
}

//创建尾指针的单循环链表
func createList(list **Noder) {
	if !isempty(*list) {
		cleanList(*list)
	}
	var val int
	p, q := *list, *list
	fmt.Println("请输入结点数据,输入0结束输入")
	fmt.Scanf("%d", &val)
	for val != 0 {
		pnew := new(Noder)
		pnew.data = val
		pnew.Next = p
		q.Next = pnew
		q = pnew
		fmt.Scanf("%d", &val)
	}
	*list = q
}

//清空循环链表
func cleanList(list *Noder) {
	if isempty(list) {
		return
	}
	phead := list.Next  //头结点
	p := list.Next.Next //第一个结点
	q := p
	for p != list.Next {
		q = p.Next
		p = nil
		p = q
	}
	phead.Next = phead
}

//插入结点
func insertList(list **Noder) {
	var index, val int
	fmt.Printf("请输入要插入的位置：（值范围：1-%d）\n", listLength(*list)+1)
	fmt.Scanf("%d", &index)
	if index < 1 || index > listLength(*list)+1 {
		fmt.Println("位置值越界")
		return
	}
	fmt.Println("请输入要插入的值：")
	fmt.Scanf("%d", &val)
	j := 1
	p, q := (*list).Next, (*list).Next //头结点
	for j < index {
		p = p.Next
		j++
	}
	pnew := new(Noder)
	pnew.data = val
	pnew.Next = p.Next
	p.Next = pnew
	if pnew.Next == q {
		*list = pnew
	}
}
func deleList(list **Noder) {
	var index int
	fmt.Printf("请输入要删除的位置：（值范围：1-%d）\n", listLength(*list))
	fmt.Scanf("%d", &index)
	if index < 1 || index > listLength(*list) {
		fmt.Println("位置值越界")
		return
	}
	j := 1
	p, q := (*list).Next, (*list).Next //头结点
	//查找index-1结点
	for j < index {
		p = p.Next
		j++
	}
	cur := p.Next
	p.Next = cur.Next
	if p.Next == q {
		*list = p
	}
	cur = nil
}
func locateList(list *Noder) {
	fmt.Println("请输入要查找的值：")
	var val int
	fmt.Scanf("%d", &val)
	q := list.Next.Next //第一个结点
	var loc int = 0
	for q != list.Next {
		loc++
		if q.data == val {
			break
		}
		q = q.Next
	}
	if loc == 0 {
		fmt.Println("链表中未找到你要的值")
	} else {
		fmt.Printf("你查找的值的位置为：%d\n", loc)
	}
}
func traverse(list *Noder) {
	if isempty(list) {
		fmt.Println("空链表")
		return
	}
	fmt.Println("链表内容如下：")
	p := list.Next.Next //第一个结点
	for p != list.Next {
		fmt.Printf("%5d", p.data)
		p = p.Next
	}
	fmt.Println()
}
func isempty(list *Noder) bool {
	if list.Next == list {
		return true
	} else {
		return false
	}
}
func listLength(list *Noder) int {
	if isempty(list) {
		return 0
	}
	var len int = 0
	p := list.Next.Next //第一个结点
	for p != list.Next {
		len++
		p = p.Next
	}
	return len
}
func main() {
	list := initList()
	var flag int
	fmt.Println("1.初始化链表")
	fmt.Println("2.插入结点")
	fmt.Println("3.删除结点")
	fmt.Println("4.返回结点位置")
	fmt.Println("5.遍历链表")
	fmt.Println("6.清空链表")
	fmt.Println("0.退出")
	fmt.Println("请选择你的操作:")
	fmt.Scanf("%d", &flag)
	for flag != 0 {
		switch flag {
		case 1:
			createList(&list)
		case 2:
			insertList(&list)
		case 3:
			deleList(&list)
		case 4:
			locateList(list)
		case 5:
			traverse(list)
		case 6:
			cleanList(list)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("无效操作")
		}
		fmt.Println("请选择你的操作:")
		fmt.Scanf("%d", &flag)
	}
}
