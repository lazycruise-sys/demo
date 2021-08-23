package main

import (
	"fmt"
)

// var metersPerLiter float64

func paintNeeded(width float64, length float64) float64 {
	area := width * length
	return area / 10.0
}

func main() {
	var amount, total float64
	amount = paintNeeded(3.4, 6.7)
	fmt.Printf("%.2f liters of paint is needed\n", amount)
	total += amount
	amount = paintNeeded(5.6, 9.2)
	fmt.Printf("%.2f liters of paint is needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f is needed\n", total)
}
