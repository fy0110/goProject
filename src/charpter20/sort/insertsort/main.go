package main

import "fmt"

func insertsort(arr *[7]int) {
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
	arr := [7]int{23, 0, 12, 56, 34, -1, 55}
	insertsort(&arr)
	fmt.Println(arr)
}
