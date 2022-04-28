package benchmark_lock

import (
	"sync"
	"testing"
)

var (
	nums = 2 << 22
)

func BenchmarkRWMutexWithRLock(b *testing.B) {
	lock := sync.RWMutex{}
	b.ResetTimer()
	for i := 0; i < nums; i++ {
		go func() {
			lock.RLock()
			defer lock.RUnlock()
		}()
	}
}

func BenchmarkRWMutexWithLock(b *testing.B) {
	lock := sync.RWMutex{}
	b.ResetTimer()
	for i := 0; i < nums; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
		}()
	}
}

func BenchmarkMutex(b *testing.B) {
	lock := sync.Mutex{}
	b.ResetTimer()
	for i := 0; i < nums; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
		}()
	}
}


