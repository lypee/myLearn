package main

import "log"

// lc 283.移动零

func main(){
	arr := []int{0,1,0,3,12}
	moveZeroes(arr)
}

// 0,1,0,3,12
func moveZeroes(nums []int) {
	length := len(nums)
	l, r := 0, 0
	for r < length &&  l <= r {
		if nums[r] != 0{
			nums[l] , nums[r] = nums[r] , nums[l]
			l++
		}
		r++
	}
	log.Println(nums)
	return
}

func moveZeroes2(nums []int) {
	l , r ,length := 0 ,0 , len(nums)
	for r < length {
		if nums[r] != 0 {
			nums[l] , nums[r] = nums[r] , nums[l]
			l++
		}
		r++
	}
}

