package prime_sextuplets_v1

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

func CheckSum(n int) bool {
	if n%90 == 0 && IsPrime(n/6-8) {
		return true
	}
	return false
}

func MakeAndCheckSextuplet(n int) []int {
	p := n/6 - 8
	if IsPrime(p+4) && IsPrime(p+6) && IsPrime(p+10) && IsPrime(p+12) && IsPrime(p+16) {
		return []int{p, p + 4, p + 6, p + 10, p + 12, p + 16}
	}
	return nil
}

func FindPrimesSextuplet(sum int) []int {
	var arr []int
	for i := sum + 1; i > sum; i++ {
		if CheckSum(i) {
			arr = MakeAndCheckSextuplet(i)
			if arr != nil {
				return arr
			}
		}
	}
	return arr
}

func Start() {
	sum := 26999999
	result := FindPrimesSextuplet(sum)
	fmt.Println(result)
}
