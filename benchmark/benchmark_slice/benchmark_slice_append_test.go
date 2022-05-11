package benchmark_slice_append

import (
	"sync"
	"testing"
)


type UnSafeSlice struct {
	s    []int
}

type SafeSlice struct {
	s    []int
	lock *sync.RWMutex
}

var (
	nums  = 1000 // 协程数
	times = 1000 //循环次数
)

func Benchmark_AppendWithLock(b *testing.B) {
	wg := &sync.WaitGroup{}
	safeSlice := &SafeSlice{
		s:    []int{},
		lock: new(sync.RWMutex),
	}
	wg.Add(nums)

	testFunc := func(wg *sync.WaitGroup, safeSlice *SafeSlice, time int) {
		defer wg.Done()
		for i := 0; i < time; i++ {
			safeSlice.lock.Lock()
			safeSlice.s = append(safeSlice.s, i)
			safeSlice.lock.Unlock()
		}
	}
	b.ResetTimer()
	for i := 0; i < nums; i++ {
		go testFunc(wg, safeSlice, times)
	}
	wg.Wait()

	//log.Println("length: ",len(safeSlice.s), " cap: " ,cap(safeSlice.s))
}

func Benchmark_AppendNoLock(b *testing.B) {
	unsafeSlice := UnSafeSlice{s: []int{}}

	b.ResetTimer()
	for i := 0 ; i < times * nums ; i++{
		unsafeSlice.s = append(unsafeSlice.s , i)
	}
}

func safeAppend(wg *sync.WaitGroup, safeSlice *SafeSlice, time int) {
	defer wg.Done()
	for i := 0; i < time; i++ {
		safeSlice.lock.Lock()
		safeSlice.s = append(safeSlice.s, i)
		safeSlice.lock.Unlock()
	}

}
