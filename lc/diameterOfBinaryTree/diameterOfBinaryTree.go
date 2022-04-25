package diameterOfBinaryTree

import "math"

// lc 二叉树的直径

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	result int
)

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	help(root)
	return result
}

func help(node *TreeNode) int{
	if node == nil{
		return 0
	}
	left := help(node.Left)
	right := help(node.Right)
	result = max(result, left + right)
	return max(left , right) + 1
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