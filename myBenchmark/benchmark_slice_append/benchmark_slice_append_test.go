package benchmark_slice_append

import (
	"log"
	"sync"
	"testing"
)
type SafeSlice struct{
	s []int
	lock *sync.RWMutex
}

func Benchmark_append(b *testing.B){
	wg := &sync.WaitGroup{}
	safeSlice := &SafeSlice{
		s:    []int{},
		lock: new(sync.RWMutex),
	}
	nums := 4
	wg.Add(nums)

	testFunc := func(wg *sync.WaitGroup , safeSlice *SafeSlice ,time int) {
		defer wg.Done()
		for i := 0 ; i< time ; i++{
			safeSlice.lock.Lock()
			safeSlice.s = append(safeSlice.s , i)
			safeSlice.lock.Unlock()
		}
	}
	b.ResetTimer()
	for i := 0 ; i< nums ; i++{
		go testFunc(wg  , safeSlice , b.N)
	}
	wg.Wait()

	log.Println("length: ",len(safeSlice.s), " cap: " ,cap(safeSlice.s))
}

func safeAppend(wg *sync.WaitGroup , safeSlice *SafeSlice ,time int){
	defer wg.Done()
	for i := 0 ; i< time ; i++{
		safeSlice.lock.Lock()
		safeSlice.s = append(safeSlice.s , i)
		safeSlice.lock.Unlock()
	}

}
