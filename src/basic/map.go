package basic

import "fmt"

func LearnMap() {
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
