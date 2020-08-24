package main

import "fmt"

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("------------欢迎登录多人聊天系统-------------------")
		fmt.Println("\t\t1.登录聊天室")
		fmt.Println("\t\t2.注册用户")
		fmt.Println("\t\t3.退出聊天系统")
		fmt.Println("\t\t请选择（1-3）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出聊天系统")
			loop = false
		default:
			fmt.Println("选择错误，请重新选择！")
		}

	}
}
