package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Golang"
	// languages[1] = "Ruby"
	// languages[2] = "C#"
	// languages[3] = "Javascript"
	// languages[4] = "Python"

	languages := [...]string{
		"C++",
		"Ruby",
		"Python",
		"Javascript",
		"Golang",
		"C#",
	}

	for index, lang := range languages {
		fmt.Println("index :", index, " language :", lang)
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)
}
