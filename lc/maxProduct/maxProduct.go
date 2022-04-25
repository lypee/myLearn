package main

import "math"

func maxProduct(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	numMin, numMax, ans := nums[0], nums[0], nums[0]
	for i := 1 ; i < length ; i++  {
		nMin, nMax := numMin, numMax
		numMax = max(nMax*nums[i] , max(nums[i] , nMin * nums[i]))
		numMin = min(nMin*nums[i] , min(nums[i] , nMax * nums[i]))
		ans = max(numMax, numMin ,ans )
	}
	return ans
}

func min(arr ...int) int {
	res := math.MaxInt64
	for i := range arr {
		if res > arr[i] {
			res = arr[i]
		}
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
