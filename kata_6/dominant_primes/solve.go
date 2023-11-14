package dominant_primes

import (
	"fmt"
)

func IsPrime(n int) bool {
	//if n < 2 {
	//	return false
	//}
	//for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
	//	if n%i == 0 {
	//		return false
	//	}
	//}
	//return true
	for i := 2; i*i < n+i; i++ { // этот алгоритм взят из ответов, мой выше закомментирован. Задача была сдана с моим)
		if n%i == 0 {
			return false
		}
	}
	return true
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
