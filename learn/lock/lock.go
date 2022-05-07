package main

import (
	"log"
	"sync"
	"unsafe"
)

func main(){
	testLock()
}


func testLock(){
	rwLock := sync.RWMutex{}

	rwLock.RLock()

	lock := &sync.Mutex{}
	lock.Lock()

	log.Println(unsafe.Pointer(lock))
}