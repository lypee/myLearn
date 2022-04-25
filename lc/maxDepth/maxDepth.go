package main

import (
	"log"
	"math"
)

// lc 104 二叉树的最大深度

func main() {
	ss := make([]int, 2)
	ss = append(ss, 1)
	log.Println(ss)
	s := []int{1, 2, 3, 4, 5}
	s1 := make([]int, len(s)+1)
	copy(s1[1:], s)
	s1[0] = 0
	log.Println(s1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return height(root)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
