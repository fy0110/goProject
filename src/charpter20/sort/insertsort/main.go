package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertsort(arr *[80000]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//if insertIndex+1 != i {
		arr[insertIndex+1] = insertVal
		//}
	}
}

func main() {
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	starttime := time.Now().Unix()
	insertsort(&arr)
	//fmt.Println(arr)
	endtime := time.Now().Unix()
	fmt.Println(endtime - starttime)
}
