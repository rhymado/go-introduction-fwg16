package pointer

import "fmt"

func Pointer() {
	var value int = 12

	// variabel pointer
	// referencing => mengambil address dari value => &
	// dereferencing => mengambil value dari adddress => *
	var addressValue *int = &value

	// pass by value
	copyValue := value

	fmt.Println(value)
	fmt.Println(*addressValue)
	fmt.Println(&addressValue)

	*addressValue = 20
	fmt.Println(value)
	fmt.Println(copyValue)
}
