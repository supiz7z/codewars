package some_functions

import (
	"errors"
	"math/big"
	"math/rand/v2"
	"strconv"
	"strings"
	"unicode"
)

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*(i-1) < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func IsPrimeOld(n int) bool {
	for i := 2; i*(i-1) < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func IsPrimeBig(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

func PrimeFactors(n int) []int {
	var arr []int
	for i := 2; n >= 2; {
		if IsPrime(n) {
			arr = append(arr, n)
			return arr
		}
		if n%i == 0 {
			n /= i
			arr = append(arr, i)
		} else {
			i += 1
		}
	}
	return arr
}

func DigitsCount(n int) int {
	if n/10 == 0 {
		return 1
	}
	return DigitsCount(n/10) + 1
}

func DigitsCountEven(n int) int {
	var total int
	if (n%10)%2 == 0 {
		total += 1
	}
	if n/10 == 0 {
		return total
	}
	return DigitsCountEven(n/10) + total
}

func DigitsCountOdd(n int) int {
	var total int
	if (n%10)%2 != 0 {
		total += 1
	}
	if n/10 == 0 {
		return total
	}
	return DigitsCountOdd(n/10) + total
}

func ReverseString(s string) string {
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}

func ReverseDigitsOfNumberByInt(n int) int {
	var result int
	for n > 0 {
		memory := n % 10
		result = result*10 + memory
		n /= 10
	}
	return result
}

func ReverseDigitsOfNumberByString(n int) int {
	s := strconv.Itoa(n)
	arr := []rune(s)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	n, _ = strconv.Atoi(string(arr))
	return n
}

func ReverseSlice(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func RemoveFromSliceByIndex(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func InsertIntoSliceByIndex(arr []int, elem, i int) []int { //функция говна, но можно юзать при запихивании не в конец
	var resultArr []int
	resultArr = append(resultArr, arr[:i]...)
	resultArr = append(resultArr, elem)
	resultArr = append(resultArr, arr[i:]...)
	return resultArr
}

func RandFill(minimum, maximum int) int {
	return rand.IntN(maximum-minimum+1) + minimum
}

func CutAfter(s, after string, length int) (string, error) {
	if i := strings.Index(s, after); i >= 0 {
		if len(s) < i+len(after)+length {
			return "", errors.New("too short string")
		}
		return s[i+len(after) : i+len(after)+length], nil
	}
	return "", errors.New("not found substring")
}

func CreateLogFuncName(s string) string {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if unicode.IsUpper(slice[i]) {
			if i == 0 {
				slice[i] = unicode.ToLower(slice[i])
				continue
			}
			slice[i] = unicode.ToLower(slice[i])
			var add []rune
			add = append(add, []rune("-")...)
			add = append(add, slice[i:]...)
			slice = append(slice[:i], add...)
		}
	}
	return string(slice)
}

func Gcd(n, d int) int {
	for n != 0 && d != 0 {
		if n > d {
			n = n % d
		} else {
			d = d % n
		}
	}
	return n + d
}

func GCD(a, b uint64) uint64 {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func Ctz(x uint64) uint64 {
	// This uses a binary search algorithm from Hacker's Delight.
	n := uint64(1)
	if (x & 0x0000FFFF) == 0 {
		n = n + 16
		x = x >> 16
	}

	if (x & 0x000000FF) == 0 {
		n = n + 8
		x = x >> 8
	}

	if (x & 0x0000000F) == 0 {
		n = n + 4
		x = x >> 4
	}

	if (x & 0x00000003) == 0 {
		n = n + 2
		x = x >> 2
	}

	return n - (x & 1)
}

func BinaryGCD(u, v uint64) uint64 {
	shift := uint64(0)

	if u == 0 {
		return v
	}

	if v == 0 {
		return u
	}

	shift = Ctz(u | v)
	u >>= Ctz(u)

	for {
		v >>= Ctz(v)
		if u > v {
			t := v
			v = u
			u = t
		}
		v = v - u

		if v == 0 {
			break
		}
	}

	return u << shift
}
