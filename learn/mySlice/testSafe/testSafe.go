package main

import (
	"log"
	"sync"
)

type SafeSlice struct{
	s []int
	lock *sync.RWMutex
}

func main(){
	TestWithLock()
}

func TestWithLock(){
	nums := 100
	wg := &sync.WaitGroup{}
	safeSlice := SafeSlice{
		s:    []int{},
		lock: new(sync.RWMutex),
	}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0 ; i< nums ; i++{
			safeSlice.lock.Lock()
			safeSlice.s = append(safeSlice.s , i)
			safeSlice.lock.Unlock()
		}
	}()


	for i := 0 ; i < nums ; i++{ // 并行nums个协程做append
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			for idx := 0 ; idx < nums ; idx++{
				safeSlice.lock.Lock()
				safeSlice.s = append(safeSlice.s , idx)
				safeSlice.lock.Unlock()
			}
		}()
	}
	wg.Wait()
	log.Println(len(safeSlice.s))
}


