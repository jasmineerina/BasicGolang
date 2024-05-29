package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Jasmine"
	middleName = "Erina"
	lastName = "Firdaus"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
