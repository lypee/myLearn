package main

import "fmt"

func main() {
	ch := make(chan int ,1)
	//ch <- 1
	go producer("生产者1" ,10 , ch)
	go producer("生产者2" , 10 , ch)
	go consumer("消费者1", ch)
	go consumer("消费者2" ,ch)

	select {}
	return
}

func producer(name string, sum int, ch chan int) {
	i := 0
	for {
		i++
		if i > sum {
			break
		}
		ch <- i
		fmt.Println("producer--", name, ":", i)
	}
}

func consumer(name string, ch chan int) {
	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println("consumer--", name, ":", data)
			}
		default:
		}
	}

	return
}
