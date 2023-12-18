package beeramid

import "fmt"

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 {
		return 0
	}
	var check int
	count := int(float64(bonus) / price)
	for i := 1; ; i++ {
		check += i * i
		if check > count {
			return i - 1
		}
	}
}

func Start() {
	bonus := 1500
	price := 2.0
	result := Beeramid(bonus, price)
	fmt.Println(result)
}
