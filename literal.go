package main

import (
	"fmt"
	"time"
)

func main() {
	// initializing array
	var notes [7] string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes)


	var dates [3] time.Time
	dates[0] = time.Unix(3289239010, 0)
	fmt.Println(dates[0])

	var primes [5] int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	fmt.Println(primes)
	primes[1]++
	fmt.Println(primes)

	names := [3] string {"akinlua", "olorunfemi", "praise"}
	fmt.Println(names)
}