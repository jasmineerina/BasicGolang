package main

import (
	"belajar-pertama/learning"
	"fmt"
)

func main() {
	subject := learning.Subject{Id: 1, Title: "cobain aja"}

	fmt.Println(subject.Id)
	fmt.Println(subject.Title)
}
