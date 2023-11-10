package basic

import "fmt"

func Looping() {
	// perulangan

	// for => for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke", i)
	}

	// for => while loop
	var i int8 = 1
	for i <= 5 {
		fmt.Printf("Loop-%d\n", i)
		i++
	}
}
