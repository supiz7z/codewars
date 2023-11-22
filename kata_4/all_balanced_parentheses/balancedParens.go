package all_balanced_parentheses

import "fmt"

func AppendElements(str string, n, leftCount, rightCount int) []string {
	var arr []string
	if leftCount == rightCount && leftCount == n {
		return append(arr, str)
	}
	if leftCount < n {
		arr = append(arr, AppendElements(str+"(", n, leftCount+1, rightCount)...)
	}
	if leftCount > rightCount {
		arr = append(arr, AppendElements(str+")", n, leftCount, rightCount+1)...)
	}
	return arr
}

func BalancedParens(n int) []string {
	var arr []string
	var str string
	var leftCount, rightCount int
	arr = AppendElements(str, n, leftCount, rightCount)
	return arr
}

func Start() {
	n := 3
	result := BalancedParens(n)
	fmt.Println(result)
}
