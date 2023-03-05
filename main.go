package main

import (
	"flag"
	"fmt"
	"lab1/bno"
	"lab1/sm"
)

func main() {

	firstNumber := flag.Float64("firstNumber", 5.155, "user's first number")
	secondNumber := flag.Float64("secondNumber", 3.25, "user's second number")
	flag.Parse()

	flag.Parse()
	firstValue := *firstNumber
	secondValue := *secondNumber

	var binFloat [][]float64
	binFloat = bno.Summation(bno.ConvertToBin(firstValue), bno.ConvertToBin(secondValue))
	fmt.Println(binFloat, "  len mantissa ", len(binFloat[2]))

	fmt.Println(bno.ConvertToDec(bno.Summation(bno.ConvertToBin(firstValue), bno.ConvertToBin(secondValue))))

	var firstNum [16]float64
	var secondNum [16]float64

	arr := sm.Convert(firstValue)
	for i := 0; i < 16; i++ {
		firstNum[i] = arr[i]
	}

	arr = sm.Convert(secondValue)
	for i := 0; i < 16; i++ {
		secondNum[i] = arr[i]
	}

	binNum := sm.Sum(firstNum, secondNum)
	dec := sm.ConvToDec(binNum)
	fmt.Println("sum ", binNum, dec)

	binNum = sm.Div(firstNum, secondNum)
	dec = sm.ConvToDec(binNum)
	fmt.Println("div ", binNum, dec)

	binNum = sm.Multiply(firstNum, secondNum)
	dec = sm.ConvToDec(binNum)
	fmt.Println("mul ", binNum, dec)
}
