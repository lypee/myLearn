package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	node1 ,node2 := headA , headB
	for node1 != node2{
		if node1 != nil{
			node1 = node1.Next
		}else{
			node1 = headB
		}
		if node2 != nil{
			node2 = node2.Next
		}else{
			node2 = headA
		}
	}
	return node1
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA , curB := headA , headB
	for curA != curB{
		if curA == nil{
			curA = headB
		}else{
			curA = curA.Next
		}

		if curB == nil{
			curB = headA
		}else{
			curB = curB.Next
		}
	}
	return curA
}

