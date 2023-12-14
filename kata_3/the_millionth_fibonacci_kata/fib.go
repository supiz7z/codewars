package the_millionth_fibonacci_kata

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func Find(n int64) *big.Int {
	m := big.NewInt(n + 1)

	f1 := big.NewInt(0)
	f2 := big.NewInt(1)

	if m.Cmp(big.NewInt(1)) == 0 {
		return f1
	}
	if m.Cmp(big.NewInt(2)) == 0 {
		return f2
	}
	for i := 3; m.Cmp(big.NewInt(int64(i))) >= 0; i++ {
		next := big.NewInt(0)
		next.Add(f1, f2)
		f1 = f2
		f2 = next
	}
	return f2
}

func PlusOrMinus(n int64) int64 {
	if n%2 == 0 {
		return -1
	}
	return 1
}

func Fib(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0).Mul(big.NewInt(PlusOrMinus(n)), Find(int64(math.Abs(float64(n)))))
	}
	return Find(n)
}

func Start() {
	t := time.Now()
	var n int64 = 500000
	//var n int64 = 50
	result := Fib(n)
	fmt.Println(time.Now().Sub(t), result)
}
