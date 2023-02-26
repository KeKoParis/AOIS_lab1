package bno // Package bno is "binary numbers operations"

import (
	"fmt"
	"math"
)

func ConvertToBin(number float64) {
	//fmt.Println(convertIntPart(number))
	//fmt.Println(convertFracPart(number))
	fmt.Println(formMantissa(math.Abs(number)))
	fmt.Println(formMantissa(number))
}

func formMantissa(number float64) ([]float64, float64) {
	var mantissa []float64
	var exp float64
	mantissa = append(convertIntPart(number), convertFracPart(number)...)

	if len(mantissa) == 71 && mantissa[0] == 1 {
		exp = 0
	}
	if len(mantissa) > 71 {
		exp = float64(len(mantissa) - 71)
	}
	if len(mantissa) == 71 && mantissa[0] == 0 {
		exp = getExp(mantissa)
	}

	if number < 0 {
		reverse(&mantissa)
		addOne(&mantissa)
	}

	exp = exp + 1024
	mantissa = mantissa[:53]
	return mantissa, exp
}

func getExp(mantissa []float64) float64 {
	for i := 1; i < len(mantissa); i++ {
		if mantissa[i] == 1 {
			return float64(i+1) * (-1)
		}
	}
	return 71
}

func convertIntPart(number float64) []float64 { // converts integer part of a number to binary
	number = math.Floor(math.Abs(number))
	var binNum, onesPosition, zero []float64

	if number == 0 {
		zero = append(zero, 0)
		return zero
	}

	for number >= 1 {
		onesPosition = append(onesPosition, math.Floor(math.Log2(number)))
		number = number - math.Exp2(math.Floor(math.Log2(number)))
	}

	for i := 0; i < int(onesPosition[0]+1); i++ {
		binNum = append(binNum, 0)
	}

	for i := len(onesPosition) - 1; i >= 0; i-- {
		binNum[len(binNum)-int(onesPosition[i])-1] = 1
	}

	return binNum
}

func convertFracPart(number float64) []float64 {
	number = math.Abs(number) - math.Floor(math.Abs(number))
	var binNum []float64

	for i := 0; i < 71; i++ {
		number = number * 2
		if number >= 1 {
			binNum = append(binNum, 1)
			number--
		} else {
			binNum = append(binNum, 0)
		}
	}

	return binNum
}

func reverse(mantissa *[]float64) {
	for i := 0; i < len(*mantissa); i++ {
		if (*mantissa)[i] == 0 {
			(*mantissa)[i] = 1
		} else {
			(*mantissa)[i] = 0
		}
	}
}

func addOne(mantissa *[]float64) {
	flagZero := 0

	for i := len(*mantissa) - 1; i >= 0; i-- {
		if (*mantissa)[i] == 0 {
			(*mantissa)[i] = 1
			flagZero = 1
			break
		}
	}

	if flagZero == 0 {
		var arr []float64
		arr = append(arr, 1)
		for i := 0; i < len(*mantissa); i++ {
			arr = append(arr, (*mantissa)[i])
		}
		*mantissa = arr
	}
}

func calcExp() {

}
