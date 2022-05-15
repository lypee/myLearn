package main

import (
	"log"
	"time"
)

func main() {
	tmp4()
}

var (
	i int
)

func tmp4() {
	i = 0
	ch := make(chan struct{}, 1)
	go tmp5(ch)
	for idx := 0; idx < 10; idx++ {
		go func() {
			ch <- struct{}{}
		}()
	}
	time.Sleep(time.Second * 3)
	log.Println(i)
}

func tmp5(ch chan struct{}) {
	for {
		select {
		case <-ch:
			i++
		}
	}
}

func test() {
	var x *int
	x = nil
	test1(x)
}

func test1(x interface{}) {
	if x == nil {
		log.Println("nil")
	} else {
		log.Println("not-nil")
	}
}
