package main

import (
	"fmt"
)

func paintNeeded(width float64, length float64) {
	area := width * length
	fmt.Printf("%.2f liters of paint is needed.\n", area/10.0)
}

func main() {
	paintNeeded(5.2, 3.4)
	paintNeeded(3.4, 7.0)
	paintNeeded(6.7, 3.4)
}
