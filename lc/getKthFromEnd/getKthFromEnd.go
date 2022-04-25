package main

type ListNode struct{
	Val int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil{
		return nil
	}
	curHead := head
	nextNode := head
	for i := 0 ; i < k ; i++{
		curHead = curHead.Next
	}
	for curHead != nil{
		curHead = curHead.Next
		nextNode = nextNode.Next
	}
	return nextNode
}
