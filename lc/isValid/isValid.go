// lc 20
package main

import "log"

func main() {
	s := "()[]{}"
	log.Println(isValid(s))
}

// '('，')'，'{'，'}'，'['，']'
func isValid(s string) bool {

	stack := make([]int32, 0, len(s))
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')':
			if len(stack) < 1 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) < 1 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		case ']':
			if len(stack) < 1 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) < 1 {
		return true
	}
	return false
}
