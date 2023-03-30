package main

import (
	"Lab01/modul"
	"fmt"
)

func main() {

	fmt.Println("You have:", modul.Months(22), "months")
	fmt.Println("You have:", modul.Days(22), "days")
	fmt.Println("Your age on Mars:", modul.YearsOnMars(22))
	modul.YearsOnWenus(22)
	modul.Delta(2, 9, 5)

	fmt.Println(6 / 2 * (1 + 2))
	fmt.Println(9 - 3/(1/3) + 1)
}
