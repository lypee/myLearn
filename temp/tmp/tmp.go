package main

import (
	"log"
	"lpynnng/learn/lockMap"
)

type s1 struct {
	i  int
	ii bool
	p  Person
}

type s2 struct {
	p  Person
	ii bool
	i  int
}
type Person struct {
	Age  int
	Name string
	PP   string
}

type T1 struct {
	a [2]int8
	b int64
	c int16
}
type T2 struct {
	a [2]int8
	c int16
	b int64
}

const (
	ii = 3
)
func main() {
	var i int
	log.Println(&i, i)
}

func get2() {
	lockMap.Mcache.Get("2")
}
