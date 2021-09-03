package main

import (
	"fmt"
	"reflect"
)

func anoda() {
	fmt.Println("-------another test---------")
}

func main() {

	// long declarative statement
	var notes []string
	notes = make([]string, 7)
	notes[1] = "re"
	fmt.Println(notes)
	fmt.Println(reflect.TypeOf(notes))

	// // short declarative statement
	// primes := make([]int, 5)
	// primes[0] = 2
	// fmt.Println(len(primes))

	// numbers := []int{3, 4, 5}

	// uni := []string{
	// 	"OAU",
	// 	"UI",
	// 	"CU",
	// 	"FUTA",
	// }

	// fmt.Println(numbers, uni)

	// uArray := [5]string{"a", "b", "c", "d", "e"}
	// newSlice := uArray[1:3]
	// fmt.Println(reflect.TypeOf(newSlice))
	// fmt.Println(newSlice)
	// uArray[2] = "g"
	// fmt.Println(newSlice)
	// sSlice := uArray[:3]
	// eSlice := uArray[2:]
	// fmt.Println(sSlice, eSlice)

	// unto a new slice, we shall append
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))

	anoda()
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append(s1, 3, 5)
	s3 := append(s1, 6, 7, 8)
	fmt.Println(s1, s2, s3)
	s3[0] = 0
	fmt.Println(s1, s2, s3)

	anoda()
	// another underlying array test
	c1 := []string{"s1", "s2"}
	c2 := append(c1, "s2", "s2")
	c3 := append(c2, "s3", "s3")
	c4 := append(c3, "s4", "s4")
	fmt.Println(c1, c2, c3, c4)
	c4[0] = "XX"
	fmt.Println(c1, c2, c3, c4)

	anoda()
	var intSlice []int
	var boolSlice []bool
	fmt.Printf("intSlice: %#v\nboolSlice: %#v\n", intSlice, boolSlice)
	fmt.Println(len(intSlice), len(boolSlice))

	anoda()
	var emptySlice []string
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, "first content")
	}
	fmt.Println(emptySlice)

	anoda()
	// var wow []float64
	wow := make([]float64, 8)
	fmt.Printf("wow: %#v\n", wow)
	wow = append(wow, 1)
	fmt.Println(wow)
}
