package main

import "fmt"

func quciksort(left int, right int, arr *[6]int) {
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
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	quciksort(0, 5, &arr)
	fmt.Println(arr)
}
