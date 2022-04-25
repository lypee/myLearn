package main

import (
	"log"
	"math"
)

func main(){
	arr := []int{2,3,4,5,18,17,6}
	log.Println(maxArea(arr))
}
func maxArea(height []int) int {
	res := math.MinInt64
	left, right := 0, len(height)-1
	for left < right{
		abs := min(height[left] , height[right])  *(right - left)
		res = max(res , abs)
		if height[left] < height[right]{
			left++
		}else{
			right--
		}
	}
	return res
}

func min(arr ...int) int {
	res := math.MaxInt64
	for _, num := range arr {
		if num < res {
			res = num
		}
	}
	return res
}

func max(arr ...int) int {
	res := math.MinInt64
	for _, num := range arr {
		if num > res {
			res = num
		}
	}
	return res
}

func maxArea2(height []int) int{
	length := len(height)
	ans := 0
	l , r := 0 ,length  - 1
	for l <= r {
		abs := min(height[l] , height[r]) * (r - l)
		ans = max(ans ,abs)
		if height[l] < height[r]{
			l++
		}else{
			r--
		}
	}
	return ans
}