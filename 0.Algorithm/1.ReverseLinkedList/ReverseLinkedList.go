package main

import "fmt"

//链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//创建链表
	head := new(ListNode)
	head.Val = 0

	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 2

	node3 := new(ListNode)
	node3.Val = 3

	node4 := new(ListNode)
	node4.Val = 4

	node5 := new(ListNode)
	node5.Val = 5

	//连接链表
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	Println(ReverseList(head))
}

func ReverseList(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	var pre *ListNode

	//反转链表
	for node != nil {
		temp := node.Next
		node.Next = pre
		pre = node
		node = temp
	}

	return pre
}

//遍历打印链表值
func Println(list *ListNode) {
	for list != nil {
		fmt.Printf("val:%v ", list.Val)
		list = list.Next
	}
}
