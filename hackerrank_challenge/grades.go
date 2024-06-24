package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var score int32
	fmt.Scan(&score)
	// var grade string

	if score <= 59 && score > 0 {
		fmt.Println("F")
	} else if score <= 69 && score > 60 {
		fmt.Println("D")
	} else if score <= 79 && score > 70 {
		fmt.Println("C")
	} else if score <= 89 && score > 80 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}

}
