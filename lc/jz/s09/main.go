package main

import "log"

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	s1 := make([]int , 0)
	s2 := make([]int , 0)
	return CQueue{
		s1: s1,
		s2: s2,
	}
}


func (this *CQueue) AppendTail(value int) {
	if this.s1 == nil{
		this.s1 = make([]int , 0)
	}
	this.s1 = append(this.s1 , value)
	return
}

func (this *CQueue) DeleteHead() int {
	if len(this.s1) < 1 {
		return -1
	}
	this.s2 = make([]int , len(this.s1))
	copy(this.s2 , this.s1)
	this.s1 = make([]int , 0)
	value := this.s2[0]
	for i := 1 ; i < len(this.s2) ; i++{
		this.s1 = append(this.s1 , this.s2[i])
	}
	this.s2 = []int{}
	return value
}




func main(){
	cq := Constructor()
	cq.AppendTail(1)
	cq.AppendTail(2)
	cq.AppendTail(3)
	cq.DeleteHead()
	cq.DeleteHead()
	log.Println(cq)
	return
}