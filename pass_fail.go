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

func getFloat() (float64, error) {
	// collecting data from keyboard
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// removing spaces and newline to enable conversion to other type
	input = strings.TrimSpace(input)

	// converting to input data to float64 using the strconv package which takes the number of bits
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	// limit the input to between 0 and 100
	if number > 100 {
		err = fmt.Errorf("the number %.2f is not a gradeable number", number)
		return 0, err
	}

	return number, nil
}

func main() {

	// collecting data from keyboard
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
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
