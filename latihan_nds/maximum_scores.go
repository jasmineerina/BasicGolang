package main

import "fmt"

// For an array arr of n integers,
// the mathematician can perform the following move on the array:

// 1. Choose an index i (0 <= i < length(arr)) and add arr[i] to their scores.
// 2. Discard either the left partition (i.e. arr[0, 1... i -1])
// or the right partition (i.e. arr[i+1,i+2... length(arr)-1].
// The partition discarded can be empty too.
// The selected partition then becomes the new value of arr
// and is used for subsequent operations.

// Starting with an initial score of 0, the mathematician wishes to find the maximum
// achievable score after k moves.

// Example
// Consider n = 6, arr = [4, 6, -10, -1, 10, -20], and k = 4

func maxScore(arr []int, k int) int {
	n := len(arr)
	scores := make([]int, n+1) // untuk membuat array (baru) score yang panjangnya sampai n+1

	for i := 1; i < n; i++ { // untuk jumlahin kumulatif (nilai total dr arr scores)
		scores[i] = scores[i-1] + arr[i-1]
	}

	dp := make([][]int, n+1) // untuk matriks
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ { // perulangan i j utk isi nilai dari matriks
		for j := 1; j <= k; j++ {
			maxScore := 0 // simpan skor maks dr setiap perulangan

			for x := 0; x <= i; x++ {
				leftScore := scores[x]              // penjumlahan dri awal sampai x
				rightScore := scores[i] - scores[x] // penjumlahan dri x+1 sampai i

				score := leftScore + rightScore + dp[x][j-1]
				if score > maxScore {
					maxScore = score
				}
			}
			dp[i][j] = maxScore
		}
	}

	return dp[n][k]

}

func main() {
	array := []int{4, 6, -10, -1, 10, -20}

	fmt.Println(maxScore(array, 4)) // output 19
}
