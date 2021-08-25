// pass_fail reports whether a grade is passing or failing

/* This is a multiline comment.
And it can be used for long ass comments.
Therefore, use sparringly.
Package widget includes all the functions used
for processing widgets.
*/

package main

import (
	"fmt"
	"github.com/lazycruise-sys/keyboard"
	"log"
)


func main() {

	// collecting data from keyboard
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// conditional statement to compare the grade variable to determine the pass/fail status
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)

	// fmt.Println(input)
	// fmt.Println(reflect.TypeOf(reader))

}
