//Projekt na dodatkowe punkty 10.06.2023

package main

import "las/las"
import (
	"fmt"
	"os"
)

func showOptions() {
	fmt.Println("Opcje programu:")
	fmt.Println("1 - Pojedyncza symulacja 10x10 z wyborem ilości drzew")
	fmt.Println("2 - Symulacja 10x10 z podaną ilościa drzew i iteracji na symulacje")
	fmt.Println("3 - Wszystkie dowolne parametry")
	fmt.Println("4 - Generowanie wykresu z pliku data.csv")
	fmt.Println("5 - Wyjście")
}
func option1() {

	clearData()
	width := 10
	height := 10
	var treeCount int

	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&treeCount)
	las.BurnTRREES(width, height, treeCount)
}
func option2() {

	clearData()

	width := 10
	height := 10

	var treeCount int
	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&treeCount)

	var simCount int
	fmt.Print("Podaj ilość iteracji na symulacje: ")
	fmt.Scan(&simCount)

	for counter := treeCount; counter < 100; counter += 5 {
		for i := 0; i < simCount; i++ {
			las.BurnTRREES(width, height, counter)
		}
	}

}
func option3() {

	clearData()

	var width int
	fmt.Print("Podaj szerokość lasu: ")
	fmt.Scan(&width)

	var height int
	fmt.Print("Podaj wysokość lasu: ")
	fmt.Scan(&height)

	var treeCount int
	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&treeCount)

	var simCount int
	fmt.Print("Podaj ilość iteracji na symulacje: ")
	fmt.Scan(&simCount)

	var addedTrees int
	fmt.Print("Podaj o jaka ilosc drzew maja róznić sie kolejnych symulacji: ")
	fmt.Scan(&addedTrees)

	for counter := treeCount; counter < (height * width); counter += addedTrees {
		for i := 0; i < simCount; i++ {
			las.BurnTRREES(width, height, counter)
		}
	}

}
func clearData() {
	filePath := "data.csv"

	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
}

func main() {
	running := true
	for running {
		showOptions()
		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			option1()
		case 2:
			option2()
		case 3:
			option3()
		case 4:
			las.PlotData()
		case 5:
			running = false
			os.Exit(0)
		default:
			fmt.Println("Podano zła wartość :(")
		}
	}
}
