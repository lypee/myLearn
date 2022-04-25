package main

import (
	"log"
	"math"
)

// lc 128 最长连续序列

func main() {
	arr := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	quickSort(arr, 0 ,  len(arr)  -1 )
	log.Println(arr)
}
func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	nMap := make(map[int]struct{}, 0)
	minNum := nums[0]
	nMap[nums[0]] = struct{}{}
	for i := 1; i < len(nums); i++ {
		minNum = min(minNum, nums[i])
		nMap[nums[i]] = struct{}{}
	}
	maxLen := 0
	for key := range nMap {
		if _, exist := nMap[key-1]; !exist {
			l := 1
			v := key + 1
			for {
				if _, ex := nMap[v]; !ex {
					break
				}
				l++
				v++

			}
			maxLen = max(maxLen, l)
		}
	}
	return maxLen
}

func min(arr ...int) int {
	res := math.MaxInt64
	for i := range arr {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

func quickSort(arr []int, left, right int) {
	length := len(arr)
	if length < 1 {
		return
	}

	l, r := left, right
	num := arr[0]
	for l < r {
		for l < r && arr[r] > num {
			r--
		}

		for l < r && arr[l] < num {
			l++
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	quickSort(arr, l, right)
	quickSort(arr, left, r)
	return
}
