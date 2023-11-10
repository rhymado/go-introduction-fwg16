package models

import "fmt"

type Dog struct{}

func (d Dog) Voice() {
	fmt.Println("Woof")
}

func (d Dog) Attack() {
	fmt.Println("Woof Woof Woof")
}
