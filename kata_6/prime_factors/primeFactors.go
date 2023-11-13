package prime_factors

import (
	"fmt"
	"math/big"
)

func PrimeFactors(n int) []int {
	var arr []int
	for i := 2; n >= 2; {
		if big.NewInt(int64(n)).ProbablyPrime(0) {
			arr = append(arr, n)
			return arr
		}
		if n%i == 0 {
			n /= i
			arr = append(arr, i)
		} else {
			i += 1
		}
	}
	return arr
}

func Start() {
	var n = 2147483648
	//var n = 2147483647
	factors := PrimeFactors(n)
	fmt.Println(factors)
}
