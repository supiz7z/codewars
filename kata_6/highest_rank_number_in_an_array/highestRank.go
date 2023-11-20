package highest_rank_number_in_an_array

import "fmt"

func HighestRank(arr []int) int {
	var total, result int
	count := make(map[int]int)
	for _, elem := range arr {
		count[elem]++
		if count[elem] > total {
			total = count[elem]
			result = elem
		} else {
			if count[elem] == total {
				if elem > result {
					result = elem
				}
			}
		}
	}
	return result
}
func Start() {
	arr := []int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}
	result := HighestRank(arr)
	fmt.Println(result)
}
