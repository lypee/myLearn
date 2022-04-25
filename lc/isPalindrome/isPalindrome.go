package main

import "log"

func main(){
	log.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xx := x
	res := 0
	for x != 0 {
		n := x % 10
		res = res *10 + n
		x /= 10
	}
	return res == xx
}
