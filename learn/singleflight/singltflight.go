package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type Person struct {
	Name string
	Age  int
}


var (
	sfKey1 = "key1"
	wg     *sync.WaitGroup
	sf     singleflight.Group
	nums   = 10
)


func main() {
	getValueBySingleflight(sfKey1)
}

func sfWithOutChan(key string) int {
	wg = &sync.WaitGroup{}
	return sfWithOutChanRun()
}

func sfWithOutChanRun() int {
	var num int
	wg.Add(nums)
	for i := 0; i < nums; i++ {
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
	wg.Add(nums)
	done := make(chan bool, 10)
	personChan := make(chan Person, 10)
	//ch := make(chan singleflight.Result,1)
	for i := 0; i < nums; i++ {
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

func singleFlightWithChan(key string) {
}

func getValueBySingleflight(key string) {
	var val string
	wg = &sync.WaitGroup{}
	wg.Add(nums)
	for idx := 0; idx < nums; idx++ { //模拟多个协程同时请求
		go func(idx int) {
			defer wg.Done()
			value, _ := getAndSetCacheNoChan(idx, key) //简化代码，不处理error
			log.Printf("request %v get value: %v", idx, value)
			val = value
		}(idx)
	}
	wg.Wait()
	log.Println("val: ", val)
	return
}

func getAndSetCacheNoChan(idx int, cacheKey string) (string, error) {
	log.Printf("idx %v into-cache...", idx)
	value, _, _ := sf.Do(cacheKey, func() (ret interface{}, err error) {
		log.Printf("idx %v is-setting-cache", idx)
		// 休眠0.1s以捕获并发的相同请求
		time.Sleep(100 * time.Millisecond)
		log.Printf("idx %v set-cache-success!", idx)
		return "myValue", nil
	})
	return value.(string), nil
}

func enterGetAndSetCacheWithChan(ctx context.Context, key string) (str string, err error) {
	tag := "enterGetAndSetCacheWithChan"
	sonCtx, _ := context.WithTimeout(ctx, 2*time.Second)
	val := ""
	nums := 10 //协程数
	wg = &sync.WaitGroup{}
	wg.Add(nums)
	for idx := 0; idx < nums; idx++ {
		go func() {
			defer wg.Done()
			val, err = getAndSetCacheWithChan(sonCtx, idx, key)
			if err != nil {
				log.Printf("err:[%+v]", err)
				return
			}
			str = val
		}()
	}
	wg.Wait()
	log.Printf("tag:[%s] val:[%s]", tag, val)
	return
}

func getAndSetCacheWithChan(ctx context.Context, idx int, cacheKey string) (string, error) {
	tag := "getAndSetCacheWithChan"
	log.Printf("tag: %s ;idx %d into-cache...", tag, idx)
	ch := sf.DoChan(cacheKey, func() (ret interface{}, err error) { //do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB
		log.Printf("idx %v is-setting-cache", idx)
		time.Sleep(100 * time.Millisecond)
		log.Printf("idx %v set-cache-success!", idx)
		return "myValue", nil
	})
	for {
		select {
		case <-ctx.Done():
			return "", errors.New("ctx-timeout")
		case data, _ := <-ch:
			return data.Val.(string), nil
		default:
		}
	}
}

type (
	Group struct {
		mu sync.Mutex       // 互斥锁
		m  map[string]*call // lazily initialized
	}

	call struct {
		wg sync.WaitGroup
		// 存储 调用singleflight.Do()方法返回的结果
		val interface{}
		err error

		// 软删除-调用singleflight.Forget(key)并不会将key从map中删除，仅仅做标记。
		forgotten bool

		// 通俗的理解成singleflight合并的并发请求数
		dups  int
		// 存储 调用singleflight.DoChan()方法返回的结果
		chans []chan<- Result
	}

	Result struct {
		Val    interface{}
		Err    error
		Shared bool
	}
)


//func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
//	g.mu.Lock() // 加互斥锁
//	if g.m == nil { // 懒加载map
//		g.m = make(map[string]*call)
//	}
//	if c, ok := g.m[key]; ok { // 检查是否有相同的请求
//		c.dups++
//		g.mu.Unlock()
//		c.wg.Wait() // 如果有，调用waitGroup的wait()方法阻塞住本次调用
//
//		if e, ok := c.err.(*panicError); ok { //如果调用完成，发生error ，将error上抛
//			panic(e)
//		} else if c.err == errGoexit {
//			runtime.Goexit()
//		}
//		return c.val, c.err, true
//	}
//	c := new(call) // 相同的请求第一次进入singleflight
//	c.wg.Add(1)
//	g.m[key] = c // new一个call实体，放其singleflight.call这个map
//	g.mu.Unlock()
//
//	g.doCall(c, key, fn) //实际执行的函数
//	return c.val, c.err, c.dups > 0
//}
//
//
//func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result {
//	ch := make(chan Result, 1)
//	g.mu.Lock()
//	if g.m == nil {
//		g.m = make(map[string]*call)
//	}
//	if c, ok := g.m[key]; ok {
//		c.dups++
//		c.chans = append(c.chans, ch)
//		g.mu.Unlock()
//		return ch
//	}
//	c := &call{chans: []chan<- Result{ch}}
//	c.wg.Add(1)
//	g.m[key] = c
//	g.mu.Unlock()
//
//	go g.doCall(c, key, fn)
//
//	return ch
//}
//
//
//
//// doCall handles the single call for a key.
//func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {
//	normalReturn := false
//	recovered := false
//
//	defer func() { //使用双重defer来区分error panic和runtime.Goexit类型的错误
//		if !normalReturn && !recovered { //
//			c.err = errGoexit
//		}
//
//		c.wg.Done()
//		g.mu.Lock()
//		defer g.mu.Unlock()
//		if !c.forgotten { // 检查key是否调用了Forget()
//			delete(g.m, key)
//		}
//
//		if e, ok := c.err.(*panicError); ok {
//			// 防止channel被永久阻塞
//			if len(c.chans) > 0 {
//				go panic(e)  // panic无法被恢复
//				select {} // 阻塞本goroutinue.
//			} else {
//				panic(e)
//			}
//		} else if c.err == errGoexit {
//
//		} else {
//			// 将结果正常地返回
//			for _, ch := range c.chans {
//				ch <- Result{c.val, c.err, c.dups > 0}
//			}
//		}
//	}()
//
//	func() {
//		defer func() {
//			if !normalReturn {
//				// 如果非正常逻辑执行返回
//				// 此时与panic相关的堆栈已经被丢弃(调用的fn()) ，无法通过堆栈跟踪去确定error类型
//				if r := recover(); r != nil {
//					c.err = newPanicError(r)
//				}
//			}
//		}()
//
//		c.val, c.err = fn() // 执行我们实际的业务逻辑，并将业务方法的返回值赋给singleflight.call的val和err属性
//		// 如果fn()发生panic，normalReturn无法被赋值为true，而是进入上面的defer()
//		normalReturn = true
//	}()
//	// 检查fn()执行时是否发生了panic，仅当未发生panic()时 ,normalReturn才能被置为true
//	if !normalReturn {
//		recovered = true // recovered标志位置为true
//	}
//}
