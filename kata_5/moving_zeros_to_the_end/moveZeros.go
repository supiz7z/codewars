package moving_zeros_to_the_end

import "fmt"

func RemoveFromSliceByIndex(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func MoveZeros(arr []int) []int {
	var count int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count++
		}
	}
	if count == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(RemoveFromSliceByIndex(arr, i), 0)
			count--
			i--
		}
		if count == 0 {
			break
		}
	}
	return arr
}

func Start() {
	arr := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	result := MoveZeros(arr)
	fmt.Println(result)
}
