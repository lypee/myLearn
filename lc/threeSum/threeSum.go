package main

import (
	"log"
	"sort"
)

func main() {
	arrs := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(arrs)
	log.Println(res)
	return
}
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	n := len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue //去重
		}

		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			if second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

// -1,0,1,2,-1,-4
func threeSum2(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := nums[i] * -1
		r := length - 1
		for l := i + 1; l < r; l++ {
			if l > i+1 && nums[l] == nums[l-1] {
				continue
			}
			if l > i+1 && nums[l]+nums[r] > target {
				r--
			}
			if l == r {
				break
			}
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			}
		}
	}
	return res
}
func threeSum3(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := nums[i] * -1
		r := length - 1
		for l := i + 1; l < length; l++ {
			if l > i+1 && nums[l] == nums[l-1] {
				continue
			}
			for l < r && nums[l]+nums[r] > target {
				r--
			}
			if l >= r{
				break
			}
			if target == nums[l]+nums[r] {
				res = append(res, []int{nums[l], nums[r], nums[i]})
			}
		}
	}
	return res
}
