package main

import (
	"fmt"
	"other/queue"
	"testing"
)

//go语言数据结构 环形队列
func TestQueen(t *testing.T) {
	var testLength = int64(4)
	q := queue.MakeQueen(testLength)
	if q.Length != testLength {
		t.Error("MakeQueen(4)的容量不是4")
	}
	q.Append(10)
	q.Append(12)
	q.Append(14)
	q.Append(16)
	//q.Append(18)

	q.OutElement()
	//fmt.Println(q.OutElement())
	if q.Capacity != 3 {
		t.Error("队队长度不正确")
	}

	q.Each(func(node interface{}) {
		fmt.Println(node)
	})

	q.Clcear()
	if q.Capacity != 0 {
		t.Error("queen的长度不为0")
	}
	q.Each(func(node interface{}) {
		fmt.Println(node)
	})

	q.Append("B")
	q.Append('A')

	q.Each(func(node interface{}) {
		fmt.Println(node)
	})
}
