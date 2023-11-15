package consonant_value

import (
	"fmt"
	"strings"
)

func Solve(s string) int {
	var res, count int32
	var split = "aeiou"
	for _, elem := range s {
		if strings.Contains(split, string(elem)) {
			s = strings.Replace(s, string(elem), "*", -1)
		}
	}
	strArr := strings.Split(s, "*")
	for i := 0; i < len(strArr); i++ {
		count = 0
		for _, elem := range strArr[i] {
			count += elem - 96
		}
		if count > res {
			res = count
		}
	}
	return int(res)
}

func Start() {
	s := "abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"
	res := Solve(s)
	fmt.Println(res)
}
