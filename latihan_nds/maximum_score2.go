package main

import (
	"fmt"
	"math"
)

func solution(arr []int, k int) int {
	n := len(arr)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt64
		}
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for m := 1; m <= k; m++ {
		for j := 0; j < n; j++ {
			for i := 0; i <= j; i++ {
				// currentSum := prefixSum[j+1] - prefixSum[i]
				dp[m][j+1] = max(dp[m][j+1], arr[i]+dp[m-1][i])
			}
		}
	}

	result := math.MinInt64
	for j := 0; j <= n; j++ {
		result = max(result, dp[k][j])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// n := 6
	arr := []int{4, 6, -10, -1, 10, -20}
	k := 4
	fmt.Println(solution(arr, k)) // output 19
}
