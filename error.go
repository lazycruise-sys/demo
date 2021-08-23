package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	fmt.Println(reflect.TypeOf(err))
}
