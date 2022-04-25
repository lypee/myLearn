package main

import (
	"context"
	"log"
	"time"
)

func main() {
	fatherCtx := context.Background()
	sonCtx, cancel := context.WithTimeout(fatherCtx, 5*time.Second)

	ticker := time.NewTicker(time.Second * 1)
	defer cancel()
	for {
		select {
		case <-sonCtx.Done():
			log.Println("接口超时！")
			return
		case <-ticker.C:
			log.Println("ticker.1s")
		//default:
		//	test1()
		//}
		}
	}
	return
}

func test1() {
	log.Println("2")
	return
}
