package main

import "log"

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	log.Println(searchRange(arr, 7))
}

func searchRange(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{0 ,0}
	}
	left, right := 0, len(nums)-1
	l, r := -1, -1
	for left < right {
		if nums[left] == target && l == -1 {
			l = left

		}
		if nums[right] == target && r == -1 {
			r = right
		}
		if l == -1 {
			// left 动
			left++
		}
		if r == -1 {
			right--
		}
		if l != -1 && r != -1 {
			return []int{l, r}
		}
	}
	return []int{l, r}
}
