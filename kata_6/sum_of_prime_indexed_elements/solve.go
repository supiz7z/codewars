package sum_of_prime_indexed_elements

import "fmt"

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

func Solve(arr []int) int {
	var total int
	for i := 0; i < len(arr); i++ {
		if IsPrime(i) {
			total += arr[i]
		}
	}
	return total
}

func Start() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	total := Solve(arr)
	fmt.Println(total)
}
