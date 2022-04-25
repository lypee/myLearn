package main

import "log"

func main(){
	//arr := [][]int{
	//	[]int{1, 4, 7, 11, 15},
	//	[]int{2, 5, 8, 12, 19},
	//	[]int{3, 6, 9, 16, 22},
	//	[]int{10, 13, 14, 17, 24},
	//	[]int{18, 21, 23, 26, 30},
	//}
	arr := [][]int{
		[]int{-5},
	}
	log.Println(findNumberIn2DArray(arr , -5))
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows < 1 {
		return false
	}
	cols := len(matrix[0])
	if cols < 1 {
		return false
	}
	j := 0
	for i := rows - 1; i >= 0 && j < cols; {
		num := matrix[i][j]
		if target == num {
			return true
		}
		if target > num {
			j++
		} else {
			i--
		}
	}
	return false
}
