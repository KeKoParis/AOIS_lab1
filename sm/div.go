package sm

const size = 16

func Div(num1 [size]float64, num2 [size]float64) []float64 {
	var firstNum, secondNum []float64
	firstNum = num1[:]
	secondNum = num2[:]

	division := 0.0

	sign1 := 1
	sign2 := 1

	if firstNum[0] == 1 {
		reverse(&firstNum)
		addOne(&firstNum)
		sign1 = -1
	}

	if secondNum[0] == 1 {
		reverse(&secondNum)
		addOne(&secondNum)
		sign2 = -1
	}
	reverse(&secondNum)
	addOne(&secondNum)

	var divider [size]float64

	for i := 0; i < size; i++ {
		divider[i] = secondNum[i]
	}

	for true {
		firstNum = Sum(num1, divider)
		for i := 0; i < size; i++ {
			num1[i] = firstNum[i]
		}
		if firstNum[0] == 1 {
			break
		}
		division++
	}

	division = division * float64(sign1*sign2)
	return Convert(division)
}
