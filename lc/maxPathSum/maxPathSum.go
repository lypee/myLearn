package maxPathSum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	maxSum int
)

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt64
	maxGain(root)
	return maxSum
}

//func maxGain(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//	leftGain := max(maxGain(node.Left), 0)
//	rightGain := max(maxGain(node.Right), 0)
//
//	maxSum = max(maxSum, leftGain+rightGain+node.Val)
//	return node.Val + max(leftGain, rightGain)
//}

func max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}


func maxPathSum2(root *TreeNode) int {
	maxSum = math.MinInt64
	maxGain(root)
	return maxSum
}


func maxGain(node *TreeNode) int {
	if node == nil{
		return 0
	}
	leftGain := max(maxGain(node.Left) , 0)
	rightGain := max(maxGain(node.Right) , 0)

	maxSum = max(maxSum , node.Val + leftGain + rightGain)

	return node.Val + max(leftGain , rightGain)


}
