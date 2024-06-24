package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	var p int

	//     cari 1/3 dari n
	consume := n / 3
	//     hasil consume di + n
	totalAmount := n + consume
	//     totalAmount dikurang 1000ml soda, untuk mengetahui B berapa ml
	p = 1000 - totalAmount

	fmt.Println(p)

}
