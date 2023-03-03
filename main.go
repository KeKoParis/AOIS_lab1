package main

import (
	"fmt"
	"lab1/bno"
	"lab1/sm"
)

func main() {

	c := 5.155
	b := 3.25
	var binFloat [][]float64
	binFloat = bno.Summation(bno.ConvertToBin(c), bno.ConvertToBin(b))
	fmt.Println(binFloat, "  len mantissa ", len(binFloat[2]))

	fmt.Println(bno.ConvertToDec(bno.Summation(bno.ConvertToBin(c), bno.ConvertToBin(b))))

	c = 27
	b = 2
	var firstNum [16]float64
	var secondNum [16]float64

	arr := sm.Convert(c)
	for i := 0; i < 16; i++ {
		firstNum[i] = arr[i]
	}

	arr = sm.Convert(b)
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
