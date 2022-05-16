package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	tmp6()
}

var (
	i int
)

func tmp6(){
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

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
