package main

import "log"

func main(){
	arr := []int{1,2,3,4}
	log.Println(exchange(arr))
}

func exchange(nums []int) []int {
	left, right := 0, len(nums) - 1
	for left < right{
		for ; !isOdd(nums[right]) && right > left ;right-- {}
		for ; isOdd(nums[left]) && right > left ; left++ {}
		nums[left] , nums[right] = nums[right] , nums[left]
	}
	return nums
}

func isOdd(num int) bool {
	if num%2 == 1 {
		return true
	}
	return false
}
