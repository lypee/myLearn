package main

import "log"

func main(){
	arr := []int{3,1,2,3}
	log.Println(findRepeatNumber(arr))
}

func findRepeatNumber(nums []int) int {
	res := -1
	if len(nums) < 2 {
		return res
	}
	for i := 0 ; i < len(nums)  ; i++{
		num := nums[i]
		if num == i {
			continue
		}
		if num == nums[num]{
			res = num
			return res
		}
		nums[i] , nums[num] = nums[num] , nums[i]
	}
	return res
}