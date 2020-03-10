package main

import "fmt"

func main() {
	// console writer implements writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("hallo"))

}

// basics
func hallo(name, surname string) {
	println("hallo " + name + " " + surname)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// composing interfaces together

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

// type conversion of a interface to another
// bwc:= writerCloser.(* bufferWriterCloser)

// empty interface

// if you want to work with multiple types without knowing the type close to generics
type Empty interface{}

// type case switches

func typeSwitchExample() {
	var i interface{} = 0
	switch i.(type) {
	case int:
		println("i ia an integer")
	case string:
		println("i is a string")
	default:
		println("its something else")
	}
}

// when writing an interface if we use a value type instace all the the methods have to have value type receivers (v Value)
// on the other hand when i have a reference type receiver it parses all value and pointer type receivers (v * Value)

// best practices
// prefer many small interfaces
// accept interfaces wherever possible in functions and methods
