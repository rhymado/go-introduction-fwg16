package basic

import "fmt"

func Conditional() {
	// pengondisian

	var username string = "user12"
	password := "abcd123"

	if username != "user1" {
		fmt.Println("username tidak ditemukan")
	} else if password != "abcd1234" {
		fmt.Println("password tidak cocok")
	} else {
		fmt.Println("login berhasil")
	}
}
