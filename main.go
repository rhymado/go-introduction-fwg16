package main

import "fwg16/src/intermediate"

type Animal interface {
	Voice()
	Attack()
}

func main() {
	// basic.Declaration()
	// basic.Conditional()
	// basic.Looping()
	// basic.Collection()
	// basic.LearnMap()

	// IIFE => immediately invoked function expression
	// defer bersifat LIFO (Last in First out)
	// defer func() {
	// 	number := 14
	// 	sent := fmt.Sprintf("%d tahun yang lalu", number)
	// 	fmt.Println(sent)
	// }()
	// defer fmt.Println("Program Berakhir")

	// os.Exit(1) //dipaksa untuk berhenti

	// firstName := "FWG"
	// lastName := "16"
	// fullName, _ := setFullName(firstName, lastName)
	// printFullName(fullName)
	// pointer.Pointer()
	// models.Struct()
	// var gilang = models.User{}

	// gilang.Username = "gilang"
	// gilang.Age = 24
	// gilang.Address = "Yogya"
	// gilang.Phone = "+627890123456"

	// var text = "Pointer merupakan alamat dari suatu value"
	// var newPost = models.CreatePost("Pointer", text, gilang)
	// fmt.Println(newPost.Content)

	// var newTitle = "Method"
	// var newContent = "Method merupakan function yang berada didalam Struct"
	// newPost.EditPost(newTitle, newContent)

	// fmt.Println(newPost)

	// cat := models.Cat{}
	// dog := models.Dog{}

	// cat.Sleep()
	// dog.Attack()

	// animalVoice(cat)
	// animalVoice(dog)

	// var variable interface{} = 1
	// var variable any = 1

	// fmt.Println(variable)
	// variable = "satu"
	// fmt.Println(variable)
	// variable = true
	// fmt.Println(variable)

	// result := 10

	// result += variable.(int)
	// fmt.Println(result)

	// str := "Hello"
	// str += strconv.Itoa(variable.(int))
	// fmt.Println(str)

	// intermediate.Routine()
	// intermediate.Channel()
	intermediate.Racing()
}

func AnimalVoice(animal Animal) {
	animal.Voice()
}

// // return function
// func setFullName(firstName string, lastName string) (string, int) {
// 	fullName := firstName + " " + lastName
// 	return fullName, 12
// }

// // void function
// func printFullName(fullName string) {
// 	fmt.Printf("%s\n", fullName)
// }
