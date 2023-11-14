package dominant_primes

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i*(i-1) < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func Solve(a, b int) int {
	var count, sum int
	for i := 2; i <= b; i++ {
		if IsPrime(i) {
			count += 1
			if i >= a && IsPrime(count) {
				sum += i
			}
		}
	}
	return sum
}

func Start() {
	var a = 4000
	var b = 500000
	total := Solve(a, b)
	fmt.Println(total)
}
