package main

import "log"

func main(){
	n1 , n2, n3, n4,n5,n6 := 0,0,0,0,0,0
	ch1 := make(chan int )
	//ch1 <- 1
	//ch1 <- 1
	//ch1 <- 1

	ch2 := make(chan int )
	//ch2 <- 2
	//ch2 <- 2
	//ch2 <- 2

	ch3 := make(chan int )
	//ch3 <- 3
	//ch3 <- 3
	//ch3 <- 3

	ch4 := make(chan int)
	//ch4 <- 4
	//ch4 <- 4
	//ch4 <- 4

	ch5 := make(chan int)
	//ch5 <- 5
	//ch5 <- 5
	//ch5 <- 5
	for i := 0 ; i < 100 ;i++{
		select{
		case <- ch1 :
			n1++
		case <- ch2 :
			n2++
		case <- ch3 :
			n3++
		case <- ch4 :
			n4++
		case <- ch5 :
			n5++
		default:
			n6++
		}
	}
	log.Println(n1 ," " ,n2 ," " ,n3 ," " ,n4 ," " ,n5 ," ",n6)
	return
}
