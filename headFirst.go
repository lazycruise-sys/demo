package main

import (
	"fmt"
	// "math"
	// "strings"
	// "reflect"
	"time"
)

func main() {

	// fullScore := math.Floor(32.4)
	// title := strings.Title("score card - Akinlua, Olorunfemi")

	// fmt.Println(reflect.TypeOf(32.4))
	// fmt.Println(reflect.TypeOf("head first go"))	
	// fmt.Println(reflect.TypeOf(23))
	// fmt.Println(reflect.TypeOf(false))

	// fmt.Println(fullScore)
	// fmt.Println(title)

	var now time.Time = time.Now()
	var year int = now.Year()
	var month time.Month = now.Month()
	var day int = now.Day()
	var hour int = now.Hour()
	var minute int = now.Minute()
	var second int = now.Second()

	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(second)

	fmt.Println(time.Now().Year())

}