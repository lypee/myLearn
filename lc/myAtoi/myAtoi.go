package main

import (
	"log"
	"math"
)

func main() {
	str := " -42"
	log.Println(myAtoi(str))
}

func myAtoi(s string) int {
	var res int32
	var isNegative, start bool
	for _, c := range s {
		if c == '-' {
			isNegative = true
		}
		if !isNumber(c) && start {
			break
		}
		if !isNumber(c) || c == ' ' {
			continue
		}

		start = true
		beforeRes := res
		res *= 10
		res += c - '0'

		if beforeRes > res { // 超限
			if isNegative {
				return math.MinInt64
			} else {
				return math.MaxInt64
			}
		}

	}
	if isNegative {
		return int(res * int32(-1))
	}
	return int(res)
}

func isNumber(c int32) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isEnglish(c int32) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
