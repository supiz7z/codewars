package primes_with_even_digits

import "fmt"

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

func DigitsCount(n int) int {
	if n/10 == 0 {
		return 1
	}
	return DigitsCount(n/10) + 1
}

func DigitsCountEven(n int) int {
	var total int
	if (n%10)%2 == 0 {
		total += 1
	}
	if n/10 == 0 {
		return total
	}
	return DigitsCountEven(n/10) + total
}

func F(n int) int {
	var evenCount, evenMax, result int
	for i := n - 1; i > 1; i-- {
		if DigitsCount(i/10) <= evenMax {
			break
		}
		if IsPrime(i) {
			evenCount = DigitsCountEven(i)
			if evenCount > evenMax {
				evenMax = evenCount
				result = i
			}
		}
	}
	return result
}

func Start() {
	n := 5000000 // 1000 <= n <= 5000000
	result := F(n)
	fmt.Println(result)
}
