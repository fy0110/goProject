package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// fmt.Println("服务起再等待输入", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //等待客户端conn发送信息，若客户端没有发送信息，那么协程就阻塞在这里
		if err != nil {
			fmt.Println("客户端已退出", err)
			return
		}
		fmt.Print(string(buf[:n]))

	}
}
func main() {
	fmt.Println("服务器开始监听。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端链接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err)
		} else {
			fmt.Println(conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
