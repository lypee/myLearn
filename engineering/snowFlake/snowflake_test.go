package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func BenchmarkSnowflake(b *testing.B) {
	w := NewWorker(5, 5)
	// 64bit = 8b
	ctx := context.Background()

	sonCtx , cancel := context.WithCancel(ctx)
	count := 40000
	ch := make(chan uint64, count+1)
	defer close(ch)

	go countMap(sonCtx , ch)
	nums := 100000
	wg.Add(nums)
	b.ResetTimer()
	//并发 count个goroutine 进行 snowFlake ID 生成
	for i := 0; i < nums; i++ {
		go func() {
			defer wg.Done()
			id, _ := w.NextID()
			ch <- id
		}()
	}
	wg.Wait()
	cancel()
	time.Sleep(1 * time.Second)
	return
}

func countMap(ctx context.Context, ch chan uint64) {
	syncMap := make(map[uint64]int  , 1)
	dupNum := 0
	unDupNum := 0
	for {
		select {
		case <-ctx.Done():
			log.Println(dupNum)
			log.Println(unDupNum)
			log.Println("done")
			return
		case data, _ := <-ch:
			if _, ok := syncMap[data]; ok {
				dupNum++
			} else {
				unDupNum++
			}
			syncMap[data] = 1
		default:

		}

	}

}
