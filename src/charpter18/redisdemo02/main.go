package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type monster struct {
	name  string
	age   int
	skill string
}

func (mon *monster) saveToRedis() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("HMSet", "user01", "name", mon.name, "age", mon.age, "skill", mon.skill)
	if err != nil {
		fmt.Println("存放失败", err)
	} else {
		fmt.Println("存放成功")
	}
	read, err := redis.Strings(conn.Do("HMget", "user01", "name", "age", "skill"))
	if err != nil {
		fmt.Println("读取失败", err)
		return
	}
	for i, v := range read {
		fmt.Printf("read[%v] is %v\n", i, v)
	}
}
func main() {
	var mon monster
	fmt.Println("请输入姓名：")
	fmt.Scanln(&mon.name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&mon.age)
	fmt.Println("请输入技能：")
	fmt.Scanln(&mon.skill)
	mon.saveToRedis()
}
