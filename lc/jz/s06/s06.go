package s06


func reversePrint(head *ListNode) []int {
	if head == nil{
		return []int{}
	}
	var pre *ListNode
	count := 0
	for head != nil{
		count++
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	arr := make([]int ,0 , count)
	for head != nil{
		arr = append(arr , head.Val)
		head = head.Next
	}
	return arr
}


type ListNode struct{
	Val int
	Next *ListNode
}