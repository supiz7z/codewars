package square_and_cube_of_a_number_become_prime_when_reversed

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

func ReverseDigitsOfNumber(n int) int {
	var result int
	for n > 0 {
		memory := n % 10
		result = result*10 + memory
		n /= 10
	}
	return result
}

func ToSq(n int) int {
	return int(math.Pow(float64(n), 2))
}
func ToCube(n int) int {
	return int(math.Pow(float64(n), 3))
}

func LocalPrime(n int) []int {
	var arr []int
	for i := 89; i <= n; i++ { // 2147483647
		if IsPrime(ReverseDigitsOfNumber(ToSq(i))) && IsPrime(ReverseDigitsOfNumber(ToCube(i))) {
			arr = append(arr, i)
		}
	}
	return arr
}

var localPrimes = LocalPrime(57200)

func SqCubRevPrime(n int) int {
	return localPrimes[n-1]
}

func Start() {
	n := 250
	result := SqCubRevPrime(n)
	fmt.Println(result)
}
