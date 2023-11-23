package race_ceremony

import "fmt"

func RacePodium(n int) [3]int {
	//arr := [3]int{n-3,2,1}

	return [3]int{0, 0, 0}
}

func Start() {
	n := 100000 // 33334, 33335, 33331
	result := RacePodium(n)
	fmt.Println(result)
}
