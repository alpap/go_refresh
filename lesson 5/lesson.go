package main

import (
	"fmt"
)

func main() {
	// maps and struct
	// map keys can only be types that can be tested for equality
	// for example maps slices arrays cannot be map keys
	scores := map[string]int{
		"alex": 20,
		"john": 40,
		"kara": 50,
	}
	printOut(scores)
	// creating maps // used when you dont have an initialization
	makeScores := make(map[string]string)
	makeScores = map[string]string{
		"alex": "20",
		"john": "40",
		"kara": "50",
	}
	printOut(makeScores)
	// manipulate the values of the map
	// adding an element is at a random possition not in order of insertion
	scores["brian"] = 44
	printOut(scores["brian"])
	printOut(scores)
	// remove elements from a map
	delete(scores, "brian")
	printOut(scores)
	// after deleting the value you can still ask to return it but it will return 0 or empty string
	// this happens because it initializes a new variable of the selected type in that spot
	// so watch out and validate your
	printOut(scores["brian"])

	// validation of the map
	// ok will return true or false depending if the array was initialized
	a, ok := scores["brian"]
	printOut(a)
	printOut(ok)
	b, ok := scores["alex"]
	printOut(b)
	printOut(ok)

	printOut(len(scores))

	// maps as slices get passed by references

	c := scores
	printOut(c)
	delete(c, "alex")
	printOut(scores)
	// if you want to avoid this from happening use the make declaration as shown above

	printOut(scores)
	printOut(makeScores)
	delete(makeScores, "alex")
	printOut(scores)
	printOut(makeScores)

	// struct
	// check out the struct defined under the main function
	person := Person{
		id:      3,
		name:    "alex",
		surname: "papad",
	}
	printOut(person)
	printOut(person.name)

	// anonymus structs
	name := struct{ name string }{name: "john"}
	printOut(name)

	// structs are copied by value not by reference
	person2 := person
	person2.name = "jack"

	printOut(person)

	// copy by reference
	person3 := &person
	person3.name = "jack"
}

// if i want to export the struct i should name it with a capital letter as well as the fields as well if i want them to be populated outside the package

type Person struct {
	id      int
	name    string
	surname string
}

func printOut(value interface{}) {
	fmt.Printf("%+v\n", value)
}
