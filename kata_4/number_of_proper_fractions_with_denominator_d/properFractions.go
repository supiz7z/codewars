package number_of_proper_fractions_with_denominator_d

import (
	"fmt"
	"math"
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

func ProperFractions(n int) int {
	if n == 1 {
		return 0
	}
	total := 1
	m := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < m; i++ {
		if IsPrime(i) {
			if n%i == 0 {
				total *= i - 1
				n /= i
			}
			for n%i == 0 {
				total *= i
				n /= i
			}
		}
	}
	if n > 1 {
		total *= n - 1
	}
	return total
}

func Start() {
	//n := 15 //8
	n := 25 //20
	//n := 12345678912345
	//n := 1
	total := ProperFractions(n)
	fmt.Println(total)
	//fmt.Println(UniqueFactors(PrimeFactors(n)))
}
