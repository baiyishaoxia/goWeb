package tree

import "fmt"

type Object interface{}

type TreeNode struct {
	Data       Object
	LeftChild  *TreeNode
	RightChild *TreeNode
}

//(完全)二叉树结构
type Tree struct {
	RootNode *TreeNode
}

//追加元素 (广度优先，即按层级遍历后添加)
func (this *Tree) Add(object Object) {
	node := &TreeNode{Data: object}
	if this.RootNode == nil {
		this.RootNode = node
		return
	}
	queue := []*TreeNode{this.RootNode}
	for len(queue) != 0 {
		cur_node := queue[0]
		queue = queue[1:]

		if cur_node.LeftChild == nil {
			cur_node.LeftChild = node
			return
		} else {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild == nil {
			cur_node.RightChild = node
			return
		} else {
			queue = append(queue, cur_node.RightChild)
		}
	}
}

//广度遍历
func (this *Tree) BreadthTravel() {

	if this.RootNode == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, this.RootNode)

	for len(queue) != 0 {
		//fmt.Printf("len(queue):%d\n", len(queue))
		cur_node := queue[0]
		queue = queue[1:]

		fmt.Printf("%v  ", cur_node.Data)

		if cur_node.LeftChild != nil {
			queue = append(queue, cur_node.LeftChild)
		}
		if cur_node.RightChild != nil {
			queue = append(queue, cur_node.RightChild)
		}
	}

}

/*
深度遍历:
1.先序遍历:根->左->右
2.中序遍历:左->中->右
3.后序遍历:左->右->根
*/

//先序遍历
func (this *Tree) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v  ", node.Data)

	//if node.LeftChild != nil {
	this.PreOrder(node.LeftChild)
	//}
	//if node.RightChild != nil {
	this.PreOrder(node.RightChild)
	//}
}

//中序遍历
func (this *Tree) InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.InOrder(node.LeftChild)
	fmt.Printf("%v  ", node.Data)
	this.InOrder(node.RightChild)
}

func (this *Tree) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.PostOrder(node.LeftChild)
	this.PostOrder(node.RightChild)
	fmt.Printf("%v  ", node.Data)
}
