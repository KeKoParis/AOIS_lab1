package main

import (
	"fmt"
	"lab1/bno"
)

func main() {
	var binNum [][]float64
	binNum = bno.ConvertToBin(5.25)

	fmt.Println(binNum)
}
