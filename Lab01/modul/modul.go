package modul

import "fmt"

func Months(x int) int {

	return x * 12

}
func Days(x int) int {
	return x * 365
}

func YearsOnMars(x float32) float32 {

	z := 365
	m := 686
	s := float32(z) / float32(m)
	return x * s
}

func YearsOnWenus(x float32) {

	z := 365
	m := 225
	s := float32(z) / float32(m)
	fmt.Println(x * s)
}

func Delta(a int, b int, c int) {
	delta := (b * b) - 4*(a*c)
	fmt.Println(delta)
	if delta < 0 {
		fmt.Println("Trójmian nie ma pierwiastków")
	} else if delta > 0 {
		fmt.Println("Trójmian ma 2 pierwiastki")
	} else {
		fmt.Println("Trójmian ma 1 pierwiastek")
	}

}
