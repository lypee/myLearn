package main

import (
	"log"
	"reflect"
)

type Person struct {
	Name string `json:"name" gson:"123"`
	Age  int    `json:"age"`
}

func main() {
	var s []string
	s = append(s , "1")
	//log.Println(s)
	//log.Println(1)
	t := Person{}
	log.Println(reflect.TypeOf(t).Field(0).Tag.Get("gson"))
	for  i := 0 ; i < reflect.TypeOf(t).NumField() ; i++{
		log.Println(reflect.TypeOf(t).Field(i).Tag.Get("json"))
	}
}
