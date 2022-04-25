package main

import (
	"log"
	"reflect"
)
// slice map func 不能比较 .要比较的话使用refluect.DeepEqual
type person struct {
	Name     string
	children string
	Clothes  []cloth
}

type cloth struct{
	price int
}

func main() {
	p1 := person{
		Name: "123",
	}

	p2 := person{
		Name: "123",
	}

	log.Println(reflect.DeepEqual(p1, p2 ))
	return
}
