package main

import (
	"errors"
	"log"
	"sort"
)



type ListNode struct{
	Val int
	Next *ListNode
}
//

type Person struct{
	Age int
}
func main() {
	p1 := Person{
		Age: 1,
	}
	p2 := Person{
		Age: 2,
	}
	p3 := Person{
		Age: 3,
	}
	pList := []Person{p1, p2, p3}
	sort.Slice(pList , func(i, j int) bool {
		return pList[i].Age > pList[j].Age
	})
	log.Println(pList)
}


func mergeTwoLists( l1 *ListNode ,  l2 *ListNode ) *ListNode {
	// write code here
	//
	//1.  head
	var  head  ListNode
	cur := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			cur.Val = l1.Val
			l1 = l1.Next
		}else{
			cur.Val = l2.Val
			l2 = l2.Next
		}
		cur.Next = new(ListNode)
		cur = cur.Next


	}
	if l1 != nil{
		cur.Next = l1
	}
	if l2 != nil{
		cur.Next = l2
	}
	return head.Next
	//
	//
}
// []int{9,11,13,15,1,2, 5} target(7)    ,not exist:-1 ,invalid input:-2
func findTarget(arrs []int, target int) (int ,error) {
	length := len(arrs)
	if length < 1 {
		return -1,errors.New("invalid input")//输入不合法
	}
	left, right := 0, length-1
	for left < right {
		mid := (left + right) / 2
		num := arrs[mid]
		if num == target {
			return mid ,nil
		}
		if num > target{
			if num < arrs[right]{
				right = mid +1
			}else{
				left  = mid - 1
			}
		}else{
			// num < target
			if num < arrs[right]{
				right = mid +1
			}else{
				left = mid - 1
			}
		}
	}
	return -1 , nil
}
