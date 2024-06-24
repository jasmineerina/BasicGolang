package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	var hasil int

	for i := 0; i <= n; i++ {
		hasil += i
	}

	fmt.Println(hasil)
}
