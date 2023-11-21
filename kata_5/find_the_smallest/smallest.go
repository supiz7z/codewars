package find_the_smallest

import (
	"fmt"
	"strconv"
)

func Smallest(n int64) []int64 {
	minimum := n
	s := strconv.FormatInt(n, 10)
	var result = []int64{n, 0, 0}
	for i := range s {
		elem := string(s[i])
		str := s[:i] + s[i+1:]
		for j := range s {
			resStr := str[:j] + elem + str[j:]
			resStrToNumber, _ := strconv.ParseInt(resStr, 10, 64)
			if resStrToNumber < minimum {
				minimum = resStrToNumber
				result = []int64{minimum, int64(i), int64(j)}
			}
		}
	}
	return result
}

func Start() {
	var n int64 = 10000 //209917
	result := Smallest(n)
	fmt.Println(result)
}
