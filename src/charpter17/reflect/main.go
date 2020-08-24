package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%v完成了减法运算，%v-%v=%v", name,
		c.Num1, c.Num2, c.Num1-c.Num2)
}
func reflectDemo(b interface{}) {
	val := reflect.ValueOf(b)
	// typ := reflect.TypeOf(b)
	fild := val.NumField()
	for i := 0; i < fild; i++ {
		fmt.Println(val.Field(i))
	}
	var namestring []reflect.Value
	namestring = append(namestring, reflect.ValueOf("tom"))
	val.Method(0).Call(namestring)
}
func main() {
	var cal Cal
	cal.Num1 = 8
	cal.Num2 = 3
	reflectDemo(cal)
}
