package the_millionth_fibonacci_kata

import (
	"fmt"
	"math/big"
	"math/bits"
	"time"
)

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func CalA(a, b *big.Int, ch chan *big.Int) {
	ch <- Mul(a, Sub(Mul(b, big.NewInt(2)), a))
}

func CalB(a, b *big.Int, ch chan *big.Int) {
	ch <- Add(Mul(a, a), Mul(b, b))
}

func Find(m int64) *big.Int {
	n := uint(m)
	if n == 0 {
		return big.NewInt(0)
	}
	chA := make(chan *big.Int)
	chB := make(chan *big.Int)
	a, b := big.NewInt(0), big.NewInt(1)
	mask := bits.RotateLeft(1, bits.Len(n-1))
	for {
		if mask == 1 {
			if n&mask > 0 {
				a = Add(Mul(a, a), Mul(b, b))
			} else {
				a = Mul(a, Sub(Mul(b, big.NewInt(2)), a))
			}
			return a
		}
		go CalA(a, b, chA)
		go CalB(a, b, chB)

		if n&mask > 0 {
			a = <-chB
			b = Add(a, <-chA)
		} else {
			a = <-chA
			b = <-chB
		}
		mask = bits.RotateLeft(mask, -1)
	}
}

func PlusOrMinus(n int64) int64 {
	if n%2 == 0 {
		return -1
	}
	return 1
}

func Fib(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0).Mul(big.NewInt(PlusOrMinus(n)), Find(-n))
	}
	return Find(n)
}

func Start() {
	t := time.Now()
	//var n int64 = 500000
	var n int64 = 2000000
	result := Fib(n)
	fmt.Println(time.Now().Sub(t), result)
}
