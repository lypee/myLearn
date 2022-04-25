package main

import (
	"log"
	"sync"
)

func main() {
	ch1, ch2, ch3 := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go print1(ch1, ch2, wg)
	go print2(ch2, ch3, wg)
	go print3(ch3, ch1, wg)
	ch1 <- struct{}{}
	wg.Wait()
	log.Println("finish")
}

func print1(ch1, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ch1:
			log.Println("1")
			ch2 <- struct{}{}
			if i > 99 {
				return
			}
			i++
		}
	}
}

func print2(ch2, ch3 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ch2:
			log.Println("2")
			ch3 <- struct{}{}
			if i > 99 {
				return
			}
			i++
		}
	}
}

func print3(ch3, ch1 chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-ch3:
			log.Println("3")
			ch1 <- struct{}{}
			if i > 99 {
				return
			}
			i++
		}
	}
}
