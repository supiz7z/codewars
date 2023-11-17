package reverse_or_rotate

import (
	"fmt"
	"math"
	"strconv"
)

func Reverse(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}

func Rotate(s string) string {
	if len(s) == 1 {
		return s
	}
	b := string(s[0])
	s = s[1:] + b
	return s
}

func StrElements(s string, n int) []string {
	var arr []string
	m := len(s)
	for i := 0; i < m; i++ {
		arr = append(arr, s[:n])
		s = s[n:]
		if len(s) < n {
			break
		}
	}
	return arr
}

func ReverseRotate(s string, n int) string {
	var result string
	var arr []string
	if n <= 0 || s == "" || n > len(s) {
		return ""
	}
	arr = StrElements(s, n)
	for i := 0; i < len(arr); i++ {
		sum := 0
		for _, elem := range arr[i] {
			a, _ := strconv.Atoi(string(elem))
			sum += int(math.Pow(float64(a), 3))
		}
		if sum%2 == 0 {
			arr[i] = Reverse(arr[i])
		} else {
			arr[i] = Rotate(arr[i])
		}
		result += arr[i]
	}
	return result
}

func Start() {
	s := "733049910872815764"
	n := 5
	result := ReverseRotate(s, n)
	fmt.Println(result)
}
