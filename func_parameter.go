package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Halo", firstName, lastName)
}

func main() {
	sayHelloTo("Kwon", "Soonyoung")
	sayHelloTo("Lee", "Seokmin")
}
