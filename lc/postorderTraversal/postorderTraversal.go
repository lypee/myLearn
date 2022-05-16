package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res []int
)

func postorderTraversal(root *TreeNode) []int {
	res = []int{}
	help(root)
	return res
}

func help(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		help(node.Left)
	}
	if node.Right != nil {
		help(node.Right)
	}
	res = append(res, node.Val)
	return
}
