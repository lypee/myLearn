package goroutine_pool

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	ctx := context.Background()
	pool := NewPool(ctx , 4)
	for i := 0; i < 9; i++ {
		pool.AddTask(Print1)
		//if i == 2{
		//	pool.Stop()
		//}
	}
	//pool.AddTask(Print2)
	time.Sleep(time.Second)
	pool.Stop()
	pool.AddTask(Print2)
	//pool.AddTask(Print2)
	time.Sleep(time.Minute)
}

func Print1() {
	time.Sleep(time.Second * 2)
	log.Println("1")
}

func Print2(){
	time.Sleep(time.Second * 3)
	log.Println("2")
}