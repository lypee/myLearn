package maxProfit

import (
	"math"
)

// lc 121 买卖股票的最佳时机

func maxProfit(prices []int) int {
	length := len(prices)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// dp[i][j] 表示当天的现金数 j=0不持股 ，j=1持股
	dp[0][0] = 0
	dp[0][1] = -1 * prices[0]
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], -1*prices[i])
	}
	return dp[length-1][0]
}
func maxProfit11(prices []int) int {
	length := len(prices)
	dp := make([][]int , length)
	for i := range dp{
		dp[i] = make([]int , 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1 ; i < length ; i++{
		dp[i][0] = max(dp[i-1][0] , dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1] , dp[i-1][0] -prices[i])
	}
	return dp[length-1][0]
}


func max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

// lc 122 买卖股票的最佳时机二

func maxProfit2(prices []int) int {
	length := len(prices)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -1 * prices[0]
	for i := 0; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[length-1][0]
}

// lc 123 买卖股票的最佳时机三
func maxProfit3(prices []int) int {
	length := len(prices)
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < length; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}


