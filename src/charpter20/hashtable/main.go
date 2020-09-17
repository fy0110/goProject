package main

import (
	"fmt"
	"os"
)

type Emp struct {
	ID   int
	Name string
	Next *Emp
}

//不带表头的 头节点
type EmpLink struct {
	Head *Emp
}
type HashTable struct {
	LinkArr [7]EmpLink
}

//链表插入
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   //辅助指针
	var pre *Emp = nil //还是一个空指针，pre在cur前面
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是一个空链表,给emp找到对应的位置并插入
	for {
		if cur != nil {
			if cur.ID >= emp.ID {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出时，时候到最后一个位置还是插入到中间位置
	emp.Next = cur
	pre.Next = emp

}

//显示当前链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("%d为空\t", no)
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s", no, cur.ID, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//确定插入那个数组的那个位置
func (this *HashTable) Insert(emp *Emp) {
	linNo := this.HashFun(emp.ID)
	this.LinkArr[linNo].Insert(emp)
}

//显示所有雇员
func (this *HashTable) Showall() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
		fmt.Println()
	}
}
func (this *HashTable) HashFun(id int) int {
	return id % 7 //散列值 对应链表的下标
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("==============雇员系统菜单============")
		fmt.Println("=============input 添加雇员==========")
		fmt.Println("=============show 显示雇员==========")
		fmt.Println("=============find 查找雇员==========")
		fmt.Println("=============del 删除雇员==========")
		fmt.Println("=============exit 退出==========")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员姓名")
			fmt.Scanln(&name)
			emp := Emp{
				ID:   id,
				Name: name,
			}
			hashtable.Insert(&emp)
		case "show":
			hashtable.Showall()
		case "find":
			fmt.Println("未做")
		case "exit":
			os.Exit(1)
		}
	}

}
