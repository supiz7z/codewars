package non_decomposable_primes_as_sums

import (
	"fmt"
)

func IsPrime(n int) bool {
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

func LocalPrimes(n int) []int {
	var arr []int
	for i := 3; i <= n; i += 4 {
		if IsPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

var localPrimesForIthNondecompPrime = LocalPrimes(1000000)

func IthNondecompPrime(n int) int {
	return localPrimesForIthNondecompPrime[n-1]
}

func Start() {
	n := 20
	result := IthNondecompPrime(n)
	fmt.Println(result)
}
