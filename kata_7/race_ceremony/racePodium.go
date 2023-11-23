package race_ceremony

import (
	"fmt"
)

func RacePodium(n int) [3]int {
	if n%3 == 0 {
		return [3]int{n / 3, n/3 + 1, n/3 - 1}
	}
	if n%3 == 1 {
		if n/3 == 2 {
			return [3]int{n / 3, n/3 + 2, n/3 - 1}
		}
		if n/3 == 3 {
			return [3]int{n/3 + 1, n/3 + 2, n/3 - 2}
		}
		return [3]int{n/3 + 1, n/3 + 2, n/3 - 2}
	}
	if n%3 == 2 {
		return [3]int{n/3 + 1, n/3 + 2, n/3 - 1}
	}
	return [3]int{0, 0, 0}
}

func Start() {
	n := 16 // 33334, 33335, 33331
	result := RacePodium(n)
	fmt.Println(result)
}
