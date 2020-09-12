package main

import "fmt"

func selectsort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		maxnum := arr[j]
		maxIndex := j
		for i := j; i < len(arr); i++ {
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
	arr := [5]int{10, 34, 19, 100, 80}
	selectsort(&arr)
	fmt.Println(arr)
}
