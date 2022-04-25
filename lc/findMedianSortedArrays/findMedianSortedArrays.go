package main

import "log"

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2}
	log.Println(findMedianSortedArrays2(arr1, arr2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	m, n := 0, 0
	left, right := 0, 0
	for i := 0; i <= (len1+len2)/2; i++ {
		left = right
		if m < len1 && (n >= len2 || nums1[m] < nums2[n]) {
			right = nums1[m]
			m++
		} else {
			right = nums2[n]
			n++
		}
	}
	if (len1+len2)%2 == 1 {
		return float64(right)
	}
	return float64(right+left) / 2.0
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	leftNum, rightNum, left, right := 0, 0, 0, 0
	for i := 0; i < (len1+len2)/2; i++ {
		leftNum = rightNum
		// 动left
		if left < len1 && (right >= len2 || nums1[left] < nums2[right]) {
			rightNum = nums1[left]
			left++
		} else {
			rightNum = nums2[right]
			right++
		}
	}
	if (len1+len2)%2 == 1 {
		return float64(rightNum)
	}
	return (float64(rightNum) + float64(leftNum)) / 2.0
}
