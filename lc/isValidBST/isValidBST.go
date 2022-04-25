package isValidBST
type TreeNode struct {
	    Val int
	     Left *TreeNode
	     Right *TreeNode
	 }


var (
	arrs []int
)

func isValidBST(root *TreeNode) bool {

	arrs = []int{}
	arrs = help(root)
	if len(arrs) <= 1 {
		return true
	}
	for i := 0 ; i < len(arrs) - 1 ; i++{
		if arrs[i]  >= arrs[i+1]{
			return false
		}
	}
	return true
}

func help(root *TreeNode) []int{
	if root == nil{
		return arrs
	}
	help(root.Left)
	arrs = append(arrs , root.Val)
	help(root.Right)
	return arrs
}