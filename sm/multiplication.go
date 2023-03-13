package sm

import (
	"math"
)

func Multiply(firstNum [size]float64, secondNum [size]float64) []float64 {
	var sum []float64
	multiplier := math.Abs(ConvToDec(secondNum[:]))
	var arr [size]float64

	for i := 0; i < int(multiplier)-1; i++ {
		if i == 0 {
			sum = Sum(firstNum, firstNum)
		} else {
			for j := 0; j < size; j++ {
				arr[j] = sum[j]
			}
			sum = Sum(arr, firstNum)
		}
	}

	return sum
}
