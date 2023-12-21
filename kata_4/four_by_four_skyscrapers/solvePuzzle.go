package four_by_four_skyscrapers

import "fmt"

func Cutter(arr [][]int) [][]int {
	n := len(arr) - 1
	arr = arr[1:n]
	for i := range arr {
		arr[i] = arr[i][1:n]
	}
	return arr
}

func SolvePuzzle(clues []int) [][]int {
	n := len(clues) / 4
	//создадим матрицу (n+2)*(n+2)
	arr := make([][]int, n+2)
	for i := range arr {
		arr[i] = make([]int, n+2)
	}
	var count int //счетчик для итерации по срезу clues
	//пройдем по периметру матрицы и заполним во внешнем слое количество видимых небоскребов;
	//сразу заполним известные поля ( 4 - где видно 1; 1,2,3,4 - где видно 4 )
	//движемся из левого верхнего угла в правый верхний
	for i := 1; i < len(arr)-1; i++ {
		arr[0][i] = clues[count]
		if clues[count] == 1 {
			arr[1][i] = n
		}
		if clues[count] == n {
			line := 1
			for j := 1; j < len(arr)-1; j++ {
				arr[j][i] = line
				line++
			}
		}
		count++
	}
	//движемся из правого верхнего в правый нижний
	for i := 1; i < len(arr)-1; i++ {
		arr[i][len(arr)-1] = clues[count]
		if clues[count] == 1 {
			arr[i][len(arr)-2] = n
		}
		if clues[count] == n {
			line := 1
			for j := len(arr) - 2; j > 0; j-- {
				arr[i][j] = line
				line++
			}
		}
		count++
	}
	//движемся из правого нижнего в левый нижний
	for i := len(arr) - 2; i > 0; i-- {
		arr[len(arr)-1][i] = clues[count]
		if clues[count] == 1 {
			arr[len(arr)-2][i] = n
		}
		if clues[count] == n {
			line := 1
			for j := len(arr) - 2; j > 0; j-- {
				arr[j][i] = line
				line++
			}
		}
		count++
	}
	//движемся из левого нижнего в левый верхний
	for i := len(arr) - 2; i > 0; i-- {
		arr[i][0] = clues[count]
		if clues[count] == 1 {
			arr[i][1] = n
		}
		if clues[count] == n {
			line := 1
			for j := 1; j < len(arr)-1; j++ {
				arr[i][j] = line
				line++
			}
		}
		count++
	}
	return Cutter(arr)
	//return arr
}

func Start() {
	clues := []int{
		0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0,
	}
	result := SolvePuzzle(clues)
	for i := range result {
		fmt.Println(result[i])
	}
}
