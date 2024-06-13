package main

import "fmt"

func buatSegitiga(input int) int {
	fmt.Println(" ", "*")

	for i := 0; i < input; i++ {
		for j := 0; j < input-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Print("/")
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("\\")
	}
	return input
}

func main() {
	fmt.Println(buatSegitiga(3))
	fmt.Println(buatSegitiga(5))
}
