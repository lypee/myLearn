package main

import (
	"log"
	"time"
)

var (
	formatStr = "2006.01.02 15:04"
	formatStr2 = "2006-01-02 15:04:05"
)
func timeDiff() {
	t1 := time.Now()
	//t2 := time.Unix(1647846990, 0)
	log.Println(t1.Format(formatStr2))
	//log.Println(t1.Before(t2))
}


type Person struct{

}

func main() {
	timeDiff()
	//timeNow := time.Now().Unix()
	//log.Println(time.Unix(timeNow ,0).Format(formatStr))
	resMap := map[int]int{
		1:1,
		2:5,
		3:7,
	}
	//list := []int{1,2,3,4,5}
	//for _ , data := range list{
	//	if data ,exist := resMap[data] ;exist {
	//		log.Println(data , " " , exist)
	//	}
	//}
	//
	for data , exist := range resMap{
		log.Println(data , " " , exist)
	}
	//log.Println(resMap[5])
}


