package snail

import "fmt"

func Snail(matrix [][]int) []int {
	if len(matrix[0]) == 0 {
		return make([]int, 0)
	}
	var arr []int
	rowStart := 0
	rowEnd := len(matrix) - 1
	colStart := 0
	colEnd := len(matrix) - 1
	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			arr = append(arr, matrix[rowStart][i])
		}
		rowStart++
		for i := rowStart; i <= rowEnd; i++ {
			arr = append(arr, matrix[i][colEnd])
		}
		colEnd--
		for i := colEnd; i >= colStart; i-- {
			arr = append(arr, matrix[rowEnd][i])
		}
		rowEnd--
		for i := rowEnd; i >= rowStart; i-- {
			arr = append(arr, matrix[i][colStart])
		}
		colStart++
	}
	return arr
}

func Start() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // {1, 2, 3, 6, 9, 8, 7, 4, 5}
	//matrix := [][]int{{}, {}, {}}
	for i := range matrix {
		fmt.Println(matrix[i])
	}
	fmt.Println()
	result := Snail(matrix)
	fmt.Println(result)
}
