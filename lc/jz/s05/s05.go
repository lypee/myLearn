package main

import "log"

func main(){
	log.Println("")
}
func replaceSpace(s string) string {
	res := make([]byte , 0)
	for i := 0 ; i < len(s) ; i++{
		if s[i] != ' '{
			res = append(res,  s[i])
		}else{
			res = append(res , []byte("%20")...)
		}
	}
	return string(res)
}


func search(arr []int, target int) int {
	left , right := 0 , len(arr) -1
	for left < right{
		if arr[left] == target{
			return left
		}
		mid := (left + right ) / 2
		if arr[mid] == target{
			right = mid
		}else if arr[left] < arr[mid] {
			if arr[left] < target && target < arr[mid]{
				right = mid-1
			}else{
				left = mid +1
			}
		}else if arr[left] > arr[mid]{
			if arr[right] > target && target > arr[mid]{
				left = mid + 1
			}else{
				right = mid - 1
			}
		}else{
			left++
		}
	}
	return -1
}
