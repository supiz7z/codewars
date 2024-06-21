package slice_chunked

import (
	"codewars/some_functions"
	"fmt"
)

func Start() {
	slice := make([]int, 101)
	n := 13
	minimum := 3
	maximum := 5
	for i := range slice {
		slice[i] = some_functions.RandFill(minimum, maximum)
	}
	mem := shit(slice, n)
	for i := range mem {
		fmt.Println(mem[i])
	}
}

func shit(slice []int, n int) [][]int {
	lenOut := len(slice)/n + 1
	if len(slice)%n == 0 {
		lenOut = len(slice) / n
	}
	mem := make([][]int, lenOut)
	for i := range mem {
		if i == len(mem)-1 && lenOut == len(slice)/n+1 {
			mem[i] = make([]int, len(slice)%n)
			break
		}
		mem[i] = make([]int, n)
	}
	j := 0
	k := 0
	for i := range slice {
		mem[j][k] = slice[i]
		k++
		if (i+1)%n == 0 {
			j++
			k = 0
		}
	}
	return mem
}
