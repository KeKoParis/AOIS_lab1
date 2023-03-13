package sm

import (
	"math"
)

/*
	functions for 16 bit nums
*/

func Convert(number float64) []float64 {
	var binNumber []float64
	if math.Floor(math.Abs(number)) != 0 {
		binNumber = getBin(number)
	} else {
		binNumber = getZero()
	}

	if number < 0 {
		reverse(&binNumber)
		addOne(&binNumber)
	}

	return binNumber
}

func getBin(number float64) []float64 {
	var onesPlace, binNumber []float64
	number = math.Abs(number)

	for i := 0; i < 16; i++ {
		binNumber = append(binNumber, 0)
	}

	for number >= 1 {
		onesPlace = append(onesPlace, math.Floor(math.Log2(number)))
		number = number - math.Exp2(math.Floor(math.Log2(number)))
	}

	for i := 0; i < len(onesPlace); i++ {
		binNumber[size-int(onesPlace[i])-1] = 1
	}

	return binNumber
}

func getZero() []float64 {
	var binNumber []float64

	for i := 0; i < size; i++ {
		binNumber = append(binNumber, 0)
	}

	return binNumber
}

func reverse(binNumber *[]float64) {
	for i := 0; i < len(*binNumber); i++ {
		if (*binNumber)[i] == 0 {
			(*binNumber)[i] = 1
		} else {
			(*binNumber)[i] = 0
		}
	}
}

func addOne(binNumber *[]float64) {
	for i := len(*binNumber) - 1; i >= 0; i-- {
		if (*binNumber)[i] == 0 {
			(*binNumber)[i] = 1
			break
		} else {
			(*binNumber)[i] = 0
		}
	}
}
