package bno // Package bno is "binary numbers operations"

import (
	"math"
)

// convertExpToDec function converts bin exponent to decimal
func convertExpToDec(number []float64) float64 {
	exp := 0.0

	for i := 0; i < len(number); i++ {
		exp += math.Exp2(float64(len(number)-i)-1) * number[i]
	}

	exp = exp - 1023
	return exp
}

// Summation function summarizes two float bin numbers.
func Summation(firstNum [][]float64, secondNum [][]float64) [][]float64 {
	form1, form2, exp := formForSum(firstNum, secondNum)
	var sum [][]float64
	for i := len(form2) - 1; i >= 0; i-- {
		if form2[i] == 1 {
			findZeroPos(&form1, i)
		}
	}
	sign := []float64{form1[0]}
	sum = append(sum, sign)

	if (form1[0] == 0 && form1[1] == 1) || (form1[0] == 1 && form1[1] == 0) {
		exp += 1024
		sum = append(sum, calcExp(exp))
		sum = append(sum, form1[2:54])
	} else {
		sum = append(sum, calcExp(exp+1023))
		sum = append(sum, form1[3:55])
	}

	return sum
}

// formForSum function makes convenient form of bin numbers
// for summation.
func formForSum(firstNum [][]float64, secondNum [][]float64) ([]float64, []float64, float64) {
	var form1, form2 []float64
	var exp1, exp2, exp float64

	exp1 = convertExpToDec(firstNum[1])
	exp2 = convertExpToDec(secondNum[1])
	form1 = makeBeginning(firstNum)
	form2 = makeBeginning(secondNum)

	if exp1 >= exp2 {
		for i := 0; i < int(exp1-exp2); i++ {
			form2 = append(form2, secondNum[0][0])
		}
		exp = exp1
	} else {
		for i := 0; i < int(exp1-exp2); i++ {
			form1 = append(form1, firstNum[0][0])
		}
		exp = exp2
	}

	makeFull(&form1)
	makeFull(&form2)
	appendMantissa(&form1, firstNum[2])
	appendMantissa(&form2, secondNum[2])
	return form1, form2, exp
}

// makeBeginning function fills the beginning with sign and zero or one.
func makeBeginning(number [][]float64) []float64 {
	var form []float64

	form = append(form, number[0][0], number[0][0])

	return form
}

// makeFull function adds missing the most significant bit.
func makeFull(number *[]float64) {
	if (*number)[0] == 0 {
		*number = append(*number, 1)
	} else {
		*number = append(*number, 0)
	}
}

// appendMantissa function adds mantissa to form number.
func appendMantissa(form *[]float64, mantissa []float64) {
	*form = append(*form, mantissa...)
	*form = (*form)[:55]
}

// findZeroPos function changes the first one zero to one from position.
// It changes all ones along the way.
func findZeroPos(number *[]float64, position int) {
	for i := position; i >= 0; i-- {
		if (*number)[i] == 0 {
			(*number)[i] = 1
			break
		}
		(*number)[i] = 0
	}
}
