package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

func GenRandNumberByRange(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	n := start + rand.Intn(end-start)
	return n
}

type Energy struct {
	EnergyType int
	EnergyText string
	Now        time.Time
}

func main() {
	e1 := Energy{
		EnergyType: 2,
		EnergyText: "2",
		Now:        time.Now(),
	}
	e2 := Energy{
		EnergyType: 1,
		EnergyText: "1",
		Now:        time.Unix(1231241231241, 0),
	}

	eList := make([]Energy, 0)
	eList = append(eList, e1, e2)
	var energy Energy
	log.Println(unsafe.Sizeof(e1))
	for i := range eList {
		energy = eList[i]
		log.Println(fmt.Sprintf("%p", &energy))
	}
	return
}

func print1() {
	ch := make(chan int, 1)
	defer close(ch)
	go printNum(ch)
	go printLetter(ch)
	time.Sleep(time.Minute)
}
func printNum(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		if i%2 == 1 {
			log.Println("g1: ", i)
		}
	}
}

func printLetter(ch chan int) {
	for i := 0; i < 100; i++ {
		<-ch
		if i%2 == 0 {
			log.Println("g2: ", i)
		}
	}
}
