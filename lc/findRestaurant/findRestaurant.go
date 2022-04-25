package main

import (
	"log"
	"math"
)

func main(){
	log.Println(math.Pow10(1))
}

func findRestaurant(list1 []string, list2 []string) []string {
	list1Map := make(map[string]int, 0)
	for index, data := range list1 {
		list1Map[data] = index
	}
	res := make([]string , 0)
	diff :=  math.MaxInt64
	for index2 , str2 := range list2{
		if index1 , exist := list1Map[str2] ; exist{
			if index1 + index2 < diff{
				res = make([]string , 0)
				res = append(res , str2 )
				diff = index1 + index2
			}else if index1 + index2 == diff{
				res  = append(res ,str2)
			}
		}
	}
	return res
}

func min(arr ...int) int {
	res := math.MaxInt64
	for _ , data := range arr{
		if data < res{
			res = data
		}
	}
	return res
}