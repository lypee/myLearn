package wordBreak

func wordBreak(s string, wordDict []string) bool {
	mmap := make(map[string]bool, 0)
	for i := range wordDict {
		mmap[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if mmap[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) bool {
	length := len(s)
	dp := make([]bool ,length + 1 )
	dp[0] = true
	mmap := make(map[string]bool, 0)
	for i := 0 ; i < length ; i++{
		for j := 0; j < i ; j++{
			if dp[j] && mmap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[length]
}
