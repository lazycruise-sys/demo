package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("sleepy creek county population:", population)
	fmt.Println("congratulations, kevin and anna! it's a girl!")
	population += 1
	fmt.Println("sleepy creek county population:", population)
}