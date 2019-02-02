package queue

import "other/linkedList"

/*
队列的特性较为单一，基本操作即初始化、获取大小、添加元素、移除元素等。
最重要的特性就是满足先进先出
*/
type Queue struct {
	linkedList.List
}

//加入队列
func (this *Queue) Put(data linkedList.Object) {
	this.Add(data)
}

//pop出队列
func (this *Queue) Pop() linkedList.Object {
	if this.GetHeadNode() == nil {
		panic("this queue is nil")
	}
	headNode := this.GetHeadNode()
	this.RemoveAtIndex(0)
	return headNode
}

//获得队列的长度
func (this *Queue) GetSize() linkedList.Object {
	return this.GetSize()
}
