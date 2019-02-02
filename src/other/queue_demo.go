package main

import (
	"fmt"
	"other/queue"
)

func main() {
	q := queue.Queue{}
	q.Put("queue_1")
	q.Put("queue_2")
	travselQueue(&q)

	q.Pop()
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
