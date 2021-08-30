package main

import (
	"fmt"
	"time"
	"reflect"
)

func main() {

	// declaring and initializing array
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes)

	var dates [3]time.Time
	dates[0] = time.Unix(3289239010, 0)
	fmt.Println(dates[0])

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	fmt.Println(primes)
	primes[1]++
	fmt.Println(primes)

	names := [3]string{"akinlua", "olorunfemi", "praise"}
	fmt.Println(names)

	text := [3]string{"well, this is going to be long",
		"how about we put it across multiple lines",
		"however, make sure you use the comma even at the end, weird",
	}

	fmt.Println(text)
	fmt.Printf("%#v\n", text)
	fmt.Println(len(dates))

	// typical for loop structure involving arrays
	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	// the safer for..range structure involving arrays
	for index, note := range notes {
		fmt.Println(index, note) // "index" holds the element index, and "note" holds the array element itself and "notes" is the array.
	}

	// one thing, you have to use all variables declared in go which means the "index" variable has to be used
	// even when it is not needed. to avoid this, the use of the blank identifier is useful (as with the error usage)
	// this can also be done for the "note" variable
	for _, note := range notes {
		fmt.Println(note)
	}

	for index := range notes {
		fmt.Println(index)
	}

	beefs := [3] float64 {71.6, 73.4, 76.5}
	var numOfBeefs float64
	for _, beef := range beefs {
		numOfBeefs += beef
	}
	fmt.Println(numOfBeefs)

	fmt.Println(reflect.TypeOf(primes))
	var numbers [3]float64
	// i := 0
	fmt.Println(numbers)
	fmt.Println(reflect.TypeOf(numbers))
}
