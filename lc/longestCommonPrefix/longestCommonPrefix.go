package main

import "log"

func main() {
	strs := []string{"ab", "a"}
	log.Println(longestCommonPrefix(strs))
	return
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	if minLen < 1 {
		return ""
	}
	//maxIndex := 0
	for i := 0 ; i < len(strs[0]) ; i++{
		for j := 1 ; j < len(strs) ; j++{
			if i == len(strs[j]) || strs[j][i] != strs[0][i]{
				return strs[0][:i]
			}
 		}
	}

	return strs[0]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	for i := 0 ; i < len(strs[0]) ; i++{
		for j :=  1 ;  j < len(strs) ; j++{
			if i == len(strs[j]) || strs[j][i] != strs[0][i]  {
				return strs[0][:i]
			}
		}
	}
	return  strs[0]
}
