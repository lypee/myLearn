package main

import (
	"log"
	"sync"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift
	mutexWaiterShift2
)

var (
	wg *sync.WaitGroup
)

func main() {
	log.Println(1 << 16)
	return
}
