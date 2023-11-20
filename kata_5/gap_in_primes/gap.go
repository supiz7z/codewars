package gap_in_primes

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i*(i-1) < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func Gap(g, m, n int) []int {
	var primes, result []int
	for i := m; i <= n; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	for i := 0; i < len(primes)-1; i++ {
		if primes[i+1] == primes[i]+g {
			result = append(result, primes[i], primes[i+1])
			return result
		}
	}
	return nil
}

func Start() {
	//g, m, n := 6, 100, 110
	g, m, n := 10, 300, 400
	result := Gap(g, m, n)
	fmt.Println(result)
}
