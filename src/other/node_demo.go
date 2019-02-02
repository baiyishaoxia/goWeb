package main

import (
	"fmt"
	"other/linkedList"
)

//测试节点
func main() {
	list := linkedList.List{}

	//1.往单链表末尾追加元素2, 3, 4, 5
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	linkedList.Traverse(list.GetHeadNode()) //打印链表
	//2.从头部添加元素head_node
	list.Add("head_node")

	fmt.Println("长度======", list.Length())

	//3.判断是否为空链表
	bool := list.IsEmpty()
	fmt.Println(bool)

	//4.在指定位置2插入元素 2indexValue
	list.Insert(2, "2_index_value")
	travselLinkList(&list)

	//5.是否包含元素2_index_value
	isContain := list.Contain("2_index_value")
	fmt.Println("isContain[2_index_value]:", isContain)

	//6.删除元素2_index_value
	list.Remove("2_index_value")
	travselLinkList(&list)

	//7.从位置2删除元素
	list.RemoveAtIndex(2)
	travselLinkList(&list)

}

func travselLinkList(list *linkedList.List) {
	//遍历
	head := list.GetHeadNode()
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
	fmt.Println("--------------------")
}
