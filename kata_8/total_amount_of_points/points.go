package total_amount_of_points

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	var total, count int
	for i := range games {
		games[i] = strings.Replace(games[i], ":", "", -1)
		count, _ = strconv.Atoi(games[i])
		if count/10 > count%10 {
			total += 3
		} else {
			if count/10 == count%10 {
				total += 1
			}
		}
	}
	return total
}

func Start() {
	arr := []string{"3:1", "2:2", "0:1", "0:3", "1:2", "4:3", "5:3", "3:3"}
	total := Points(arr)
	fmt.Println(total)
}
