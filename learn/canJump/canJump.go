package canJump

func canJump2(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length)
	dp[0] = true
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[length-1]
}
