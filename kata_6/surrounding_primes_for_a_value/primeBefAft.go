package surrounding_primes_for_a_value

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i*(i-1) < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func PrimeBefAft(n int) [2]int {
	var arr [2]int
	for i := n - 1; i > 1; i-- {
		if IsPrime(i) {
			arr[0] = i
			break
		}
	}
	for i := n + 1; i > n; i++ {
		if IsPrime(i) {
			arr[1] = i
			break
		}
	}
	return arr
}

func Start() {
	n := 97 // {89, 101}
	primeBefAft := PrimeBefAft(n)
	fmt.Println(primeBefAft)
}
