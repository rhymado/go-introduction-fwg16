package models

import "fmt"

type Cat struct{}

func (c Cat) Voice() {
	fmt.Println("Meow")
}

func (c Cat) Sleep() {
	fmt.Println("zzz")
}

func (c Cat) Attack() {
	fmt.Println("Meooowr")
}
