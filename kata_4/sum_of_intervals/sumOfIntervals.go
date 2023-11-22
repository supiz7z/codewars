package sum_of_intervals

import "fmt"

func CrossNew(arr1, arr2 [2]int) bool {
	if arr1[1] > arr2[0] && arr1[0] < arr2[1] {
		return true
	}
	return false
}

func SumOfIntervals(arr [][2]int) int {
	if len(arr) == 1 {
		return arr[0][1] - arr[0][0]
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Println(arr[i], arr[j])
			if CrossNew(arr[i], arr[j]) {
				arr[i][0] = min(arr[i][0], arr[j][0])
				arr[i][1] = max(arr[i][1], arr[j][1])
				arr = append(arr[:j], arr[j+1:]...)
				i--
				break
			}
		}

	}
	fmt.Println(arr)
	var count int
	for i := 0; i < len(arr); i++ {
		count += arr[i][1] - arr[i][0]
	}
	return count
}

func Start() {
	//arr := [][2]int{{1, 4}, {7, 10}, {3, 5}} // 7
	//arr := [][2]int{{0, 20}, {-100_000_000, 10}, {30, 40}} // 100.000.030
	//arr := [][2]int{{-1_000_000_000, 1_000_000_000}} // 2.000.000.000
	arr := [][2]int{{1, 3}, {-7, -3}} // 6
	result := SumOfIntervals(arr)
	fmt.Println(result)
}
