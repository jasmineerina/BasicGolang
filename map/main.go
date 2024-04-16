package main

import "fmt"

func main() {
	var myMap map[string]int // var NAMAVARIABEL map[key]value
	myMap = map[string]int{} // deklarasiin lagi

	myMap["ayam"] = 5
	myMap["daging"] = 4
	myMap["sayur"] = 4
	myMap["ikan"] = 9

	fmt.Println(myMap["ayam"])
	fmt.Println("=============")

	language := map[string]string{
		"ruby": "is easy",
		"go":   "is convinence",
		"java": "is old but cool",
	}

	for key, value := range language {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	fmt.Println("=============")

	delete(language, "ruby") // untuk delete sebuah data

	for key, value := range language {
		fmt.Println("Key: ", key, "Value: ", value)
	} // harus dipanggil lagi untuk melihat data yang sudah di delete

	value, isAvail := language["go"]
	fmt.Println(isAvail) // hasil akan false, kalo tidak ada key yg sesuai di map
	fmt.Println(value)
}
