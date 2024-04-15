package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Golang :", i)
	// }

	// contoh model untuk while
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Golang :", i)
	// 	i++
	// }

	// for each (array)
	title := "belajar go"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index :", index, " letter: ", string(letter))
	// 	}
	// }

	for index, letter := range title {
		letterString := string(letter)

		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
			fmt.Println("index :", index, " letter: ", string(letter))

		}
	}
}
