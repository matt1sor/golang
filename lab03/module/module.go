package module

import "fmt"

func Collatza(x int) int {

	if x%2 == 0 {
		x = x / 2
		//fmt.Println(x)
		return x

	} else if x%2 == 1 {
		x = (x * 3) + 1
		//fmt.Println(x)
		return x

	}
	return x
}

func CollatzaDriver(i int) []int {

	var tablica = make([]int, 0)

	for i != 1 {
		i = Collatza(i)
		tablica = append(tablica, i)
	}
	return tablica
}

func CollatzRange(x int, y int) {

	length := 0
	liczba := 0 //??????

	for i := x; i <= y; i++ {

		z := CollatzaDriver(i)

		if len(z) > length {
			length = len(z)
			liczba = i

		}

	}

	fmt.Println("Długość:", length, "Liczba:", liczba)

}
