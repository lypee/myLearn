package reversePrint

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head != nil {
		next := head.Next

		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reversePrint(head *ListNode) []int {
	tailHead := reverse(head)
	arrs := make([]int, 0)
	for tailHead != nil {
		arrs = append(arrs, head.Val)
		tailHead = head.Next
	}
	return arrs
}
