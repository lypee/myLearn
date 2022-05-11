package main

import "log"

func main() {
	s := "babad"
	log.Println(longestPalindrome(s))
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

//
//func expend(s string, l, r int) (int, int) {
//
//	for l >= 0 && r < len(s) && s[l] == s[r] {
//		l--
//		r++
//	}
//	return l + 1, r - 1
//}

func longestPalindrome2(s string) string {
	length := len(s)
	start, end := 0, 0
	for i := 0; i < length-1; i++ {
		l1, r1 := expend(s, i, i)
		l2, r2 := expend(s, i, i+1)
		if end-start < r1-l1 {
			start, end = l1, r1
		}
		if end-start < r2-l2 {
			start, end = l2, r2
		}
	}
	return s[start:end]
}

func expend(s string, l, r int) (int, int) {
	for l <= r && l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}
