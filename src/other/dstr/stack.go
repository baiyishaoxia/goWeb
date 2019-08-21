package dstr

// 栈信息
type Stack struct {
	list *SingleList
}

// Init 初始化栈
func (s *Stack) Init()  {
	s.list = new(SingleList)
	s.list.Init()
}

// Push 压入栈
func (s *Stack)Push(data interface{}) bool {
    node := &SingleNode{
		Data: data,
	}
	return s.list.Insert(0, node)
}

// Pop 压出栈
func (s *Stack)Pop() interface{}{
	node := s.list.Get(0)
	if node != nil {
		s.list.Delete(0)
		return node.Data
	}
	return nil
}

// Peek 查看栈顶元素
func (s *Stack)Peek() interface{}{
	node := s.list.Get(0)
	if node != nil {
		return node.Data
	}
	return nil
}

// Size 获取栈的长度
func (s *Stack)Size()uint{
	return s.list.Size
}