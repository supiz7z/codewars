package prime_sextuplets_v2

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

func FindPrimesSextuplet(n int) []int {
	var total, index int
	for i := 0; total <= n; i++ {
		index = i
		if IsPrime(30*i+7) && IsPrime(30*i+11) && IsPrime(30*i+13) && IsPrime(30*i+17) && IsPrime(30*i+19) && IsPrime(30*i+23) {
			total = 180*i + 90
		}
	}
	return []int{30*index + 7, 30*index + 11, 30*index + 13, 30*index + 17, 30*index + 19, 30*index + 23}

}

func FindPrimesSextupletAccept(n int) []int {
	var sum int
	arr := []int{7, 97, 16057, 19417, 43777, 1091257, 1615837, 1954357, 2822707, 2839927, 3243337, 3400207, 6005887}
	for i := 0; i < len(arr); i++ {
		sum = arr[i]*6 + 48
		if sum > n {
			return []int{arr[i], arr[i] + 4, arr[i] + 6, arr[i] + 10, arr[i] + 12, arr[i] + 16}
		}
	}
	return nil
}

func Start() {
	sum := 26999999
	result := FindPrimesSextuplet(sum)
	fmt.Println(result)
}
