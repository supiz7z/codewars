package make_a_spiral

import "fmt"

func Spiralize(n int) [][]int {
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	rowStart := 0
	rowEnd := n - 1
	colStart := 0
	colEnd := n - 1
	for rowStart <= rowEnd && colStart <= colEnd {
		if rowStart != 0 {
			arr[rowStart][colStart-1] = 1
		}

		for i := colStart; i <= colEnd; i++ {
			arr[rowStart][i] = 1
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			arr[i][colEnd] = 1
		}
		colEnd--
		if colStart >= colEnd {
			break
		}
		for i := colEnd; i >= colStart; i-- {
			arr[rowEnd][i] = 1
		}
		rowEnd--

		for i := rowEnd; i >= rowStart+1; i-- {
			arr[i][colStart] = 1
		}
		rowStart++
		colStart += 2
		rowEnd--
		colEnd--

	}
	return arr
}

func Start() {
	//n := 5 // [[1,1,1,1,1],[0,0,0,0,1],[1,1,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
	n := 6
	result := Spiralize(n)
	for i := range result {
		fmt.Println(result[i])
	}
}
