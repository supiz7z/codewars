package explosive_sum_v2

import "fmt"

func ExpSum(n uint64) uint64 {
	res := make([]uint64, n+1)
	res[0] = 1

	for num := uint64(1); num <= n; num++ {
		for i := num; i <= n; i++ {
			res[i] += res[i-num]
		}
	}
	return res[n]
}

func Start() {
	var n uint64 = 1000
	count := ExpSum(n)
	fmt.Println(count)
}
