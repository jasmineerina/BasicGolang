package main

import "fmt"

// slice mirip array
func main() {
	// gamingConsoles := []string{"Nintendo Switch", "PlaysStation 5", "Xbox", "PSP"}
	// fmt.Println(gamingConsoles)
	// ini kalo mau dibuat langsung

	// pakai append untuk dinamis
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "PlayStation5")
	gamingConsoles = append(gamingConsoles, "Xbox")
	gamingConsoles = append(gamingConsoles, "PSP")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

}
