package main

import (
	"fmt"
	"reflect"
)

func main() {

	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	fmt.Println(length > float64(quantity))
	fmt.Println("Total length across all sheets is", length * float64(quantity))
	fmt.Println(float64(quantity))
	fmt.Println(reflect.TypeOf(float64(quantity)))
	fmt.Println(int(length))
	fmt.Println(reflect.TypeOf(int(length)))

}
