package main

import (
	"log"
	"sync"
	"time"
)
var nums int
func main(){
	test()
}

func test(){
	sc := sync.NewCond(&sync.Mutex{})
	for i := 0 ;  i < 1; i++{
		go test2(sc)
	}
	time.Sleep(4 * time.Second)
	nums = 11
	return
}

func test2(sc *sync.Cond){
	sc.L.Lock()
	defer sc.L.Unlock()
	for{
		if nums >= 10{
			time.Sleep(1 * time.Second)
			break
		}else{
			log.Println("waiting....")
		}
	}
	log.Println("listen")
	return
}
