package how_many_numbers

import (
	"fmt"
	"strconv"
)

func IsValid(s string, d int) bool {
	for j := 1; j < len(s); j++ {
		a, _ := strconv.Atoi(string(s[j]))
		b, _ := strconv.Atoi(string(s[j-1]))
		if a <= b || a-b > d {
			return false
		}
	}
	return true
}

func SelNumber(n, d int) int {
	var count int
	for i := 12; i <= n; i++ {
		s := strconv.Itoa(i)
		if IsValid(s, d) {
			count += 1
		}
	}
	return count
}

func Start() {
	n, d := 57492, 8 // 372
	//n, d := 7639, 7 // 245
	total := SelNumber(n, d)
	fmt.Println(total)
}
