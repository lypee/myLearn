package main

import (
	"log"
	"time"
)

type Person struct {
	Name string `json:"name" gson:"123"`
	Age  int    `json:"age"`
}

func main() {
	//var s []string
	//s = append(s , "1")
	////log.Println(s)
	////log.Println(1)
	//t := Person{}
	//log.Println(reflect.TypeOf(t).Field(0).Tag.Get("gson"))
	//for  i := 0 ; i < reflect.TypeOf(t).NumField() ; i++{
	//	log.Println(reflect.TypeOf(t).Field(i).Tag.Get("json"))
	//}
	//ch := make(chan int, 1)
	go keepAliveV2()
	//time.Sleep(time.Minute)
	select {}
}
func keepAliveV2() {
	time.Sleep(time.Second * 5)
	log.Println("alive")
	return
}

func keepAlive(ch chan int) {
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
