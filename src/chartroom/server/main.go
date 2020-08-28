package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		conn.Read(b)
	}
}
func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端链接成功")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go process(conn)
	}
}
