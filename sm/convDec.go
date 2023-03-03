package sm

import "math"

func ConvToDec(number []float64) float64 {
	decNum := 0.0
	sign := 0.0
	if number[0] == 1 {
		sign = -1
		reverse(&number)
		addOne(&number)
	} else {
		sign = 1
	}

	for i := 0; i < 16; i++ {
		decNum = decNum + math.Exp2(float64(15-i))*number[i]
	}

	if sign == -1 {
		reverse(&number)
		addOne(&number)
	}

	decNum = sign * decNum
	return decNum
}
