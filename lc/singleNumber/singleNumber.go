package main

import "log"

func main() {
	//arr := []int{4,4,3,4}
	//log.Println(singleNumber2(arr))
	//log.Println(5 & 1)
	//log.Println(11 & 1)
	//log.Println(2 & 1)
	//log.Println(3 & 1)
	//
	log.Println(hammingWeight(11))
}

// lc 136 只出现一次的数字
func singleNumber2(nums []int) int{
	res := 0
	for i := range nums{
		res ^= nums[i]
	}
	return res
}

func singleNumber(nums []int) int {
	res, bit := 0, 0
	for i := 30; i >= 0; i-- {
		for _, num := range nums {
			bit += (num >> 1) & 1
		}
		res = res >> 1
		res += bit % 3
		bit = 0
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	numInt := int(num)
	for numInt != 0 {
		if numInt&1 == 1 {
			res += 1
		}
		numInt = numInt >> 1
	}
	return res
}


