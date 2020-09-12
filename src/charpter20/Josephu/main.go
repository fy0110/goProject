package main

import "fmt"

type Boy struct {
	no   int
	next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no:   i,
			next: nil,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = curBoy.next
			curBoy.next = first
		}
	}
	return first
}

func ShowBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("当前为konlianb")
		return
	}
	curBoy := first
	for {
		fmt.Println(curBoy.no)
		if curBoy.next == first {
			break
		}
		curBoy = curBoy.next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {
	if first.next == nil {
		fmt.Println("空链表")
		return
	}
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	//让first移动到startNo
	for i := 0; i < startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	//开始数countNum ，然后删除first
	for {
		for i := 0; i < countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Println(first.no)
		first = first.next
		tail.next = first
		//tail.next = first.next//这样会丢失first的位置
		if tail == first {
			break
		}
	}
	fmt.Println(first.no)

}

func main() {
	first := AddBoy(500)
	ShowBoy(first)
	fmt.Println("=============================")
	PlayGame(first, 20, 31)
}
