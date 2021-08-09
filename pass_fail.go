// pass_fail reports whether a grade is passing or failing

/* This is a multiline comment.
And it can be used for long ass comments.
Therefore, use sparringly.
Package widget includes all the functions used
for processing widgets.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "reflect"
	"strconv"
	"strings"
)

func main() {

	// collecting data from keyboard
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// removing spaces and newline to enable conversion to other type
	input = strings.TrimSpace(input)

	// converting to input data to float64 using the strconv package which takes the number of bits
	grade, err := strconv.ParseFloat(input, 64)
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
