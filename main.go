package main

import (
	"fmt"
	"lab1/bno"
)

func main() {
	var binNum [][]float64
	binNum = bno.ConvertToBin(4)

	fmt.Println(binNum)

	fmt.Println(bno.Summation(bno.ConvertToBin(5.75), bno.ConvertToBin(3.5)))
}
