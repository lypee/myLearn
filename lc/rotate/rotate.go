package main

import "log"

// lc 旋转图像
func main() {
	s := []int{}
	log.Println(len(s), " ", cap(s))
	s = append(s, 1)
	log.Println(len(s), " ", cap(s))
	s = append(s, 1)
	log.Println(len(s), " ", cap(s))
	s = append(s, 1)
	log.Println(len(s), " ", cap(s))
	s = append(s, 1, 1)
	log.Println(len(s), " ", cap(s))
}

func rotate(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
	log.Println(matrix)
}
