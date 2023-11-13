package intermediate

import (
	"fmt"
	"math/rand"
)

func Channel() {
	var channel = make(chan int)

	go diceRoll(channel)
	diceResult := <-channel

	fmt.Printf("Hasil Dadu adalah %d", diceResult)
}

func diceRoll(ch chan int) {
	result := rand.Intn(6)
	ch <- result + 1
}
