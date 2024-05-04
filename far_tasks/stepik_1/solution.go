package stepik_1

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func Solution() int {
	file, err := os.ReadFile("/home/sup/Downloads/24__4_.txt")
	if err != nil {
		fmt.Println(errors.New("nu pizda"))
		return 0
	}
	var count, maxCount int
	for i := 0; i < len(file)-1; i++ {
		if (string(file[i]) == "a" && string(file[i+1]) == "d") || (string(file[i]) == "d" && string(file[i+1]) == "a") {
			count = 0
		}
		count++
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func Start() {
	start := time.Now()
	result := Solution()
	finish := time.Now().Sub(start)
	fmt.Println(result)
	fmt.Println(finish)
}
