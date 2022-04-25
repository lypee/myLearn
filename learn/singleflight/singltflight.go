package main

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	sfKey1 = "key1"
	wg     *sync.WaitGroup
	sf     singleflight.Group
	times  = 10
)

type Person struct {
	Name string
	Age  int
}

//func main() {
//	sfWithOutChanRun()
//}
func sfWithOutChan(key string) int {
	wg = &sync.WaitGroup{}
	return sfWithOutChanRun()
}

func sfWithOutChanRun() int {
	var num int
	wg.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			res, _, _ := sf.Do(sfKey1, func() (interface{}, error) {
				r := run()
				return r, nil
			})
			if data, ok := res.(int); !ok {
				num = -1
			} else {
				num = data
			}
		}()
	}
	wg.Wait()
	return num
}

func run() int {
	time.Sleep(5 * time.Second)
	res := rand.Intn(1000)
	log.Println("调用一次")
	return res
}

func sfWithChan() {
	wg = &sync.WaitGroup{}
	sfWithChanRun()
}

func sfWithChanRun() {
	wg.Add(times)
	done := make(chan bool, 10)
	personChan := make(chan Person, 10)
	//ch := make(chan singleflight.Result,1)
	for i := 0; i < times; i++ {
		go func() {
			defer wg.Done()
			sf.DoChan(sfKey1, func() (interface{}, error) {
				time.Sleep(5 * time.Second)
				personChan <- Person{
					Name: "lpy",
					Age:  18,
				}
				done <- true
				return "789", nil
			})
		}()
	}
	wg.Wait()
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-done:
			person := <-personChan
			log.Println(person.Age, "  ", person.Name)
		case <-ticker.C:
			log.Printf("456")
		}
	}
	return
}
func sfWait5s() (interface{}, error) {
	ch := make(chan bool)
	ch <- true
	return ch, nil
}

func main() {

	var singleSetCache singleflight.Group

	getAndSetCache := func(requestID int, cacheKey string) (string, error) {
		log.Printf("request %v start to get and set cache...", requestID)
		value, _, _ := singleSetCache.Do(cacheKey, func() (ret interface{}, err error) { //do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB
			log.Printf("request %v is setting cache...", requestID)
			time.Sleep(2 * time.Second)
			log.Printf("request %v set cache success!", requestID)
			return "VALUE", nil
		})
		return value.(string), nil
	}

	cacheKey := "cacheKey"
	var v string
	for i := 1; i < 10; i++ { //模拟多个协程同时请求
		go func(requestID int) {
			value, _ := getAndSetCache(requestID, cacheKey)
			log.Printf("request %v get value: %v", requestID, value)
			v = value
		}(i)
	}
	time.Sleep(3 * time.Second)
	log.Println("v: " , v)

	return
}
