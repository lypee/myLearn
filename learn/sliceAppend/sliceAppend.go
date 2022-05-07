package main

import (
	"log"
	"sort"
	"time"
	"unsafe"
)

type Person struct {
	Age        int
	Name       string
	CreateTime time.Time
}

func main() {
	testAppend()
	return
}

func testAppend() {
	s := make([]int, 0)
	log.Println(unsafe.Pointer(&s), " ", s, " ", len(s), " ", cap(s))

	testAppend2(&s)
	log.Println(unsafe.Pointer(&s), " ", s, " ", len(s), " ", cap(s))
}
func testAppend2(s *[]int) {
	//log.Println("testAppend2 ", unsafe.Pointer(&s), " ", s, " ", len(s), " ", cap(s))
	*s = append(*s, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	//log.Println("testAppend2 ", unsafe.Pointer(&s), " ", s, " ", len(s), " ", cap(s))

}
func testEnter() {
	p1 := Person{
		Age:  1,
		Name: "",
	}
	p2 := Person{
		Age:  2,
		Name: "",
	}
	p3 := Person{
		Age:  3,
		Name: "",
	}
	p4 := Person{
		Age: 4,
	}
	p5 := Person{
		Age: 5,
	}
	s1 := make([]Person, 0)
	s1 = append(s1, p1, p2)
	s2 := make([]Person, 0)
	s2 = append(s2, p3, p4, p5)
	//log.Println(s1)
	//handle(s1)
	//log.Println(s1)
	test()
}
func handle(arr []Person) []Person {
	arr = arr[:1]
	return arr
}
func test() {
	s := make([]Person, 0, 1)
	p1 := Person{
		Age:        1,
		Name:       "",
		CreateTime: time.Unix(1646620003, 0),
	}
	p2 := Person{
		Age:        2,
		Name:       "",
		CreateTime: time.Unix(1646621003, 0),
	}
	p3 := Person{
		Age:        3,
		Name:       "",
		CreateTime: time.Unix(1646622503, 0),
	}
	p4 := Person{
		Age:        4,
		CreateTime: time.Unix(1646622203, 0),
	}
	p5 := Person{
		Age:        5,
		CreateTime: time.Unix(1646622003, 0),
	}
	//id := 4
	s = append(s, p1, p2, p3, p4, p5)
	//for i := 0 ; i < len(s) ; i++{
	//
	//	if s[i].Age == id{
	//		if i + 1 >= len(s) {
	//			log.Println("maxxx")
	//			break
	//		}
	//		for j := i +1 ; j < len(s) ;j++{
	//			if s[j].Age > 4 {
	//				continue
	//			}else{
	//				log.Println(s[j])
	//				break
	//			}
	//
	//		}
	//	}
	//}
	//log.Println(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i].CreateTime.Before(s[j].CreateTime)
	})
	log.Println(s)
}
