package prime_factors

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*(i-1) < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func PrimeFactors(n int) []int {
	var arr []int
	for i := 2; n >= 2; {
		if IsPrime(n) {
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
	factors := PrimeFactors(n)
	fmt.Println(factors)
}
