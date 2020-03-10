package main

import "fmt"

func main() {
	name := "alex"
	surname := "papa"
	hallo(name, surname)
	println(name, surname)
	halloPointer(&name, &surname)
	println(name, surname) // value changed cause we changed th apointer
	sum, pointerSum := variadicParamsSum(1, 2, 3, 4, 5)
	println(sum)
	println(pointerSum)
	res := plus(1, 2)
	println(res)
	invokeFunction()
	n := num{number: 2}
	println(n.Multiply())
	println(n.Add())
}

// both name and surname are strings
// the variables are copied by value
// if we passed slices or maps it would pass a reference
func hallo(name, surname string) {
	println("hallo " + name + " " + surname)
}

// the variables are copied by reference
func halloPointer(name, surname *string) {
	*name = "john"
	println("hallo " + *name + " " + *surname)
}

// variatic parameter or spread operator
func variadicParamsSum(values ...int) (int, *int) {

	result := 0
	for _, v := range values {
		result += v
	}
	return result, &result
}

// in other languages the internal function stuck gets destroyed when the function exits
// so if you return a pointer it return invalid data for that pointer
// in go the moment you return a pointer it get promoted and the value is passed outside to the heap memory

// named return values, we create beforehand the value and names
func plus(a, b int) (result int) {
	result = a + b
	return
}

// returning an error from the function
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can not divide by 0")
	}
	return a / b, nil
}

func invokeFunction() {
	for i := 0; i < 5; i++ {
		func(i int) {
			println("invoked")
		}(i) // pass the outer scope value to the function
	}

	// function as variable
	// var f func(int,int) (int ,error)
	f := func() {
		println("as a variable")
	}
	f()
}

//! methods
// methods are functions attached to a struct

type num struct {
	number int
}

// value receiver if you want to work without manipulating the original data but keep in mind it takes additional memory
func (n num) Multiply() int {
	return n.number * 2
}

// pointer receiver
func (n num) Add() int {
	return n.number + n.number
}
