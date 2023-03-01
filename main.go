package main

import (
	"fmt"
	"lab1/bno"
)

func main() {

	fmt.Println(bno.Summation(bno.ConvertToBin(555.), bno.ConvertToBin(445)))
	fmt.Println(bno.ConvertToDec(bno.Summation(bno.ConvertToBin(-5.), bno.ConvertToBin(-6))))
}
