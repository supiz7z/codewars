package find_the_odd_int

import "fmt"

func FindOdd(arr []int) int {
	var count, res int
	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] {
				count += 1
			}
		}
		if count%2 == 1 {
			res = arr[i]
			break
		}
	}
	return res
}

func FindOddBestPractice(arr []int) int {
	res := 0
	for _, elem := range arr {
		res ^= elem
	}
	return res
}

func Start() {
	arr := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	//res := FindOdd(arr)
	res := FindOddBestPractice(arr)
	fmt.Println(res)
}
