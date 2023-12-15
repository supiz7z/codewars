package numbers_of_measurements_to_spot_the_counterfeit_coin

import "fmt"

func HowManyMeasurements(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 || n == 3 {
		return 1
	}
	switch n % 3 {
	case 0:
		return HowManyMeasurements(n/3) + 1
	case 1:
		return HowManyMeasurements((n+2)/3) + 1
	case 2:
		return HowManyMeasurements((n+1)/3) + 1
	}
	return 0
}

func Start() {
	n := 100
	total := HowManyMeasurements(n)
	fmt.Println(total)
}
