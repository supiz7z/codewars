package positions_average

import (
	"fmt"
	"strings"
)

func PosAverage(s string) float64 {
	arr := strings.Split(s, ", ")
	var count float64
	var total int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			total += 1 //total можно посчитать сразу как (len(arr)*(len(arr)-1))/2, не заводя счетчик
			for k := 0; k < len(arr[0]); k++ {
				if string(rune(arr[i][k])) == string(rune(arr[j][k])) {
					count += 1
				}
			}
		}
	}
	return count / float64(total*len(arr[0])) * 100
}

func Start() {
	var s = "444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490"
	//var s = "466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"
	//var s = "6900690040, 4690606946, 9990494604"
	posAverage := PosAverage(s)
	fmt.Printf("вероятность равна %.10f процентов", posAverage)
}
