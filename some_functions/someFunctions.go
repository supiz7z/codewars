package some_functions

import (
	"math/big"
	"strconv"
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

func IsPrimeOld(n int) bool {
	for i := 2; i*(i-1) < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func IsPrimeBig(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
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

func DigitCount(n int) int {
	if n/10 == 0 {
		return 1
	}
	return DigitCount(n/10) + 1
}

func ReverseString(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}

func ReverseDigitsOfNumber(n int) int {
	s := strconv.Itoa(n)
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	n, _ = strconv.Atoi(string(arr))
	return n
}

func RemoveByIndexFromSlice(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}
