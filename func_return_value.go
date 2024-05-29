package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Hoshi") // name disimpan dulu di reslut
	fmt.Println(result)

	fmt.Println(getHello("Woozi"))
}
