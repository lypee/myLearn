package main

import (
	"log"
	"math"
)

func main(){
	log.Println(cuttingRope(10))
	return
}
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int,n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j <= i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(arrs ...int) int {
	res := math.MinInt64
	for i := 0; i < len(arrs); i++ {
		if arrs[i] > res {
			res = arrs[i]
		}
	}
	return res
}
