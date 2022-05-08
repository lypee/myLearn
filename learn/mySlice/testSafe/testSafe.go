package main

import (
	"log"
	"sync"
)

type SafeSlice struct {
	s    []int
	lock *sync.RWMutex
}

func main() {
	TestWithLock()
}

func TestWithLock() {
	nums := 100
	wg := &sync.WaitGroup{}
	safeSlice := SafeSlice{
		s:    []int{},
		lock: new(sync.RWMutex),
	}
	i := 0
	for idx := 0; idx < nums; idx++ { // 并行nums个协程做append
		wg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {

				}
				safeSlice.lock.Unlock()
				wg.Done()
			}()

			safeSlice.lock.Lock()
			safeSlice.s = append(safeSlice.s, i)
			if i == 98{
				panic("123")
			}
			i++
		}()
	}
	wg.Wait()
	log.Println(len(safeSlice.s))
}
