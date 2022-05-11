package main

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	s1 := make([]int, 0) // use
	s2 := make([]int, 0) // tmp
	return CQueue{
		s1: s1,
		s2: s2,
	}
}

func (this *CQueue) AppendTail(value int) {
	if len(this.s1) < 1 {
		return
	}
	this.s1 = append(this.s1, value)
	return
}

func (this *CQueue) DeleteHead() int {
	len1 := len(this.s1)
	if len1 < 1 {
		return -1
	}
	value := this.s1[0]
	this.s1 = this.s1[1:]
	return value
}
