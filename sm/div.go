package sm

func Div(num1 [16]float64, num2 [16]float64) []float64 {
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

	var divider [16]float64

	for i := 0; i < 16; i++ {
		divider[i] = secondNum[i]
	}

	for true {
		firstNum = Sum(num1, divider)
		for i := 0; i < 16; i++ {
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
