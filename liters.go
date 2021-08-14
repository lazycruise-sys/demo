package main

import (
	"fmt"
)

func main() {
	var width, length, area float64
	width = 3.45
	length = 3.22
	area = width * length
	fmt.Println(area/10.00, "liters of paint is needed.")
}