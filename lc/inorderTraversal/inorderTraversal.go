package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res []int
)

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	help(root)
	return res
}

func help(root *TreeNode) {
	if root == nil {
		return
	}
	help(root.Left)
	res = append(res, root.Val)
	help(root.Right)
}
