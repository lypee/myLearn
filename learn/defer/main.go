package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	p1 = 312
)
func init(){
	log.Println("init")
	log.Println(p1)
}
func main() {
	var i int
	i = 90
	defer func() {
		log.Println(1, "i :", i)
	}()
	i = 11
	defer func() {
		log.Println(2, "i :", i)
	}()
	i = 22
	test_defer(i)
	//for i := 0 ; i < 10 ; i++{
	//	test1()
	//	test2()
	//	log.Println(0xf)
	//}
	return

}

func test_defer(i int) {
	i = 333
	defer func() {
		log.Println(3, "i :", i)
	}()
	return
}
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
func test1() {
	curTime := time.Now().Nanosecond()
	log.Println(1 << 20)
	log.Println("test1 ", time.Now().Nanosecond()-curTime)
}

func test2() {
	curTime := time.Now().Nanosecond()
	log.Println(1024 * 1024)
	log.Println("test2 ", time.Now().Nanosecond()-curTime)
}
