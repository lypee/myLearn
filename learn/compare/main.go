package main

import (
	"log"
	"runtime"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	runtime.GOMAXPROCS(1)
	num := 300
	wg := &sync.WaitGroup{}
	wg.Add(num)
	for index := 0 ; index < num; index++{
		i := index
		go func() {
			defer wg.Done()
			log.Println(i)
			return
		}()
	}
	wg.Wait()
	return
}
func Level1(p Person){
	p.Name = "456"
	log.Println(p)
	return
}
func Test1(){
	return
}

func DigitNum() {
	s := "1011110110111010100011001011110011111110010100000000001001"
	log.Println(len(s))
}