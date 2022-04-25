package main

import (
	"context"
	"log"
	"reflect"
)

type handler func(str string) string

type adder interface {
	add(string) int
}

func (h handler) add(str string) string {
	return str + " 1"
}

func main() {
	arr := []int{1,3,4,5}
	log.Println(arr[:len(arr)])
	//fatherCtx, cance := context.WithCancel(context.Background())
	var a int

	//var my handler = func(name string) string {
	//	go test(fatherCtx)
	//	return name + " name"
	//}
	//
	//log.Println(my.add("123")) // -> param + 闭包函数 + handler.add()

	//data ,ok := a.(int)
	//test(fatherCtx)
	//cancel(cance)
	testInterface(a)
}
func testInterface(i interface{}){
	log.Println(reflect.TypeOf(i).Kind().String())
}
func test(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("ctx.Done")
			return
		default:
			//log.Println("watch...")
		}
	}
	log.Println(1)
	return
}

func cancel(cance context.CancelFunc) {
	cance()
}
