package findPeakElement

func findPeakElement(nums []int) int {
	idx := 0
	for i , val := range nums{
		if val > nums[idx]{
			idx = i
		}
	}
	return idx
}
