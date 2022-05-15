package main

import (
	"log"
	"unsafe"
)

func main() {
	ch := make(chan int)
	log.Println(unsafe.Pointer(&ch))
	//go tPop(ch)
	//for i := 0; i < 10; i++ {
	//	go tPush(ch, i)
	//}
	//select {}
}

func tPush(ch chan int, idx int) {
	for {
		ch <- idx
	}
}

func tPop(ch chan int) {
	for {
		select {
		case data, _ := <-ch:

			log.Println("recv-data: ", data)
		default:

		}
	}
}
