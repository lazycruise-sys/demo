package main

import (
	"fmt"
	"log"
)

// var metersPerLiter float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	// var amount, total float64
	// amount = paintNeeded(3.4, 6.7)
	// fmt.Printf("%.2f liters of paint is needed\n", amount)
	// total += amount
	// amount = paintNeeded(5.6, 9.2)
	// fmt.Printf("%.2f liters of paint is needed\n", amount)
	// total += amount
	// fmt.Printf("Total: %.2f is needed\n", total)

	amount, err := paintNeeded(-3.3, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters of paint is needed\n", amount)
}
