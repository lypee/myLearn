package main

import "log"

func main(){
	arrs := []int{3,2,1,4,0}
	arr := order(arrs)
	log.Println(arr)
}

// 3、2、1、4、0
func order(arrs []int) []int {
	for i := 0 ; i  < len(arrs) ; i++{
		if arrs[i] == i{
			continue
		}
		arrs[arrs[i]] , arrs[i] = arrs[i] , arrs[arrs[i]]
	}
	return arrs
}

func minWindow(s string, t string) string {
	return ""
}
