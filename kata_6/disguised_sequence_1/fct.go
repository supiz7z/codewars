package disguised_sequence_1

import (
	"fmt"
	"math"
)

func Fct(n uint) uint {
	u := math.Pow(2, float64(n))
	return uint(u)
}

func Start() {
	var n uint = 17
	u := Fct(n)
	fmt.Println(u)
}
