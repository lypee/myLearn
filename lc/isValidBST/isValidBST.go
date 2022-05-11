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
	help(root)
	if len(arrs) < 1 {
		return true
	}
	for i := 1 ; i < len(arrs) ; i++{
		if arrs[i] <= arrs[i-1]{
			return false
		}
	}
	return true
}

func help(node *TreeNode){
	if node == nil{
		return
	}
	help(node.Left)
	arrs = append(arrs , node.Val)
	help(node.Right)
	return
}