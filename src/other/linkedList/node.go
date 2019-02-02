package linkedList

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node //头节点
}

//判断是否为空的单链表
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//单链表的长度
func (this *List) Length() int {
	cur := this.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//获取头部节点
func (this *List) GetHeadNode() *Node {
	return this.headNode
}

//从头部添加元素
func (this *List) Add(data Object) {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
}

//从尾部添加元素
func (this *List) Append(data Object) {
	node := &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

//在指定位置添加元素
func (this *List) Insert(index int, data Object) {
	if index < 0 {
		this.Add(data)
	} else if index > this.Length() {
		this.Append(data)
	} else {
		pre := this.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		//当循环退出后，pre指向index -1的位置
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除指定元素
func (this *List) Remove(data Object) {
	pre := this.headNode

	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//删除指定位置的元素
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	if index <= 0 {
		this.headNode = pre.Next
	} else if index >= this.Length() {
		//报错 err
	} else {
		count := 0 //index = 3
		for count != (index-1) && pre.Next != nil {
			count++        //2
			pre = pre.Next //2
		}
		pre.Next = pre.Next.Next
	}
}

//是否包含某个元素
func (this *List) Contain(data Object) bool {
	cur := this.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

//#遍历链表  [list为头节点]
func Traverse(list *Node) {
	if is_empty(list) {
		fmt.Println("空链表")
		return
	}
	fmt.Println("链表内容如下：")
	p := list //头节点
	for p != nil {
		fmt.Printf("%5d", p.Data)
		p = p.Next
	}
	fmt.Println()

}

//#清空循环链表
func CleanList(list *Node) {
	if is_empty(list) {
		return
	}
	list.Data = 0
	list.Next = nil
}

//#是否为空
func is_empty(list *Node) bool {
	if list.Next == list {
		return true
	} else {
		return false
	}
}

//测试代码在node_demo.go中 ,此处省略...
//实现部分在linkedList包里面，具体实现的功能有:
//1.判断是否为空的单链表
//2.单链表的长度
//3.获取头节点
//4.从头部添加元素
//5.从尾部添加元素
//6.在指定位置添加元素
//7.删除指定元素
//8.删除指定位置的元素
//9.判断是否包含某个元素
