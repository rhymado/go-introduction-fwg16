package intermediate

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}

func Routine() {
	fmt.Printf("Program Start\n\n")

	wg.Add(2)

	go printRandomNumber(6)
	go greeting()

	wg.Wait()

	fmt.Printf("\nProgram End")
}

func greeting() {
	defer wg.Done()
	// time.Sleep(time.Millisecond * 1000)
	fmt.Println("Hello World")
	// panic("Problems")
}

func printRandomNumber(limit int) {
	defer wg.Done()
	result := rand.Intn(limit)
	fmt.Println(result)
}
