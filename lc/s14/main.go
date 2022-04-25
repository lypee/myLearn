package main

import (
	"fmt"
	"math"
)

func main(){
	n := 10
	fmt.Println(cuttingRope(n))
}
func cuttingRope(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 3
	for len := 4 ; len <= n  ; len++ {
		for j := 2; j <= len-1; j++ {
			dp[len] = max(dp[len], dp[len-j]*dp[j])
		}
	}
	return dp[n]
}

func max(arr ...int)int{
	max := math.MinInt64
	for _ , num := range arr{
		if max < num{
			max = num
		}
	}
	return max
}
