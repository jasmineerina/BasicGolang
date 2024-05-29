package main

import "fmt"

func getFullName() (string, string) {
	return "Kwon", "Horanghae"
}

func main() {
	firstName, _ := getFullName() // kalo ga butuh lastName, pakai _
	fmt.Println(firstName)

	firstName, lastName := getFullName() // kalo butuh seluruh return
	fmt.Println(firstName, lastName)
}
