package main

import "fmt"

// for loop

func main() {

	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}

	fmt.Println("------------------")

	y := 10
	for y > 2 {
		fmt.Println(y)
		y -= 2
	}
	fmt.Println("happy new year, modafucka")
}
