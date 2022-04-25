package main

import (
	"log"
	"time"
)

func main() {
	ch1 := make(chan int,999)
	//ch1 <- 1
	//close(ch1)

	//ch2 := make(chan int,1 )
	log.Println(<-ch1)
	log.Println(<-ch1)
	//log.Println(<-ch1)
	//log.Println(<-ch2)
	//push_data()
	return
}
func push_data() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	go push_data1(ch1)
	for{
		select {
		case data, _ := <-ch1:
			log.Println(data)
			return
		case data, _ := <-ch2:
			log.Println(data)
		case data, _ := <-ch3:
			log.Println(data)
		default:
			//log.Println("999")
		}
	}
	return
}
func push_data1(ch1 chan int) {
	time.Sleep(time.Second)
	ch1 <- 1
	return
}

func push_data2(ch2 chan int) {
	ch2 <- 1
	return
}

func push_data3(ch3 chan int) {
	ch3 <- 1
	return
}
