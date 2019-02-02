package queue

import "log"

//环形队列实现 队列，先进先出。追加至队尾，弹出队顶
type Queen struct {
	Length   int64 //队列长度
	Capacity int64 //队列容量
	Head     int64 //队头
	Tail     int64 //队尾
	Data     []interface{}
}

//初始化
func MakeQueen(length int64) Queen {
	var q = Queen{
		Length: length,
		Data:   make([]interface{}, length),
	}
	return q
}

//判断是否为空
func (t *Queen) IsEmpty() bool {
	return t.Capacity == 0
}

//判断是否满
func (t *Queen) IsFull() bool {
	return t.Capacity == t.Length
}

//加一个元素
func (t *Queen) Append(element interface{}) bool {
	if t.IsFull() {
		log.Println("队列已满 ，无法加入")
		return false
	}
	t.Data[t.Tail] = element
	t.Tail++
	t.Capacity++
	return true
}

//弹出一个元素，并返回
func (t *Queen) OutElement() interface{} {
	if t.IsEmpty() {
		log.Println("队列为空，无法弹出")
	}
	defer func() {
		t.Capacity--
		t.Head++
	}()
	return t.Data[t.Head]
}

//遍历
func (t *Queen) Each(fn func(node interface{})) {
	for i := t.Head; i < t.Head+t.Capacity; i++ {
		fn(t.Data[i%t.Length])
	}
}

//清空
func (t *Queen) Clcear() bool {
	t.Capacity = 0
	t.Head = 0
	t.Tail = 0
	t.Data = make([]interface{}, t.Length)
	return true
}
