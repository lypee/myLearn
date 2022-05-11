package main

import (
	"encoding/json"
	"log"
)

func main(){
	s := []int{6,7,8,9,10,11,13,14,15,16,17,18}
	ss  , _:= json.Marshal(&s)
	log.Println(ss)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 0)
	for i := range res {
		res[i] = make([]int, 0)
	}
	if root == nil{
		return res
	}
	stack := make([]*TreeNode ,0)
	stack = append(stack , root)
	for stack != nil{
		length := len(stack)
		subList := make([]int , 0)
		for idx := 0 ; idx < length ; idx++{
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil{
				stack = append(stack , node.Left)
			}
			if node.Right!= nil{
				stack = append(stack , node.Right)
			}
			subList = append(subList ,node.Val)
		}
		res = append(res, subList)
	}
	return res
}
