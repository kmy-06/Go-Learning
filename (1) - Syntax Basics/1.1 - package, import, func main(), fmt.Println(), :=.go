package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	// int or float64 or bool or strings
	// var length, width int = 500, 1000 (Not The Best)
	// OR length, width= 500, 1000
	length, width := 500, 1000 // Answer already provides the type, no need to declare.
	// convert from int to float, temporarily
	// float64(length)
	// int(float64_value) -> this will drop the number behind the decimal point
	width2 := 1000

	fmt.Println(2+2 != 5)
	fmt.Println(math.Floor(2.45))
	fmt.Println(strings.Title("head first go"), float64(length)*float64(width), width2 > length)
	// same type in equation and comparison
	fmt.Println(reflect.TypeOf(3.14))
}
