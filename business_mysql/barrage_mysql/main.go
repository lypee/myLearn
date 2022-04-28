package main

import (
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
}

func main(){
	for i := 0 ; i < 128 ; i++{
		fmt.Println(fmt.Sprintf("ALTER TABLE xes_barrage_message_log_%d ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)';",i))
	}
	for i := 0 ; i < 128 ; i++{
		fmt.Println(fmt.Sprintf("ALTER TABLE xes_barrage_stu_speak_log_%d ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)';",i))
	}
	for i := 0 ; i < 128 ; i++{
		fmt.Println(fmt.Sprintf("ALTER TABLE xes_leave_message_log_%d ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)';",i))
	}
	fmt.Println(fmt.Sprintf("ALTER TABLE xes_barrage_send_log ADD COLUMN `businessline_id` int(11) NOT NULL DEFAULT 0 COMMENT '业务(10:素质app 20:高中app 30:新学科app 40:中老年项目)';"))

}

func test(head *ListNode){
	pre := new(ListNode)
	cur := head
	for head != nil{
		next := head.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
