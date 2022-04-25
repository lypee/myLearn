package main

// 搜索旋转数组 33\81
import "log"

func main() {
	arrs := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	log.Println(search3(arrs, target))
	return
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target {
			return left
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return -1
}

// 15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14
func search3(arr []int, target int) int {
	l , r := 0 , len(arr) -1
	for l <= r {
		if arr[l] == target{
			return l
		}
		mid := (l + r) >> 1
		if arr[mid] == target{
			return mid
		}else if arr[l] < arr[mid]{
			if arr[l] <= target && target < arr[mid]{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}else if arr[l] > arr[mid]{
			if arr[mid] < target && target <= arr[r]{
				l = mid + 1
			}else{
				r = mid - 1
			}
		}else{
			l++
		}
	}
	return -1
}
