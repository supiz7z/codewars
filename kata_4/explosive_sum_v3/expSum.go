package explosive_sum_v3

import "fmt"

func ExpSum(n uint64) uint64 {
	memory := make([][]uint64, n+1)
	for i := range memory {
		memory[i] = make([]uint64, n+1)
	}

	return dp(n, n, memory)
}

func dp(targetSum uint64, maxSummand uint64, memory [][]uint64) uint64 {
	if targetSum == 0 {
		return 1
	}
	if maxSummand == 0 {
		return 0
	}

	if memory[targetSum][maxSummand] == 0 {
		memory[targetSum][maxSummand] = dp(targetSum, maxSummand-1, memory)
		if targetSum >= maxSummand {
			memory[targetSum][maxSummand] += dp(targetSum-maxSummand, maxSummand, memory)
		}
	}

	return memory[targetSum][maxSummand]
}

func Start() {
	var n uint64 = 1000
	total := ExpSum(n)
	fmt.Println(total)
}
