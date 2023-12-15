package up_and_down_sorting_for_each_column

import (
	"fmt"
	"sort"
)

func UpDownColSort(matrix [][]int) [][]int {
	var count int
	var numbers []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			numbers = append(numbers, matrix[i][j])
		}
	}
	sort.Ints(numbers)
	rowStart := 0
	rowEnd := len(matrix[0]) - 1
	colStart := 0
	colEnd := len(matrix) - 1
	for rowStart <= rowEnd {
		for i := colStart; i <= colEnd; i++ {
			matrix[i][rowStart] = numbers[count]
			count++
		}
		rowStart++
		if rowStart > rowEnd {
			break
		}
		for i := colEnd; i >= colStart; i-- {
			matrix[i][rowStart] = numbers[count]
			count++
		}
		rowStart++
	}
	return matrix
}

func Start() {
	//arr := [][]int{{-20, -4, -1}, {1, 4, 7}, {8, 10, 12}}
	//arr := [][]int{{1, -1, 4, 1}, {7, -20, 12, 0}, {8, 10, -4, -3}}
	//arr := [][]int{{1, -1, 4, 1}, {0, 12, -20, 7}, {8, 10, -4, -3}}
	arr := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
		{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Println()
	result := UpDownColSort(arr)
	for i := range result {
		fmt.Println(result[i])
	}
}
