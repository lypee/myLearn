package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)


var writeFlag = 0


func main() {
	//var f1,f2 int
	//f1 , f2 = 1 , 1
	//f1 &^= writeFlag
	a := 1 &^ 1
	b := 0 &^ 0
	c := 1 &^ 0
	d := 0 &^ 1

	e := 26 &^ 5

	f := ^5
	g := 26 & f
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(d)
	log.Println(e)
	log.Println(f)
	log.Println(g)

	//for i := 0 ; i<= 999 ; i++{
	//	log.Println(i &^ 5)
	//}

}




func consumer(cname string, ch chan int) {

	//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。
LOOP:
	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Println("consumer--------", cname, ":", i)
			} else {
				break LOOP
			}

		}

	}
	fmt.Println(cname + " ch closed.")
	return

}

func producer(pname string, maxCount int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < maxCount; i++ {
		fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
	return
}


func start() {
	ch := make(chan int)
	Nums := 10
	MaxCount := 99999
	for i := 0; i < Nums; i++ {
		name := strconv.Itoa(i)
		go consumer("消费者"+name, ch)
	}
	wg := &sync.WaitGroup{}
	wg.Add(Nums)
	for i := 0; i < Nums*10; i++ {
		name := strconv.Itoa(i)
		go producer("生产者"+name, MaxCount, ch, wg)
	}
	wg.Wait()
	close(ch)
	log.Println("all-done")
	return
}

func t1(){
	mmap := make(map[int]int , 0)
	mmap[1] =1
}

//for i := range ch {
//fmt.Println("consumer-----------", cname, ":", i)
//
//}
//
//// 这个也可以改成：LOOP:
//for {
//select {
//case i,ok:=<-ch:
//if ok {
//fmt.Println("consumer--------", cname, ":", i)
//} else {
//break LOOP
//}
//
//}
//
//}

//注意： i := <- ch  从空的channel中读取数据不会panic, i读取到的值是0,  如果channel是bool的，那么读取到的是false

//判断channel是否关闭，可以使用像上面的ok pattern
