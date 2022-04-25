package main

import (
	"log"
	"math"
)

func main() {
	str := "pwwkew"
	log.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	left, res := 0, 0
	mmap := make(map[int32]int, 0)
	for i, c := range s {
		if numIndex, exist := mmap[c]; exist {
			left = max(left, numIndex+1)
		}
		res = max(res, i-left+1)
		mmap[c] = i
	}
	return res
}
func max(arr ...int) int {
	res := math.MinInt64
	for _, num := range arr {
		if num > res {
			res = num
		}
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	cMap := make(map[int32]int , 0)
	left ,res := 0 ,0
	for i , c := range s{
		if index , exist := cMap[c] ; exist{
			left = max(left, index+1)
		}
		res = max(res,  i - left + 1 )
		cMap[c] = i
	}
	return res
}


