package main

import (
	"fmt"
	"strconv"
)

// declare multiple variabless this case its a package level variable
var (
	a = "a"
	b = "b"
	c = "c"
)

// if a variable is lowercase its only exposed to this package if it is uppercase it is exposed to the other packages
var R = 42

// shadowing example value that can be re-declated and reused later <line 16>
var i int = 2

func main() {
	// variables
	var i int
	i = 42
	i = 27
	println(i)

	k := 4.9
	k = 27
	fmt.Printf("%v, %T \n", k, k)
	// conversion

	i = int(k)
	fmt.Printf("%v, %T \n", i, i)

	// when converting to string it reads the int as an byte so it will try to convert it to a character
	s := string(i)
	fmt.Printf("%v, %T \n", s, s)

	// correct way of conversion to string
	s = strconv.Itoa(i)
	fmt.Printf("%v, %T \n", s, s)

	// variable names
	// Pascal => Pascal
	// Camel Case => camelCase
	// Acronyms => newURL
}
