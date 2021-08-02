package main

import (
	"fmt"
	"strings"
	"sort"
)


func main() {

	greeting := "Hello, Welcome to Cruise"

	fmt.Println(strings.Contains(greeting, "hello")) // case-sensitive
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) // the ReplaceAll() function doesn't change the original string
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "W"))
	fmt.Println(strings.Split(greeting, " "))

	// Creation of Slice from String
	string_slice := strings.Split(greeting, " ")

	new_string := strings.ReplaceAll(greeting, "Hello", "Hi") // assigning the changed outcome to a new string
	fmt.Println(new_string)

	// To test for changes
	fmt.Println(greeting)
	fmt.Println(string_slice, string_slice[1])

	ages := []int{23, 45, 67, 89, 23, 45, 31, 33}
	names := []string{"Ye", "Yonda", "Burna", "Femi", "Yamela", "Catherine"}
	
	sort.Ints(ages) // This methods alters the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 23)
	fmt.Println(index)

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Ye"))

}