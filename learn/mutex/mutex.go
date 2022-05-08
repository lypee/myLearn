package main

import (
	"log"
)
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken // 2
	mutexStarving  //4
	mutexWaiterShift = iota
	starvationThresholdNs = 1e6
)
func main() {

}


func mockPanic(){

}

func test1() {
	old := mutexLocked
	log.Printf("mutex占用内存大小:[%v]",old&(mutexLocked|mutexStarving) == mutexLocked)
	old = mutexWoken
	log.Printf("mutex占用内存大小:[%v]",old&(mutexLocked|mutexStarving) == mutexLocked)
	old = mutexStarving
	log.Printf("mutex占用内存大小:[%v]",old&(mutexLocked|mutexStarving) == mutexLocked)

}
