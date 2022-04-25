package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	arr := []int{1, 1, 1, 0}
	target := -100
	log.Println(threeSumClosest(arr, target))
}

func threeSumClosest(nums []int, target int) int {
	res := 0
	length := len(nums)
	sort.Ints(nums)
	minClosest := math.MaxInt64
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		r := length - 1

		for l := i + 1; l < length; l++ {
			if l < r && math.Abs(float64(target-(nums[l]+nums[r]+nums[i]))) < math.Abs(float64(minClosest)) {
				res = nums[l] + nums[r] + nums[i]
				minClosest = target - (nums[l] + nums[r] + nums[i])
			}
			if nums[i]+nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
