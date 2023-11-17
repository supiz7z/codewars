package are_they_the_same

import (
	"fmt"
)

func RemoveByIndex(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func Comp(arr1, arr2 []int) bool {
	var count int
	if arr1 == nil || arr2 == nil {
		return false
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr2[i] == arr1[j]*arr1[j] {
				arr1 = RemoveByIndex(arr1, j)
				count++
				break
			}
		}
	}
	return count == len(arr2)
}

func Start() {
	var arr1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var arr2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19} // true
	//var arr1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var arr2 = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 36}
	//var arr1 []int
	//var arr2 []int
	check := Comp(arr1, arr2)
	fmt.Println(check)
}
