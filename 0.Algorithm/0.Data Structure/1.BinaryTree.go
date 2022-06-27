package main

import "fmt"

func main() {

	tree := &Tree{}
	tree.Add(123)
	tree.Add(12)
	tree.Add(1)
	tree.Add(1234)
	tree.Add(12345)
	tree.Add(123456)

	fmt.Println("-----------前序遍历")
	tree.PreOrderTraverse(tree.Root)
	fmt.Println("-----------中序遍历")
	tree.InOrderTraverse(tree.Root)
	fmt.Println("-----------后序遍历")
	tree.PostOrderTraverse(tree.Root)

}

// TreeNode leaf结构
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// Tree 二叉树的实现(根节点)
type Tree struct {
	Root *TreeNode
}

func (tree *Tree) Add(data int) {
	var queue []*TreeNode
	newNode := &TreeNode{Data: data}
	if tree.Root == nil {
		tree.Root = newNode
	} else {
		queue = append(queue, tree.Root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)

			if data > cur.Data {
				//右树添加节点
				if cur.Right == nil {
					cur.Right = newNode
				} else {
					queue = append(queue, cur.Right)
				}
			} else {
				//左树添加节点
				if cur.Left == nil {
					cur.Left = newNode
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}
	}
}

// PreOrderTraverse 前序遍历 根 -> 左 -> 右
func (tree *Tree) PreOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		fmt.Println(node.Data)
		tree.PreOrderTraverse(node.Left)
		tree.PreOrderTraverse(node.Right)
	}
}

// InOrderTraverse 中序遍历 左 -> 根 -> 右
func (tree *Tree) InOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.InOrderTraverse(node.Left)
		fmt.Println(node.Data)
		tree.InOrderTraverse(node.Right)
	}
}

// PostOrderTraverse 后序遍历 左 -> 右 -> 根
func (tree *Tree) PostOrderTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.PostOrderTraverse(node.Left)
		tree.PostOrderTraverse(node.Right)
		fmt.Println(node.Data)
	}
}

func Delete() {

}

func Update() {

}

func Find() {

}

func IsExists(){

}

func GetMax(){

}

func GetMin(){

}

func DelMax(){

}

func DelMin(){
	
}