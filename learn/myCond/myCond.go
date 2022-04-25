package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func main(){
	cond := new(sync.Cond)
	cond.L=  new(sync.Mutex)
	for i := 0 ; i < 10 ; i++{
		name := strconv.Itoa(i)
		go DoSomething(cond , name)
	}
	time.Sleep(2 *time.Second)
	log.Println("Broadcast")
	cond.Broadcast()
	time.Sleep(time.Minute)

}

func DoSomething(c *sync.Cond,name string){
	c.L.Lock()
	defer c.L.Unlock()
	log.Println(name + " 开始执行")
	c.Wait()
	log.Println(name + " wait结束")
	time.Sleep(time.Second)
}