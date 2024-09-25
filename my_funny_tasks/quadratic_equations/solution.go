package quadratic_equations

import (
	"fmt"
	"math"
)

type Disc struct {
	Data string
}

func (d *Disc) Error() string {
	return d.Data
}

func D(a, b, c float64) (float64, float64, error) {
	d := math.Pow(b, 2) - 4*a*c
	fmt.Printf("считаем дискриминант по формуле D = b^2 - 4ac:\nD = (%v)*(%v) - 4 * (%v) * (%v) = (%v) - (%v) = %v\nквадратный корень из дискриминанта равен %v\n", b, b, a, c, math.Pow(b, 2), 4*a*c, d, math.Sqrt(d))
	switch {
	case d < 0:
		return 0, 0, &Disc{Data: "дискриминант меньше нуля, нет действительных корней"}
	case d == 0:
		x1 := -b / (2 * a)
		return x1, x1, nil
	default:
		x1 := (-b - math.Sqrt(d)) / (2 * a)
		x2 := (-b + math.Sqrt(d)) / (2 * a)
		return x1, x2, nil
	}
}

func GetStringEquation(a, b, c float64) string {
	checkA := func(i float64) string {
		switch {
		case i == 1:
			return fmt.Sprintf("x^2")
		case i == -1:
			return fmt.Sprintf("-x^2")
		default:
			return fmt.Sprintf("%vx^2", a)
		}
	}

	checkB := func(i float64) string {
		switch {
		case i < 0:
			return fmt.Sprintf("- %vx", -b)
		case i == 0:
			return ""
		default:
			return fmt.Sprintf("+ %vx", b)
		}
	}

	checkC := func(i float64) string {
		switch {
		case i < 0:
			return fmt.Sprintf("- %v", -c)
		case i == 0:
			return ""
		default:
			return fmt.Sprintf("+ %v", c)
		}
	}
	return fmt.Sprintf("%s %s %s = 0", checkA(a), checkB(b), checkC(c))
}

func Start() {
	var a, b, c float64
	a = 3
	b = 0
	c = -9

	eq := GetStringEquation(a, b, c)

	if a == 0 {
		fmt.Println("если коэффициент \"a\" равен 0, то уравнение не квадратное, а линейное")
		return
	}

	fmt.Printf("решим уравнение %s:\na = %v, b = %v, c = %v\n", eq, a, b, c)

	x1, x2, err := D(a, b, c)
	if err != nil {
		fmt.Println(err)
		return
	}

	if x1 == x2 {
		fmt.Printf("уравнение %s имеет один корень (т.к. дискриминант равен нулю)\n%s", eq,
			fmt.Sprintf("считаем корень по формуле x = (-b)/(2a):\nx = (%v)/(2*(%v)) = %v\n", -b, a, x1))
		return
	}

	fmt.Printf("уравнение %s имеет два корня (т.к. дискриминант больше нуля)\n%s", eq,
		fmt.Sprintf("считаем корни по формулам:\n"+
			fmt.Sprintf("x1 = (-b - D^0,5)/(2a) = (-(%v) - %v)/(2*(%v)) = (%v)/(%v) = %v\n", b, math.Sqrt(math.Pow(b, 2)-4*a*c), a, -b-math.Sqrt(math.Pow(b, 2)-4*a*c), 2*a, x1)+
			fmt.Sprintf("x2 = (-b + D^0,5)/(2a) = (-(%v) + %v)/(2*(%v)) = (%v)/(%v) = %v\n", b, math.Sqrt(math.Pow(b, 2)-4*a*c), a, -b+math.Sqrt(math.Pow(b, 2)-4*a*c), 2*a, x2)))
}
