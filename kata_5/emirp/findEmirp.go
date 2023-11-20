package emirp

import (
	"fmt"
	"strconv"
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

func Reverse(n int) int {
	s := strconv.Itoa(n)
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	n, _ = strconv.Atoi(string(arr))
	return n
}

func FindEmirp(n int) [3]int {
	var count, largest, sum int
	for i := 13; i <= n; i++ {
		reverse := Reverse(i)
		if IsPrime(i) && IsPrime(reverse) && i != reverse {
			count++
			sum += i
			if i > largest {
				largest = i
			}
		}
	}
	return [3]int{count, largest, sum}
}

func Start() {
	n := 100
	result := FindEmirp(n)
	fmt.Println(result)
}
