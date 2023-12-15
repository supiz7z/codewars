package the_spiraling_box

import "fmt"

func CreateBox(width, length int) [][]int {
	arr := make([][]int, length)
	for i := range arr {
		arr[i] = make([]int, width)
	}
	rowStart := 0
	rowEnd := width - 1
	colStart := 0
	colEnd := length - 1
	jopa := 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := rowStart; i <= rowEnd; i++ {
			arr[colStart][i] = jopa
		}
		for i := colStart; i <= colEnd; i++ {
			arr[i][rowEnd] = jopa
		}
		for i := rowEnd - 1; i >= rowStart; i-- {
			arr[colEnd][i] = jopa
		}
		for i := colEnd; i > colStart; i-- {
			arr[i][rowStart] = jopa
		}
		rowStart++
		rowEnd--
		colStart++
		colEnd--
		jopa++
	}
	return arr
}
func Start() {
	w, l := 19, 19
	result := CreateBox(w, l)
	for i := range result {
		fmt.Println(result[i])
	}
	//fmt.Println(result)
}
