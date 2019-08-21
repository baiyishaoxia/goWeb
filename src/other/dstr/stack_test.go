package dstr

import(
	"testing"
)

func TestStack_Init(t *testing.T)  {
	stack := new(Stack)
	stack.Init()
	t.Log("stack init success")
}

func TestStack_Push(t *testing.T){
	stack := new(Stack)
	stack.Init()
	b := stack.Push(1)
	if !b {
		t.Error("stack push failed")
		return
	}
	t.Log("stack push success")
	data := stack.Peek()
	var (
		ok bool
		num int
	)
	if num, ok = data.(int); ok && num == 1{
		t.Log("stack push and peek success")
		return
	}
	t.Error("stack push success but peek failed")
}

func TestStack_Pop(t *testing.T){
	stack := new(Stack)
	stack.Init()
	d1 := stack.Pop()
	if d1 != nil{
		t.Error("empty stack pop error")
		return
	}
	t.Log("empty stack pop success")

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	d2 := stack.Pop()
	var (
		ok bool
		num int
	)
	if num, ok = d2.(int); ok && num == 3 && stack.Size() == 2{
		t.Log("stack pop success")
		return
	}
	t.Error("stack pop failed")
}

func TestStack_Peek(t *testing.T){
	stack := new(Stack)
	stack.Init()
	d := stack.Peek()
	if d == nil {
		t.Log("empty stack peek success")
		return
	}
	t.Error("empty stack peek fail")
}