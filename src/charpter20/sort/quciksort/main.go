package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quciksort(left int, right int, arr *[80000]int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l < r {
		if arr[l] < pivot {
			l++
		}
		if arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if left < r {
			quciksort(left, r, arr)
		}
		if right > l {
			quciksort(l, right, arr)
		}
	}
}

func main() {
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	starttime := time.Now().Unix()
	quciksort(0, len(arr)-1, &arr)
	//fmt.Println(arr)
	endtime := time.Now().Unix()
	fmt.Println(endtime - starttime)
}
