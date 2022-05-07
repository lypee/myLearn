package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	curr, carry := head, 0
	x, y := 0, 0
	for l1 != nil || l2 != nil {
		x, y = 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		curr = curr.Next
		carry = total / 10
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head != nil {
		node := head.Next
		head.Next = pre
		pre = head
		head = node
	}
	return pre
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}
	curr, carry := head, 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil{
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			y = l2.Val
			l2 = l2.Next
		}
		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		carry = total / 10
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head
}
