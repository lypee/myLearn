package main

import (
	"log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		go Produce(ch, wg)
	}
	wg.Wait()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Consumer(ch, wg)
	}

	wg.Wait()
	return
}

func Produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		ch <- i
	}
	return
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select{

		}
		if val, ok := <-ch; ok {
			log.Println(val)
		} else {
			return
		}
	}
	return
}
