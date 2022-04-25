package main

import "log"

func main() {
	log.Println(numWays(7))
}
func numWays(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
		dp[i] %= 1000000007
	}
	return dp[n]
}
