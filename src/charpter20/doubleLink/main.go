package main

import (
	"fmt"
	"time"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

func InserHeroNode(head *HeroNode, newheroNode *HeroNode) {
	//创建一个辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newheroNode
	newheroNode.pre = temp

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
		newheroNode.next = temp.next //保证后面的链表不丢失
		newheroNode.pre = temp       //连接到当前temp的链表
		if temp.next != nil {
			temp.next.pre = newheroNode //后面链表的头连接到新节点
		}
		temp.next = newheroNode //新节点连接到temp的尾巴
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
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		} else {
			fmt.Println("未找到")
		}
	}
}
func main() {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		pre:      nil,
		next:     nil,
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		pre:      nil,
		next:     nil,
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
		pre:      nil,
		next:     nil,
	}
	fmt.Println(time.Now())
	fmt.Println("==========================")
	InserHeroNode2(head, hero2)
	InserHeroNode2(head, hero3)
	InserHeroNode2(head, hero1)
	ListHeroNode(head)
	fmt.Println("==========================")
	DelHeroNode(head, 2)
	ListHeroNode(head)
}
