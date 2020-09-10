package main

import (
	"errors"
	"fmt"
)

type circleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (this *circleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	//此时的this。tail在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}
func (this *circleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列已空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}
func (this *circleQueue) ListQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("队列已空")
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}
func (this *circleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}
func (this *circleQueue) IsEmpty() bool {
	return this.tail == this.head
}
func (this *circleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}
func main() {
	queue := circleQueue{
		maxSize: 5,
		array:   [5]int{},
		tail:    0,
		head:    0,
	}
	var key string
	var flag = true
	for flag {
		fmt.Println("1.输入add 表示添加数据")
		fmt.Println("2.输入get 表示获取数据")
		fmt.Println("3.输入show 表示展示数据")
		fmt.Println("4.输入exit 表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			value := 0
			fmt.Println("请输入要添加的数据：")
			fmt.Scanln(&value)
			err := queue.Push(value)
			if err != nil {
				fmt.Println(err)
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(val)
		case "show":
			queue.ListQueue()
		case "exit":
			flag = false
		default:
			println("输入有误，请重新输入！")
		}
	}
}
