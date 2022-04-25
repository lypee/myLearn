package main

import "log"

func main() {
	s := "babad"
	log.Println(longestPalindrome(s))
}
func longestPalindrome2(s string) string {
	if len(s) < 1 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expand2(s, i, i)
		left2, right2 := expand2(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
		if right1-left1 > end-start {
			start, end = left1, right1
		}
	}
	return s[start : end+1]
}

func expand2(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		s1, e1 := expend(s, i, i)
		s2, e2 := expend(s, i, i+1)
		if end-start < e1-s1 {
			start, end = s1, e1
		}
		if end-start < e2-s2 {
			start, end = s2, e2
		}
	}

	return s[start : end+1]
}

func expend(s string, l, r int) (int, int) {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
