package main

import "fmt"

func main() {
	copy()
	pointer()
	pointerList()
	pointerToObject()
}

func copy() {
	println("======> copy")
	a := 43
	b := a
	println(a, b)
	a = 27
	println(a, b)
}

func pointer() {

	println("======> pointer")

	var a int = 43
	// creating a pointer
	var b *int = &a
	println(&a)
	println(a, b)
	a = 27
	println(a, b)
	// dereferencing
	println(a, *b)
}

func pointerList() {

	println("======> pointerList")

	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]

	fmt.Printf("a %v ", a)
	fmt.Printf("b %p ", b)
	fmt.Printf("c %p ", c)
	println()
}

type nm struct {
	number int
}

func pointerToObject() {
	println("======> pointerToObject")
	var ms *nm = &nm{number: 2}
	println(ms)

	var msn *nm
	msn = new(nm)
	println(msn)

	fmt.Printf("%p \n", (*msn))
	// here even if we have a pointer the compiler knows that we want the value so it dereferences the pointer for us
	fmt.Println(msn.number)
}
