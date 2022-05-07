package main

import "log"

func main() {
	myPanic()
}

func myPanic() {
	hasPanic := false
	normalReturn := false
	defer func() {
		log.Println("hasPanic: ", hasPanic)
		log.Println("normalReturn: ", normalReturn)
	}()

	defer func() {
		if !hasPanic{
			if r := recover() ; r != nil{
				log.Println("recover! ")
			}
		}
	}()
	func() {
		recover()
		happenPanic()
		hasPanic = true
	}()
	if !hasPanic {
		normalReturn = true
	}

}

func happenPanic() {
	defer func() {
		recover()
	}()
	panic("1")
}
