package basic

import "fmt"

func Collection() {
	// array (value type)
	var kendaraan [3]string
	kendaraan[0] = "Mobil"
	kendaraan[1] = "Becak"
	kendaraan[2] = "Motor"
	// kendaraan[3] = "Bis"

	fmt.Println(kendaraan)
	fmt.Println(len(kendaraan))

	var umur = [2]int8{10, 12}
	umurSMP := umur
	umurSMP[0] = 14
	fmt.Println(umur, umurSMP)

	// slice (reference type)
	var buah = []string{"Mangga", "Semangka", "Melon", "Jambu", "Durian"}
	buahFavorite := buah[1:3]
	buah[1] = "Pepaya"
	buahFavorite[0] = "Pepaya"
	fmt.Println(buahFavorite, buah)

	// slice of array
	kendaraanku := kendaraan[1:3]
	kendaraan[1] = "Kapal"
	kendaraanku[0] = "Kapal"
	fmt.Println(kendaraanku, kendaraan)

	// for => foreach
	for idx, val := range buah {
		if idx%2 != 0 {
			fmt.Println(val, "Ganjil")
			continue
		}
		fmt.Println(val, "Genap")
	}
}
