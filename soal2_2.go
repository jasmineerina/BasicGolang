package main

import "fmt"

func buatSegitiga(n int) int {
	fmt.Println(" ", "*")
	for i := 0; i < n; i++ {

		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		fmt.Print("/")
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}

		fmt.Print("\\")
		fmt.Println()
	}
	return n
}

func main() {
	fmt.Println(buatSegitiga(3))
	fmt.Println(buatSegitiga(5))
}
