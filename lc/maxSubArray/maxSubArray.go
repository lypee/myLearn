package main

import (
	"log"
	"math"
)

func main() {
	arrs := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	log.Println(maxSubArray2(arrs))
	return
}
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	length, res := len(nums), nums[0]
	dp := make([]int, length)
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
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

func maxSubArray2(nums []int) int {
	dp := make([]int , len(nums))
	dp[0] = nums[0]
	res := 0
	for i := 1 ;i <  len(nums) ; i++{
		dp[i] = max(dp[i-1]  + nums[i] ,nums[i])
		res = max(res,dp[i])
	}
	return res
}
