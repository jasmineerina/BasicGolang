package main

import (
	"fmt"
	"sort"
)

// Function Description
// Complete the function getMaximumAmount in the editor below.
// getMaximumAmount has the following parameter:
// int quantity[n]: the number of items of each type

// Returns
// long_int: the maximum revenue possible

// Example
// Consider n = 3, m = 4, quantity = [1, 2, 4].

func getMaximumAmount(quantity []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(quantity))) // urut slice menurun dr besar ke kecil
	var revenue int

	for i := 0; i < len(quantity); i++ { // cek dr arr quantity
		price := int(i + 1) // untuk mulai dr indeks 1
		revenue += int(quantity[i]) * price
	}

	return revenue
}

func main() {
	quantity := []int{1, 2, 4}
	result := getMaximumAmount(quantity)

	fmt.Println("max revenue possible:", result)
}
