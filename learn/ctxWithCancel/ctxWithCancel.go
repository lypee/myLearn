package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
	"unsafe"
)

type Person struct {
	Age  int
	Name string
	Avg  string
}

func main() {
	fatherCtx := context.Background()
	sonCtx, cancel := context.WithCancel(fatherCtx)
	ticker := time.NewTicker(time.Second)
	go canc(cancel)
	go watchCM(sonCtx, ticker)
	go NewObject()
	watch(sonCtx, ticker)
	select {
	case <-sonCtx.Done():
		log.Println("done")
	}
	return
}

func NewObject() {
	ticker := time.NewTicker(time.Second)
	log.Println("one大小: " ,unsafe.Sizeof(Person{}))
	for {
		select {
		case <-ticker.C:
			log.Println("分配一次对象")
			pList := make([]Person, 100)
			log.Println("list大小: " ,unsafe.Sizeof(pList))
			pList = append(pList, Person{
				Age:  1,
				Name: "1",
			})
		}
	}
}
func watchCM(ctx context.Context, ticker *time.Ticker) {
	var ms runtime.MemStats

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			runtime.ReadMemStats(&ms)
			log.Println(fmt.Sprintf("总计分配堆内存: %v kb,现分配堆内存: %v kb,  mallocs次数: %v ,goroutine数量: %v ", ms.TotalAlloc/1024, ms.HeapAlloc/1024, ms.Mallocs, runtime.NumGoroutine()))
		}
	}
}
func watch(ctx context.Context, ticker *time.Ticker) {
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("watch-done")
			return
		case <-ticker.C:
			//log.Println("watch-after-1s")
		}
	}

}

func canc(cance context.CancelFunc) {
	time.Sleep(10 * time.Hour)
	cance()
	log.Println("cancel1")
	return
}
