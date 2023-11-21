package count_ip_addresses

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IpsBetween(start, end string) int {
	var arr1, arr2 []string
	var total1, total2 int
	arr1 = strings.Split(start, ".")
	arr2 = strings.Split(end, ".")
	for i := 0; i < len(arr1); i++ {
		a, _ := strconv.Atoi(arr1[i])
		b, _ := strconv.Atoi(arr2[i])
		total1 += a * int(math.Pow(256, float64(len(arr1)-1-i)))
		total2 += b * int(math.Pow(256, float64(len(arr1)-1-i)))
	}
	return total2 - total1
}

func Start() {
	ip1 := "20.0.0.10"
	ip2 := "20.0.1.0"
	total := IpsBetween(ip1, ip2)
	fmt.Println(total)
}
