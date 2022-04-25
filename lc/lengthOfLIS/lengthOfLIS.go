package main

import (
	"log"
	"math"
)

func main() {
	arr := []int{7, 7, 7, 7, 7, 7, 7}
	log.Println(lengthOfLIS(arr))
}

// lc 300 最长递增子序列
func lengthOfLIS(nums []int) int {
	res := 0
	if len(nums) < 1 {
		return res
	}
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func lengthOfLIS2(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
