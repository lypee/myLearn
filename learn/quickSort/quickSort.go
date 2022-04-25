package quickSort

func getIndex(nums []int, left, right int) int {
	base := nums[left]
	start := left

	for left < right {
		for nums[right] >= base && left < right {
			right--
		}
		for nums[left] <= base && left < right {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[start] = nums[left]
	nums[left] = base
	return left
}

func findTopKInQuickSort(nums []int, left, right, k int) int {
	index := getIndex(nums, left, right)
	if right-index == k {
		return index
	}
	if right-index > k {
		return findTopKInQuickSort(nums, index+1, right, k)
	}
	return findTopKInQuickSort(nums, left, index, k-(right-index))
}
