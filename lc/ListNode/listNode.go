package ListNode

type ListNode struct{
	Val int
	Next *ListNode
}


func reverseListNode(head *ListNode) *ListNode{
	if head == nil{
		return nil
	}
	var pre *ListNode
	count := 0
	for head != nil{
		count++
		node := head.Next
		head.Next = pre
		pre = head
		head = node
	}
	return pre
}