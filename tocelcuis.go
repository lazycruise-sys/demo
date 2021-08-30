package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a Fahrenheit Temperature: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcuis := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees celcuis", celcuis)
}
