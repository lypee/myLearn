package main

import (
	"log"
	"math"
)

func main() {
	log.Println(reverse(0))
	log.Println(reverse(-321))

	log.Println(reverse(120))

}
func reverse(x int) int {
	res := 0
	var isNegative bool
	if x < 0 {
		isNegative = true
		x *= -1
	}
	for x != 0 {
		res = res*10 + x%10
		x /= 10
		if res >= math.MaxInt32{
			return 0
		}
	}
	if isNegative {
		res *= -1
	}
	return res
}
