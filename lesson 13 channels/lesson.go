package main

func main() {
	// routines
	go routine()
}

func correctGoRoutine() {
	go func(message string) {
		println(message)
	}("correct way")
}
