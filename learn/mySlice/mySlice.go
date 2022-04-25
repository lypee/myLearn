package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var (
	nums = 0
)
func main() {
	myTestAppend()
}
func myTestAppend() {
	//runtime.GOMAXPROCS(1)
	time := 1000
	wg := &sync.WaitGroup{}
	wg.Add(time)
	for i := 0 ; i < time ; i++{
		go func() {
			defer wg.Done()
			nums++
		}()
	}
	wg.Wait()
	log.Println(nums)
	return
}

func myMakeSlice() {
	baseStr := "指针地址为:%p"
	s1 := []int{1, 2, 3, 1, 1, 1, 1, 1}
	log.Println(fmt.Sprintf(baseStr, s1))
	s1 = append(s1, 1)
	log.Println(fmt.Sprintf(baseStr, s1))
	//var x int
	//x, _ := f()
	//x, _ = f()
	//x, y := f()
	//x, y = f()

}
func f() (int, int) {
	return 0, 0
}
func testPrint() {
	p1 := Person{
		Age:  1,
		Name: "1",
	}
	p2 := Person{
		Name: "1",
		Age:  1,
	}
	p1Str, _ := json.Marshal(&p1)
	p2Str, _ := json.Marshal(&p2)
	log.Println(string(p1Str))
	log.Println(string(p2Str))
}
func watch(p []Person) {
	for {
		return
	}
}
func test2(arr *[]int) {
	*arr = append(*arr, 2)
	return
}
func test1() {
	//stuLen := 1
	p1List := make([]Person, 0, 1)
	p1List = append(p1List, Person{
		Age:  1,
		Name: "1",
	})
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 3)
	copy(s2, s1)
	log.Println(s1)
	log.Println(s2)
	log.Println()

	s2[1] = 99
	log.Println(&s1)
	log.Println(s1)
	log.Println(s2)
	return
	//pLen := len(p1List)
	//log.Println(reflect.TypeOf(stuLen / pLen))

	//log.Println(1 / 60)

	//log.Println(p1List[0:1])
	////stuLen := 10
	//log.Println(1 /  0)
}

type Req struct {
}

func myString() {
	//str := "123"
	pageSize := 50
	pageNums := 0
	stuLen := 1

	if pageSize > 0 && pageNums >= 0 {
		pageNum := stuLen / pageSize
		if (stuLen % pageSize) > 0 {
			pageNum = pageNum + 1
		}
		curnum := pageNums % pageNum
		log.Println(curnum)
	}
	//log.Println(1 / 50)
	//str2 := str + "4"
	//log.Println(str2)
	//log.Println(reflect.TypeOf(reflect.Map).Kind())
	//p1 := Person{
	//	Age:  11,
	//	Name: "123",
	//}

	log.Println()
	return
}

func mySlice() {
	var s1 []int
	s2 := make([]int, 0)
	s3 := []int{}
	log.Println(nil == s1, " ", len(s1), " ", cap(s1))
	log.Println(nil == s2, " ", len(s2), " ", cap(s3))
	log.Println(nil == s3, " ", len(s2), " ", cap(s3))
}

type Person struct {
	Age  int
	Name string
}

//func query(o interface{}){
//	if reflect.TypeOf(o).Kind() == reflect.Struct{
//		kind := reflect.ValueOf(o)
//	}
//}
