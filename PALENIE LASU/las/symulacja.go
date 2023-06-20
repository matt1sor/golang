package las

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func fireSpread(las [][]int, coordX int, coordY int) {
	if coordY < 0 || coordY >= len(las) ||
		coordX < 0 || coordX >= len(las[coordY]) {
		return
	}

	// prawo
	if coordX < len(las[coordY])-1 && las[coordY][coordX+1] == 1 {
		las[coordY][coordX+1] = 2
		fireSpread(las, coordX+1, coordY)
	}

	//lewo
	if coordX > 0 && las[coordY][coordX-1] == 1 {
		las[coordY][coordX-1] = 2
		fireSpread(las, coordX-1, coordY)
	}

	// góra
	if coordY > 0 && las[coordY-1][coordX] == 1 {
		las[coordY-1][coordX] = 2
		fireSpread(las, coordX, coordY-1)
	}

	// dół
	if coordY < len(las)-1 && las[coordY+1][coordX] == 1 {
		las[coordY+1][coordX] = 2
		fireSpread(las, coordX, coordY+1)
	}

	// SKOSY

	// prawo-góra
	if coordX < len(las[coordY])-1 && coordY > 0 && las[coordY-1][coordX+1] == 1 {
		las[coordY-1][coordX+1] = 2
		fireSpread(las, coordX+1, coordY-1)
	}
	//lewo-góra
	if coordY > 0 && coordX > 0 && las[coordY-1][coordX-1] == 1 {
		las[coordY-1][coordX-1] = 2
		fireSpread(las, coordX-1, coordY-1)
	}
	//lewo-dół
	if coordX > 0 && coordY < len(las)-1 && las[coordY+1][coordX-1] == 1 {
		las[coordY+1][coordX-1] = 2
		fireSpread(las, coordX-1, coordY+1)
	}
	//prawo-dół
	if coordX < len(las[coordY])-1 && coordY < len(las)-1 && las[coordY+1][coordX+1] == 1 {
		las[coordY+1][coordX+1] = 2
		fireSpread(las, coordX+1, coordY+1)
	}

}

func thunder(las [][]int) {
	coordX := rand.Intn(len(las[0]))
	coordY := rand.Intn(len(las))

	//fmt.Println("Pozycja strzału pioruna", coordX+1, coordY+1)

	if las[coordY][coordX] == 1 {
		las[coordY][coordX] = 2
		fmt.Println("Piorun trafił i podalił drzewo! UCIEKAJ!!!")
		fireSpread(las, coordX, coordY)
	} else {
		fmt.Println("Piorun nie trafił w drzewo! Ufff!")
	}

}

func countTrees(las [][]int) int {
	trees := 0
	for i := 0; i < len(las[0]); i++ {
		for j := 0; j < len(las); j++ {
			if las[j][i] == 1 {
				trees++
			}
		}
	}
	return trees
}
func printForest(las [][]int) {
	for i := 0; i < len(las); i++ {
		for j := 0; j < len(las[i]); j++ {
			switch las[i][j] {
			case 0:
				fmt.Print("\033[0;37m0\033[0m")
			case 1:
				fmt.Print("\033[0;32m1\033[0m")
			case 2:
				fmt.Print("\033[0;31m2\033[0m")
			}
		}
		fmt.Println()
	}
}

func saveDataToFile(size string, startTrees, leftTrees, burnedTrees int) {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Nie można otworzyć pliku CSV:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{
		size,
		strconv.Itoa(startTrees),
		strconv.Itoa(leftTrees),
		strconv.Itoa(burnedTrees),
	}

	err = writer.Write(row)
	if err != nil {
		log.Fatal("Błąd podczas zapisu do pliku CSV:", err)
	}

	writer.Flush()
}

func plantTrees(trees int, las [][]int) {
	planted := 0
	for planted < trees {
		for i := 0; i < len(las[0]); i++ {
			for j := 0; j < len(las); j++ {
				if planted == trees {
					break
				}

				x := rand.Intn(2)
				if x == 0 && las[j][i] == 0 {
					las[j][i] = 1
					planted++
				}
			}
		}

	}
}
func generateArray(width, height int) [][]int {
	las := make([][]int, height)
	for i := 0; i < height; i++ {
		las[i] = make([]int, width)
	}
	return las
}
func BurnTRREES(width, height, trees int) {
	if width*height <= trees {
		fmt.Println("ERROR!, ilość drzew przekracza maksymalną ilość: ", width*height)
		return
	}

	las := generateArray(width, height)

	plantTrees(trees, las)
	printForest(las)
	thunder(las)
	printForest(las)
	treesAfterThunder := countTrees(las)
	burnedTrees := trees - treesAfterThunder
	fmt.Println("Ilość drzew:", treesAfterThunder)
	fmt.Println("Spalone drzewa:", burnedTrees)
	size := fmt.Sprintf("%dx%d", width, height)
	saveDataToFile(size, trees, treesAfterThunder, burnedTrees)
}
