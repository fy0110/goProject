package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("set", "name", "jerry")
	if err != nil {
		fmt.Println("添加失败！")
		return
	}
	fmt.Println("ok")

}
