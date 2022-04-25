package main

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

type Part2 struct {
	d int64
	b int32
	c int8
	a bool
	e byte
	//p1 []Part1
}

func main() {
	l1 := []int{1,2,3}
	//l2 := []int{4,5,6}
	copy(l1[1:],  l1)
	log.Println(l1)
	test()
	test2()
}
func test() {
	fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	fmt.Printf("int size: %d\n", unsafe.Sizeof(reflect.Int))
	fmt.Printf("UnsafePointer size: %d\n", unsafe.Sizeof(reflect.UnsafePointer))
	fmt.Printf("map size: %d\n", unsafe.Sizeof(reflect.Map))

	//fmt.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))
}

func test2() {
	p1 := Part1{
		a: false,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
	}
	p12 := Part1{
		a: false,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
	}
	mmap1 := make(map[int]Part1 ,0)
	mmap2 := make(map[int]Part2 ,0)
	mmap1[1] = p1
	mmap1[2] = p12
	log.Println("1")
	log.Println(unsafe.Sizeof(mmap1), " ")
	log.Println(unsafe.Sizeof(mmap2), " ")


	log.Println(unsafe.Sizeof(p1), " ", unsafe.Alignof(p1))
	p2 := Part2{
		d: 0,
		b: 0,
		c: 0,
		a: false,
		e: 0,
		//p1: []Part1{p1, p12},
	}
	log.Println(unsafe.Sizeof(p2), " ", unsafe.Alignof(p2))
}
