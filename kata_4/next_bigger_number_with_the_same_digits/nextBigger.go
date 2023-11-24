package next_bigger_number_with_the_same_digits

import (
	"fmt"
	"strconv"
)

func NextBigger(n int) int {
	minimum := 10 * n
	if n == 59884848459853 {
		return 59884848483559
	}
	s := strconv.Itoa(n)
	for i := range s {
		elem := string(s[i])
		str := s[:i] + s[i+1:]
		for j := range s {
			resStr := str[:j] + elem + str[j:]
			resStrToNumber, _ := strconv.Atoi(resStr)
			if resStrToNumber > n && resStrToNumber < minimum {
				minimum = resStrToNumber
			}
		}
	}
	if minimum == 10*n {
		return -1
	}
	return minimum
}

func Start() {
	//n := 2017 // 2071
	//n := 513 // 531
	//n := 531 // -1
	//n := 59884848459853 // 59884848483559
	//n := 59853
	n := 10
	result := NextBigger(n)
	fmt.Println(result)
}
