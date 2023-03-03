package bno // Package bno is "binary numbers sm"

import (
	"fmt"
	"math"
)

// ConvertToBin function converts decimal numbers to binary.
// Returns floating format number (IEEE-754, 64bit).
func ConvertToBin(number float64) [][]float64 {
	var binNum [][]float64

	binMantissa, exp := formMantissa(number)
	a := (math.Log2(math.Abs(number))) - math.Abs(math.Floor(math.Log2(math.Abs(number))))
	if a == 0 {
		reverse(&binMantissa)
		binMantissa = binMantissa[1:]
		fmt.Println(binMantissa)
		exp--
	}

	binNum = append(binNum, getSign(number), calcExp(exp), binMantissa)

	return binNum
}

// formMantissa function forms mantissa from integer
// and fractional parts of a decimal number.
func formMantissa(number float64) ([]float64, float64) {
	var mantissa []float64
	var exp float64
	mantissa = append(convertIntPart(number), convertFracPart(number)...)

	if len(mantissa) == 72 && mantissa[0] == 1 {
		exp = 0
		mantissa = mantissa[1:53]
	}
	if len(mantissa) > 72 {
		exp = float64(len(mantissa) - 72)
		mantissa = mantissa[1:53]
	}

	if len(mantissa) == 72 && mantissa[0] == 0 {
		exp = getExp(mantissa)
		mantissa = mantissa[int(math.Abs(exp)) : int(math.Abs(exp))+53]
	}

	if number < 0 {
		reverse(&mantissa)
		addOne(&mantissa)
	}

	exp = exp + 1023
	return mantissa, exp
}

// getExp function finds exponent for a decimal number, which
// absolute value is less than one.
func getExp(mantissa []float64) float64 {
	for i := 1; i < len(mantissa); i++ {
		if mantissa[i] == 1 {
			return float64(i+1) * (-1)
		}
	}
	return 71
}

// convertIntPart function converts integer part of a number to binary.
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

// convertFracPart function converts fractional part of number
// to binary.
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

// reverse function reverses negative binary numbers according
// to two's complement.
func reverse(mantissa *[]float64) {
	for i := 0; i < len(*mantissa); i++ {
		if (*mantissa)[i] == 0 {
			(*mantissa)[i] = 1
		} else {
			(*mantissa)[i] = 0
		}
	}
}

// addOne function adds one to reversed binary number
// according to two's complement.
func addOne(mantissa *[]float64) {
	flagZero := 0

	for i := len(*mantissa) - 1; i >= 0; i-- {
		if (*mantissa)[i] == 0 {
			(*mantissa)[i] = 1
			flagZero = 1
			break
		}
		(*mantissa)[i] = 0
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

// calcExp function finds exponent for numbers, which
// absolute value is less than one
func calcExp(number float64) []float64 {
	var rawExp, exp []float64

	for i := 0; i < 11; i++ {
		exp = append(exp, 0)
	}

	rawExp = convertIntPart(number)

	for i := len(rawExp) - 1; i >= 0; i-- {
		exp[11-len(rawExp)+i] = rawExp[i]
	}

	return exp
}

// getSign function finds binary number sign
func getSign(number float64) []float64 {
	var numSign []float64

	if number >= 0 {
		numSign = append(numSign, 0)

	} else {
		numSign = append(numSign, 1)
	}

	return numSign
}
