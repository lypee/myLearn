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
	nums := 10000
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
	wg.Wait()
	log.Println(len(safeSlice.s))
}


