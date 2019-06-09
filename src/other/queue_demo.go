package main

import (
	"fmt"
	"other/queue"
)

func main() {
	q := queue.Queue{}
	q.Put("queue_1")
	q.Put("queue_2")
	ss := map[string]interface{}{
		"test": 1,
	}
	q.Put(ss)
	travselQueue(&q)

	q.Pop() //出队列
	travselQueue(&q)
}

func travselQueue(q *queue.Queue) {
	fmt.Println("-------queue----begin-------------")
	//遍历
	head := q.GetHeadNode()
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
	fmt.Println("-------queue------end-------")
}
