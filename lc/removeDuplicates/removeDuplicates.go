package main

import "log"

// lc 80

func main() {
	nums := []int{1, 1, 2, 3, 4, 4}
	res := []int{}
	res = append(res, nums[0:1]...)
	log.Println(res)
	removeDuplicates(nums)
}
func removeDuplicates(nums []int) int {
	res := make([]int, 0)
	l, r := 0, 0
	for l = 0; l < len(nums)-1 && r < len(nums)-1; l++ {
		if nums[l] == nums[l+1] {
			r = l + 1
			for ; nums[r] == nums[r+1]; r++ {
			}
			res = append(res, nums[l:r+1]...)
			l = r
			continue
		}
		res = append(res, nums[l])
	}
	log.Println(res)
	return len(res)
}
