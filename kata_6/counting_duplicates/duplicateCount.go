package counting_duplicates

import (
	"fmt"
	"strings"
)

func DuplicateCount(s string) int {
	var count int
	charCount := make(map[rune]int)
	for _, elem := range strings.ToLower(s) {
		charCount[elem]++
		if charCount[elem] == 2 {
			count++
		}
	}
	return count
}

func Start() {
	s := "qqwwweeeerrrtttyyuuuii"
	count := DuplicateCount(s)
	fmt.Println(count)
}
