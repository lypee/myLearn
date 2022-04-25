package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)
type Person struct{
	Age int
}
func main(){
	s := &Person{Age: 1}
	log.Println(reflect.TypeOf(s).Kind() == reflect.Ptr)
}

func BenchmarkSprintf(b *testing.B){
	nums := 10
	b.ResetTimer()
	for i := 0 ; i < nums ; i++{
		fmt.Sprintf("%d" , nums)
	}
}