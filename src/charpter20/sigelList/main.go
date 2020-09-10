package main

import (
	"fmt"
	"time"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InserHeroNode2(head *HeroNode, newheroNode *HeroNode) {
	//创建一个辅助结点
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newheroNode.no {
			break
		} else if temp.next.no == newheroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("已经存在")
		return
	} else {
		newheroNode.next = temp.next
		temp.next = newheroNode
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		next:     nil,
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		next:     nil,
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
		next:     nil,
	}
	fmt.Println(time.Now())
	InserHeroNode2(head, hero2)
	InserHeroNode2(head, hero3)
	InserHeroNode2(head, hero1)
	ListHeroNode(head)
}
