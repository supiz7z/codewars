package maximum_subarray_sum

import "fmt"

func MaximumSubarraySum(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		if arr[i] <= 0 {
			continue
		}
		memory := 0
		for j := i; j < len(arr); j++ {
			memory += arr[j]
			if memory > sum {
				sum = memory
			}
		}
	}
	return sum
}

func Start() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := MaximumSubarraySum(arr)
	fmt.Println(result)
}
