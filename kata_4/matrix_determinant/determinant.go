package matrix_determinant

import (
	"fmt"
)

func MakeDowngradeMatrix(arr [][]int, index int) [][]int {
	n := len(arr)
	memory := make([][]int, n-1)
	for i := range memory {
		memory[i] = make([]int, n)
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			memory[i-1][j] = arr[i][j]
		}
	}
	for i := 0; i < len(memory); i++ {
		memory[i] = append(memory[i][:index], memory[i][index+1:]...)
	}
	return memory
}

func PlusOrMinus(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}

func Determinant(arr [][]int) int {
	var det int
	if len(arr) == 1 {
		return arr[0][0]
	}
	if len(arr) == 2 {
		return arr[0][0]*arr[1][1] - arr[0][1]*arr[1][0]
	}
	for i := 0; i < len(arr); i++ {
		det += arr[0][i] * PlusOrMinus(i) * Determinant(MakeDowngradeMatrix(arr, i))
	}
	return det
}

func Start() {
	//arr := [][]int{{10}}
	//arr := [][]int{{1, 3}, {2, 5}}
	//arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr := [][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}
	det := Determinant(arr)
	fmt.Println(det)

}
