package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return help(root , targetSum)
}

func help(node *TreeNode , diff int) bool{
	if node == nil{
		return false
	}
	if node.Left == nil && node.Right == nil{
		return node.Val == diff
	}
	return help(node.Left , diff - node.Val) || help(node.Right , diff - node.Val)
}
