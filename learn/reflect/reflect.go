package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string
}

func main() {
	//p := []Person{
	//	{Name: "1"},
	//}
	p2 := []int{}
	//print2(p)
	//print2(p2)
	//
	//var p3 []int
	//print2(p3)
	//p4 := [3]int{}
	//print2(p4)

	//var i int
	//var str string
	//print(str)
	//str := "123"
	print2(p2)
}

func print(p interface{}) {

	switch p.(type) {
	case int:
		log.Println("int")
	case string:
		log.Println("string")
	case struct{}:
		log.Println("struct{}")
	default:
		log.Println("default")
	}
}

func print2(p interface{}) {
	v := reflect.TypeOf(p).Kind()
	switch v {
	case reflect.String:
		log.Println("String")
	case reflect.Int:

		log.Println("Int")
	case reflect.Struct:
		log.Println("Struct")
	case reflect.Array:
		log.Println("Array")
	case reflect.Slice:
		vv := reflect.ValueOf(v)
		sv := reflect.MakeSlice(vv.Type(), vv.Len(), vv.Cap())
		log.Println(sv)
		log.Println("Slice")
	case reflect.Interface:
		log.Println("Interface")
	case reflect.Ptr:
		log.Println("Ptr")
	default:
		log.Println("default")
	}
}

func print3(p interface{}) bool {
	if reflect.TypeOf(p).Kind() == reflect.Int {
		return false
	}
	return true
}
