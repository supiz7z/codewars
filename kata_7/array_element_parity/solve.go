package array_element_parity

import "fmt"

func Solve(arr []int) int {
	var m, count int
	for i := range arr {
		count = 0
		for j := range arr {
			if arr[j] == -arr[i] {
				break
			} else {
				count += 1
				if count == len(arr) {
					m = arr[i]
				}
			}
		}
	}
	return m
}
func Start() {
	arr := []int{-3, 1, 2, 3, -1, -4, -2}
	total := Solve(arr)
	fmt.Println(total)
}
