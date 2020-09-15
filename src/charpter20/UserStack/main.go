package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	maxTop int
	Top    int
	arr    [20]int
}

//判断是不是操作符
func (this *Stack) isOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
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
	//numstack := &Stack{
	//	maxTop: 20,
	//	Top:    -1,
	//	arr:    [20]int{},
	//}
	//operstack := &Stack{
	//	maxTop: 20,
	//	Top:    -1,
	//	arr:    [20]int{},
	//}
	//exp:="3+2*6-2"
	//index := 0
	//for  {
	//	ch:=exp[index:index+1]
	//
	//}
}
