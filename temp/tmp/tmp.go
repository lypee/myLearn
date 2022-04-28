package main

import (
	"lpynnng/learn/lockMap"
	"time"
)

func main(){
	lockMap.Mcache.Set("1" ,"1")
	lockMap.Mcache.Get("1")
	get2()
	get3()
	time.Sleep(time.Hour)
}

func get2(){
	lockMap.Mcache.Get("2")
}