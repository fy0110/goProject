package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环形
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}
func DelCatNode(head *CatNode, no int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		return head
	}
	if temp.next == head {
		temp.next = nil
		return head
	}
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	flag := true
	for {
		if temp.next == head {
			break
		}
		if temp.no == no {
			if temp == head {
				head = head.next
			}
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == no {
			helper.next = temp.next //形成新的头节点
		} else {
			fmt.Println("没有这只猫")
		}

	}
	return head
}
func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Println(temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
		next: nil,
	}
	cat2 := &CatNode{
		no:   2,
		name: "def",
		next: nil,
	}
	cat3 := &CatNode{
		no:   3,
		name: "abc",
		next: nil,
	}
	cat4 := &CatNode{
		no:   4,
		name: "jack",
		next: nil,
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	InsertCatNode(head, cat4)
	ListCircleLink(head)
	fmt.Println("=======================")
	head1 := DelCatNode(head, 5)
	ListCircleLink(head1)
}
