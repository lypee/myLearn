package main

import (
	"context"
	"log"
	"time"
)

type person struct {
	Name string
	Age  int
}

type testMapNil struct {
	List map[int]person
}

func main() {
	ctx := context.Background()
	testWithValue(ctx)
}
func testWithValue(ctx context.Context){
	sCtx1 := context.WithValue(ctx , "key", "value")
	log.Println(sCtx1.Value("key"))
}
func timeTest(){
	timeNow := time.Now()
	tt := int64(1647846990) //2022-2-26
	timeTt := time.Unix(tt, 0)
	log.Println(timeNow.Before(timeTt))

}

func mapNil() {
	res := returnNil()
	log.Println(res.List == nil)
	return
}

func returnNil() testMapNil {
	return testMapNil{}
}
func test() {
	parentCtx := context.Background()
	//ctx1 := context.WithTimeout()
	//ctx , cancel := context.WithCancel(context.Background())
	ctx1, _ := context.WithCancel(parentCtx)
	context.WithValue(ctx1, "123", "456")
	log.Println(parentCtx.Value("123"))
}
