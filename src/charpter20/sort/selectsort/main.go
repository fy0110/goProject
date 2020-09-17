package main

import (
	"fmt"
	"math/rand"
	"time"
)

//选择当前第一个为最大值进行比较
func selectsort(arr *[80000]int) {
	for j := 0; j < len(arr)-1; j++ {
		maxnum := arr[j]
		maxIndex := j
		for i := j; i < len(arr)-j; i++ {
			if arr[i] > maxnum {
				maxnum = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

func main() {
	//arr := [6]int{10, 34, 19, 100, 80, 6}
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	starttime := time.Now().Unix()
	selectsort(&arr)
	//fmt.Println(arr)
	endtime := time.Now().Unix()
	fmt.Println(endtime - starttime)
}
