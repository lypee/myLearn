package main

import (
	"log"
	"math"
)

func main() {
	//log.Println(max(1,2))
	str1 := "abcde"
	str2 := "ace"
	log.Println(longestCommonSubsequence3(str1, str2))
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

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	dp := make([][]int, t1Len+1)
	for index := range dp {
		dp[index] = make([]int, t2Len+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[t1Len][t2Len]
}

func longestCommonSubsequence3(text1 string, text2 string) int {
	m , n := len(text1)  ,len(text2)
	dp := make([][]int , m+1)
	for i := range dp{
		dp[i] = make([]int , n+1)
	}
	for i , c1 := range text1{
		for j , c2 := range text2{
			if c1 == c2{
				dp[i+1][j+1] = dp[i][j] +1
			}else{
				dp[i+1][j+1] = max(dp[i+1][j] , dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}


