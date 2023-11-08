package main

import "fmt"

func main() {
	// // manifest type
	// const greeting string = "Hello World"
	// var firstName string = "Fazz"
	// var lastName string = "Track"
	// // type interface
	// isMarried := false

	// fmt.Println(greeting)
	// fmt.Printf("%s-%s\n", firstName, lastName)
	// fmt.Print(isMarried, isMarried)

	// pengondisian

	// var username string = "user12"
	// password := "abcd123"

	// if username != "user1" {
	// 	fmt.Println("username tidak ditemukan")
	// } else if password != "abcd1234" {
	// 	fmt.Println("password tidak cocok")
	// } else {
	// 	fmt.Println("login berhasil")
	// }

	// perulangan

	// // for => for loop
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Perulangan ke", i)
	// }

	// // for => while loop
	// var i int8 = 1
	// for i <= 5 {
	// 	fmt.Printf("Loop-%d\n", i)
	// 	i++
	// }

	// array (value type)
	// var kendaraan [3]string
	// kendaraan[0] = "Mobil"
	// kendaraan[1] = "Becak"
	// kendaraan[2] = "Motor"
	// // kendaraan[3] = "Bis"

	// fmt.Println(kendaraan)
	// fmt.Println(len(kendaraan))

	// var umur = [2]int8{10, 12}
	// umurSMP := umur
	// umurSMP[0] = 14
	// fmt.Println(umur, umurSMP)

	// slice (reference type)
	// var buah = []string{"Mangga", "Semangka", "Melon", "Jambu", "Durian"}
	// buahFavorite := buah[1:3]
	// buah[1] = "Pepaya"
	// buahFavorite[0] = "Pepaya"
	// fmt.Println(buahFavorite, buah)

	// slice of array
	// kendaraanku := kendaraan[1:3]
	// kendaraan[1] = "Kapal"
	// kendaraanku[0] = "Kapal"
	// fmt.Println(kendaraanku, kendaraan)

	// for => foreach
	// for idx, val := range buah {
	// 	if idx%2 != 0 {
	// 		fmt.Println(val, "Ganjil")
	// 		continue
	// 	}
	// 	fmt.Println(val, "Genap")
	// }
	learnMap()

	firstName := "FWG"
	lastName := "16"
	fullName, _ := setFullName(firstName, lastName)
	printFullName(fullName)
}

// return function
func setFullName(firstName string, lastName string) (string, int) {
	fullName := firstName + " " + lastName
	return fullName, 12
}

// void function
func printFullName(fullName string) {
	fmt.Printf("%s", fullName)
}

func learnMap() {
	// Map
	var response = make(map[string]string)
	response["status"] = "200"
	response["message"] = "Sukses"
	fmt.Println(response["message"])

	data := map[int]string{
		1: "satu",
		2: "dua",
		3: "tiga",
	}
	fmt.Println(data[3])
	fmt.Println(len(response), len(data))
}
