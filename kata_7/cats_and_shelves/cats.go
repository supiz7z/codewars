package cats_and_shelves

import "fmt"

func Cats(start, finish int) int {
	var count1, count3, total int
	r := finish - start
	count3 = r / 3
	count1 = r - count3*3
	total = count1 + count3
	return total
}

func Start() {
	start := 12
	finish := 46
	total := Cats(start, finish)
	fmt.Println(total)
}
