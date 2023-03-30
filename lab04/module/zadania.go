package module

func Zadanie1() [20]float64 {
	var vector1 [20]float64
	var vector2 [20]float64

	for i := range vector1 {
		vector1[i] = 2.0
		vector2[i] = 3.0
	}

	var sumOfvectors [20]float64

	for i := range vector2 {
		sumOfvectors[i] = vector1[i] + vector2[i]
	}

	return sumOfvectors
}

func Zadanie2() [5]float64 {
	vector1 := [5]float64{1, 2, 3, 4, 5}
	vector2 := [5]float64{5, 4, 3, 2, 1}

	var hadamardProduct [5]float64

	for i := range hadamardProduct {
		hadamardProduct[i] = vector1[i] * vector2[i]
	}

	return hadamardProduct
}
