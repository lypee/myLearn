package s453

import "math"

func minMoves(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	minNum := min(nums...)
	cnt := 0
	for _ , num := range  nums{
		if num > minNum {
			cnt += num - minNum
		}
	}
	return cnt
}

func min(arr ...int) int{
	min := math.MaxInt64
	for _ , num := range arr{
		if min > num{
			min = num
		}
	}
	return min
}