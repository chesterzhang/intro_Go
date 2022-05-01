package tree

import "fmt"

type TreeNode struct{
	Value int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode  {
	return &TreeNode{Value:value} //必须返回一个指针, 否则就返回了函数内部的一个局部变量
}

//相当于 print(node treeNode) 函数, node 传值, go 所有参数都是传值, 要引用则需要指针
func (node TreeNode) Print()  {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetVal(value int)  {
	 node.Value=value
}


//中序遍历
func (node *TreeNode) Traverse(){
	if node==nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}


