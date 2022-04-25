package main

import (
	"log"
	"math"
	"time"
)

//func main(){
//	arrs := []int{1,8,6,2,5,4,8,3,7}
//	log.Println(maxWater(arrs))
//	arrs2 := []int{1,1}
//	arrs3 := []int{4,3,2,1,4}
//	arrs4 := []int{1,2,1}
//	log.Println(maxWater(arrs2))
//	log.Println(maxWater(arrs3))
//	log.Println(maxWater(arrs4))
//
//	return
//}


func main() {
	values := []int{1,2,3}

	for _, v := range values {
		go func() {
			log.Println(v)
			return
		}()
	}
	time.Sleep(time.Second)
	return
}

func min(nums ...int) int {
	res := math.MaxInt64
	for i := 0 ; i < len(nums) ;  i++{
		if res > nums[i]{
			res = nums[i]
		}
	}
	return res
}

func maxWater(arrs []int) int {
	length := len(arrs)
	if length < 1 {
		return 0
	}
	left , right  := 0 , length - 1
	max := 0
	for left < right{
		area := min(arrs[left] , arrs[right]) * (right - left)
		if area > max{
			max = area
		}

		if arrs[left] > arrs[right]{
			right--
		}
		left ++
	}
	return max
}