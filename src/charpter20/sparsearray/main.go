package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Print(" ", v2)
		}
		fmt.Println()
	}
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, i2 := range chessMap {
		for i3, i4 := range i2 {
			if i4 != 0 {
				valNode := ValNode{
					row: i,
					col: i3,
					val: i4,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	fmt.Println(sparseArr)

	//恢复数组
	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Print(" ", v2)
		}
		fmt.Println()
	}
}
