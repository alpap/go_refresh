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
