package main

import (
	"fmt"
)

var metersPerLiter float64

func paintNeeded(width , length float64) float64 {
	area := width * length
	return area/metersPerLiter
}

func main() {
	metersPerLiter = 10.0
	fmt.Printf("%.2f", paintNeeded(5.2, 3.4))
	paintNeeded(3.4, 7.0)
	paintNeeded(6.7, 3.4)
}