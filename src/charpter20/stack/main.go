package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	maxTop int
	Top    int
	arr    [5]int
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.maxTop-1 {
		return errors.New("栈已满")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) List() (err error) {
	if this.Top == -1 {
		return errors.New("空栈")
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Println(this.arr[i])
	}
	return
}

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		return 0, errors.New("空栈")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

func main() {
	stack := Stack{
		maxTop: 5,
		Top:    -1,
		arr:    [5]int{},
	}
	stack.Push(10)
	stack.Push(11)
	stack.Push(12)
	stack.Push(13)
	stack.Push(14)
	stack.List()
	fmt.Println("=================")
	val, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	fmt.Println("=================")
	stack.List()
}
