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
	"os"
	"log"
	"reflect"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
	fmt.Println(reflect.TypeOf(reader))
}