package main

import (
	"log"
	"time"
)

func main() {
	//ch := make(chan int)
	////log.Println(<- ch)
	//go write(ch)
	//go read(ch)
	//time.Sleep(time.Second * 2)
	//close(ch)
	//time.Sleep(time.Second * 5
	//select {}

	forRange()
}
func write(ch chan int, times int) {
	for i := 0; i < times; i++ {
		ch <- i
		log.Printf("send: [%d]", i)
		break
	}
}
func read(ch chan int) {
	for {
		data, ok := <-ch // 阻塞的
		if ok {
			log.Printf("recv: [%d]", data)
		} else {
			log.Println("channel close ")
			break
		}
	}
}

func readV2(ch chan int) {
	for {
		select {
		case data, ok := <-ch:
			if ok {
				log.Printf("recv:[%v]", data)
			} else {
				log.Printf("ok-false")
			}
		default:
			log.Println("into-default")
		}
	}
}
func consumer() {

}
func chn() {
	ch := make(chan int)
	go write(ch , 1 )

	//time.Sleep(2  * time.Second)

	close(ch)
	write(ch ,1 )
	//log.Println(<- ch)

	time.Sleep(time.Minute)
}

func forRange() {
	ch := make(chan int, 1)
	go readV2(ch)
	go write(ch ,1 )

	time.Sleep(time.Second)
	log.Println("休眠1s")

	go write(ch ,1 )
	time.Sleep(time.Minute)
}


func onlyReadAndWrite(){
	ch := make(chan int ,1)
	onlyRead(ch)
	onlyWrite(ch)
}

// 参数为只读channel
func onlyRead(ch <-chan int ){
	<- ch
}

// 参数为只写channel
func onlyWrite(ch chan<- int ){
	 ch <- 1
}