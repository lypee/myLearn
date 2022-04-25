package s28


func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return help(root.Left , root.Right)
}

func help(left , right *TreeNode)bool{
	if left == nil && right == nil{
		return true
	}
	if left == nil || right == nil || left.Val != right.Val{
		return false
	}
	return help(left.Left , right.Right) && help(left.Right ,  right.Left)
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}