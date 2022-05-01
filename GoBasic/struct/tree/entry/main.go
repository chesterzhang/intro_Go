package main

import (
	"fmt"
	"tree"
)

//相当于继承  TreeNode, 额外实现一个后序遍历的方法
type MyTreeNode struct {
	*tree.TreeNode
}


//重载 Traverse
func (myNode *MyTreeNode ) Traverse() {
	fmt.Println("This is  shadowed method")
}

//添加一个后序遍历
func (myNode *MyTreeNode ) PostOrder(){
	if myNode==nil || myNode.TreeNode==nil {
		return
	}
	left:=MyTreeNode{myNode.TreeNode.Left}
	left.PostOrder()
	right:=MyTreeNode{myNode.TreeNode.Right}
	right.PostOrder()
	myNode.TreeNode.Print()
}

func main() {
	var root tree.TreeNode

	root=tree.TreeNode{Value:3}

	root.Left= new(tree.TreeNode)
	root.Right=&tree.TreeNode{5,nil,nil}

	root.Right.Left=new(tree.TreeNode)
	root.Left.Right=tree.CreateNode(2)

	fmt.Println(root)
	root.Print()

	root.SetVal(10)
	root.Print()

	fmt.Println("============")
	pRoot :=&root
	pRoot.Print()
	pRoot.SetVal(20)
	pRoot.Print()

	fmt.Println("============")
	//     20
	//	 /    \
	//  0       5
	//    \   /
	//	   2  0
	fmt.Println("In order traverse :")
	root.Traverse()

	fmt.Println("MyTreeNode Post order traverse :")

	var root2 =MyTreeNode{&root}
	root2.PostOrder()

	fmt.Println("MyTreeNode  in order traverse :")
	root2.Traverse()

	fmt.Println("MyTreeNode.TreeNode in order traverse :")
	root2.TreeNode.Traverse()

}

