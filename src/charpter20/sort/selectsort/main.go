package main

import "fmt"

//选择当前第一个为最大值进行比较
func selectsort(arr *[6]int) {
	for j := 0; j < len(arr)-1; j++ {
		maxnum := arr[j]
		maxIndex := j
		for i := j; i < len(arr)-j; i++ {
			if arr[i] > maxnum {
				maxnum = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != 0 {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
}

func main() {
	arr := [6]int{10, 34, 19, 100, 80, 6}
	selectsort(&arr)
	fmt.Println(arr)
}
