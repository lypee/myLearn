package main

// lc 169 多数元素
func majorityElement(nums []int) int {
	nMap := map[int]int{}
	length := len(nums)
	for i := range nums{
		nMap[nums[i]]++
		if nMap[nums[i]] > length /2 {
			return nums[i]
		}
	}
	return 0
}