package number_of_integer_partitions

import "fmt"

func dp(targetSum int, maxSummand int, memory [][]int) int {
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

func Partitions(n int) int {
	memory := make([][]int, n+1)
	for i := range memory {
		memory[i] = make([]int, n+1)
	}

	return dp(n, n, memory)
}

func Start() {
	var n = 1000
	total := Partitions(n)
	fmt.Println(total)
}
