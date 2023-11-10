package models

import "fmt"

// type user struct {
// 	username string
// 	age      int8
// 	address  string
// 	phone    string
// }

type User struct {
	Username string
	Age      int8
	Address  string
	Phone    string
}

func Struct() {
	// var andi user

	// andi.username = "andi"
	// andi.age = 8
	// andi.address = "Jakarta"
	// andi.phone = "+6280012345"

	doni := User{
		Username: "doni",
		Age:      14,
		Address:  "Bandung",
		Phone:    "+6286712345",
	}

	doni.Age = 16

	// fmt.Println(andi)
	fmt.Println(doni)
}
