package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

//
//func height(node *TreeNode) int {
//	if node == nil {
//		return 0
//	}
//	return max(height(node.Left), height(node.Right)) + 1
//}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left)) - abs(height(root.Right)) <= 1 && isBalanced2(root.Right) && isBalanced2(root.Left)
}


func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Right), height(node.Left)) + 1
}
