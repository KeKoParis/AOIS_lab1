package bno

import (
	"fmt"
	"math"
)

func ConvertToDec(number [][]float64) float64 {
	var intPart, fracPart []float64
	var exp, decNum float64

	exp = convertExpToDec(number[1])

	if number[0][0] == 1 {
		reverse(&number[2])
		addOne(&number[2])
	}

	intPart = append(intPart, 1)

	intPart = append(intPart, number[2][:int(exp)]...)

	if exp < 0 {
		fracPart = append(fracPart, 1)
	}
	fracPart = append(fracPart, number[2][int(exp):]...)
	fmt.Println(intPart)
	decNum = 0
	for i := 0; i < len(intPart); i++ {
		decNum += intPart[i] * math.Exp2(float64(len(intPart)-i-1))
	}
	for i := 0; i < len(fracPart); i++ {
		decNum += fracPart[i] * math.Exp2(float64(-1*i-1))
	}

	if number[0][0] == 1 {
		decNum = decNum * (-1)
	}

	return decNum
}
