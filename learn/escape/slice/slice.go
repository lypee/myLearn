package main

func main() {
	escapeEntrance()
	select {}
}
// longestPalindrome
func escapeEntrance() {
	s1 := make([]string, 10, 10)
	escapeTest(s1)
}

func escapeTest(s1 []string) []string {
	s2 := make([]string, 0, 0)
	for idx, _ := range s1 {
		s2 = append(s2, s1[idx])
	}
	return s2
}
