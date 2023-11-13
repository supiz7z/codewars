package the_feast_of_many_beasts

import "fmt"

func Feast(beast, dish string) bool {
	bFirst := string(rune(beast[0]))
	bLast := string(rune(beast[len(beast)-1]))
	dFirst := string(rune(dish[0]))
	dLast := string(rune(dish[len(dish)-1]))
	fmt.Println(bFirst, dFirst, bLast, dLast)
	return bFirst == dFirst && bLast == dLast
}

func Start() {
	beast := "qwerty"
	dish1 := "qjhjnkjkjkg__gg gkoky"
	dish2 := "wjjhgj  _ _  kkjfnf"
	total1 := Feast(beast, dish1)
	total2 := Feast(beast, dish2)
	fmt.Printf("ожидаем true: %t\nожидаем false: %t", total1, total2)
}
