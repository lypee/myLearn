package main

import (
	"context"
	"log"
	"time"
)
const (
	businessNums = 3
)

func main() {
	// fater
	fatherCtx := context.Background()
	arr := server(fatherCtx)
	log.Println(arr)
}

func server(ctx context.Context) []int {
	// server ctx
	serverCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	ch := make(chan int, businessNums)

	go serverText(serverCtx, ch)
	go serverImg(serverCtx, ch)
	go serverVideo(serverCtx, ch)

	for {
		select {
		case <-serverCtx.Done():
			log.Println("serverCtx.Done()")
			cancel()
			return []int{}
		}
	}
}
//func watch(ctx context.Context){
//	for {
//		select {
//		case <- ticker.C:
//
//		}
//	}
//}
func serverText(ctx context.Context, ch chan int)  {
	textCtx, _ := context.WithTimeout(ctx, time.Second*9)
	text(ch)
	for {
		select {
		case <-textCtx.Done():
			log.Println("serverText.Done")
			return
		case data, _ := <-ch:
			log.Println(data)
			return
		default:
		}
	}
	return
}
func text(ch chan int) {
	// 模拟业务方法
	return
}

func serverImg(ctx context.Context, ch chan int) int {
	return 0
}

func img(ch chan int) {
	return
}

func serverVideo(ctx context.Context, ch chan int) int {
	return 0
}

func video(ctx context.Context) int {
	return 0
}
