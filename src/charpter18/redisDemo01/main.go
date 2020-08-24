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
	_, err = conn.Do("Hset", "user01", "name", "john")
	if err != nil {
		fmt.Println("添加失败！")
		return
	}
	_, err = conn.Do("Hset", "user01", "age", "18")
	if err != nil {
		fmt.Println("添加失败！")
		return
	}
	r, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	r1, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok", r, r1)

}
