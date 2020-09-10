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

//362

func main() {
	first := AddBoy(5)
	ShowBoy(first)

}
