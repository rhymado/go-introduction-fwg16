package basic

import "fmt"

func Declaration() {
	// manifest type
	const greeting string = "Hello World"
	var firstName string = "Fazz"
	var lastName string = "Track"
	// type interface
	isMarried := false

	fmt.Println(greeting)
	fmt.Printf("%s-%s\n", firstName, lastName)
	fmt.Print(isMarried, isMarried, "\n")
}
