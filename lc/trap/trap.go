package main

import (
	"math"
)

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for l <= r {
		leftMax = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		if height[l] < height[r] {
			ans += (leftMax - height[l])
			l++
		} else {
			ans += (rightMax - height[r])
			r--
		}
	}
	return ans
}

func tra2p(height []int) int{
	length := len(height)
	lMax , rMax := 0 ,0
	res := 0
	l , r := 0 ,length - 1
	for l <= r{
		lMax = max(lMax , height[l])
		rMax = max(rMax , height[r])
		if height[l] < height[r]{
			res += lMax - height[l]
			l++
		}else{
			res += rMax - height[r]
			r--
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
