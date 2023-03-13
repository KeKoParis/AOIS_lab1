package sm

/*
	functions for 16 bit nums
*/

func Sum(num1 [size]float64, num2 [size]float64) []float64 {
	firstNumber := num1[:]
	secondNumber := num2[:]

	for i := 15; i >= 0; i-- {
		if secondNumber[i] == 1 {
			firstNumber = findZero(firstNumber, i)
		}
	}

	return firstNumber
}

func findZero(number []float64, position int) []float64 {

	for i := position; i >= 0; i-- {
		if number[i] == 0 {
			number[i] = 1
			break
		} else {
			number[i] = 0
		}
	}

	return number
}
