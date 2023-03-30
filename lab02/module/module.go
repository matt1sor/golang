package module

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func NicknameChange(x string) []uint8 {

	nickname := []byte(x)
	return nickname

}

func Factorial(y int64) *big.Int {
	x := new(big.Int)
	return x.MulRange(1, y)

}
func Fibonacci(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	return n1
}

func Driver(x string) {

	ascii := NicknameChange(x)
	var s int64 = 10

sil:
	for s < 1000 {

		silnia := Factorial(s).String()

		for i := 0; i < len(ascii); i++ {

			str := ascii[i]

			z := strconv.FormatUint(uint64(str), 10)

			if strings.Contains(silnia, z) {

				if len(ascii)-1 == i {
					fmt.Println(s)
					fmt.Println("containsa in ")
					break sil
				}
			} else {

				break
			}

		}

		s = s + 1
	}

}

//func MatchNumbers(haystack string, needle string) bool {
//	//for i :=1; i<len(Factorial(x))
//
//}

//func DriverKMP(x string) {
//	ascii := NicknameChange(x).St
//
//	factorial := 0
//
//	fibonacci := 0
//
//	var sum int64 = 1
//
//	for sum < 1000 {
//
//		for number := 0; number < len(ascii); number++ {
//			 tab := ascii[number]
//			 needle := tab
//
//			fac := Factorial(sum).String()
//			if MatchNumbers(fac,needle)
//
//
//		}
//
//		sum += sum
//	}
//}
