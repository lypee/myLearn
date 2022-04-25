package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val{
		return root
	}
	left := lowestCommonAncestor(root.Left  ,p , q)
	right := lowestCommonAncestor(root.Right , p , q)
	if left != nil && right != nil{
		return root
	}
	if left != nil{
		return left
	}
	return right
}
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode{
	if root == nil{
		return root
	}
	if root.Val == p.Val || root.Val == q.Val{
		return root
	}
	left := lowestCommonAncestor2(root.Left , p , q)
	right := lowestCommonAncestor2(root.Right , p , q)
	if left != nil && right != nil{
		return root
	}
	if left != nil{
		return left
	}
	if right != nil{
		return right
	}
	return nil
}