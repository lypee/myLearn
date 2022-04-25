package main

import "log"

func main() {
	s := "abc"
	log.Println(countSubstrings(s))
}
func countSubstrings(s string) int {
	ans := 0
	length := len(s)
	for i := 0; i < length*2-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func countSubstrings2(s string) int {
	ans := 0
	length := len(s)
	for i := 0; i < length*2-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
