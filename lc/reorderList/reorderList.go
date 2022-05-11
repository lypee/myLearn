package reorderList

import (
	"log"
	"unsafe"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	copyHead := copyList(head)
	log.Println(unsafe.Pointer(copyHead))

}
func copyList(head *ListNode) *ListNode{
	cur := head
	back := cur
	for head != nil{
		cur.Next = &ListNode{
			Val:  head.Val,
		}
		cur = cur.Next
		head = head.Next
	}
	return back.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
