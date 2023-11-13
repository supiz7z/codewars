package pell_numbers

import (
	"fmt"
	"math/big"
)

func Pell(n int) *big.Int {
	arr := []*big.Int{big.NewInt(0), big.NewInt(1)}
	for len(arr) <= n {
		arr = append(arr, big.NewInt(0).Add(big.NewInt(0).Mul(big.NewInt(2), arr[len(arr)-1]), arr[len(arr)-2]))
	}
	//fmt.Println(arr)
	return arr[len(arr)-1]
}

func Start() {
	var n = 40
	total := Pell(n)
	fmt.Println(total)
}
