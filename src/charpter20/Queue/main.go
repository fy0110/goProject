package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("队列已满")
	} else {
		this.rear++
		this.array[this.rear] = val
		fmt.Println("添加成功")
	}
	return
}
func (this *Queue) GetQueue() (err error) {
	if this.front == this.rear {
		return errors.New("队列已空")
	} else {
		this.front++
		fmt.Println(this.array[this.front])
	}
	return
}
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, this.array[i])
	}
}

func main() {
	queue := Queue{
		maxSize: 5,
		array:   [5]int{},
		front:   -1,
		rear:    -1,
	}
	var key string
	var flag bool = true
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
			err := queue.AddQueue(value)
			if err != nil {
				fmt.Println(err)
			}
		case "get":
			err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			flag = false
		default:
			println("输入有误，请重新输入！")
		}

	}

}
