package explosive_sum_v1

import "fmt"

func ExpSum(n uint64) uint64 {
	totalWays := uint64(0)
	for i := uint64(1); i <= n; i++ {
		totalWays += ExpSumDfs(i, 0, n)
	}
	return totalWays
}

func ExpSumDfs(node uint64, sumSoFar uint64, target uint64) uint64 {
	sumWithNode := sumSoFar + node
	if sumWithNode == target {
		return 1
	}

	ways := uint64(0)
	for child := node; sumWithNode+child <= target; child++ {
		ways += ExpSumDfs(child, sumWithNode, target)
	}
	return ways
}

func Start() {
	var n uint64 = 10
	count := ExpSum(n)
	fmt.Println(count)
}
