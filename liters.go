package main

import (
	"fmt"
)
func sayHi()  {
	fmt.Println("hi!")
}
func main() {
	var width, length, area float64
	width = 3.45
	length = 3.22
	area = width * length
	fmt.Printf("%.2f liters of paint is needed.\n", area/10.0)
	width = 5.22
	length = 3.5
	area = width * length
	fmt.Printf("%.2f liters of paint is needed.\n", area/10.0)
	sayHi()
}
