package main

import (
	"encoding/json"
	"log"
	"reflect"
	"sync"
	"unsafe"
)

var (
	nums = 0
)

func main() {
	var nilSlice []int
	s1 := []int{}
	log.Println(nilSlice[1])
	getSliceAllPtr("nilSlice ", &nilSlice)
	getSliceAllPtr("s1 ", &s1)

	log.Println("nil切片与nil比较的结果: ", nil == nilSlice)
	log.Println("nil切片占用内存大小: ", unsafe.Sizeof(nilSlice))
	log.Println("s1切片占用内存大小: ", unsafe.Sizeof(s1))

	//enterTransfer()
}

func enterTransfer() {
	s := []int{1, 2, 3}
	getSlicePtr("before.s ", &s)
	log.Println("before.s.Data: ", s)

	transferTest1V2(&s)

	getSlicePtr("after.s ", &s)
	log.Println("after.s.Data: ", s)

}

func transferTest(s []int) {
	log.Println("into-transferTest1")
	getSlicePtr("before.s ", &s)
	log.Println("before.s.Data: ", s)
	s[0] = 999
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	getSlicePtr("after.s ", &s)
	log.Println("out-transferTest1")
}

func transferTest1V2(s *[]int) {
	log.Println("into-transferTest1")
	getSlicePtr("before.s ", s)
	log.Println("before.s.Data: ", s)
	for i := 0; i < 10; i++ {
		*s = append(*s, i)
	}
	getSlicePtr("after.s ", s)
	log.Println("out-transferTest1")
}

const (
	basePrintInfo   = "%s切片的指针地址:[%v]，切片数组地址:[%v]"
	basePrintInfoV2 = "%s切片的指针地址:[%v]，切片数组地址:[%v] ,Len字段的地址:[%p] ,Cap字段地址:[%p]"
)

func myMakeSlice() {
	s1 := []int{}
	s1 = append(s1, 1)
	s4 := s1[:1]
	log.Println("append之前")
	getSlicePtr("s1 ", &s1)
	getSlicePtr("s4 ", &s4)
	s1 = append(s1, 2, 3, 4, 5, 6, 7, 8)
	log.Println("append之后")
	getSlicePtr("s1 ", &s1)
	getSlicePtr("s4 ", &s4)
	//log.Println("再声名一个数组看看")
	//s5 := s1[:len(s1)]
	//getSlicePtr("s1 ", s1)
	//getSlicePtr("s4 ", s5)
}

func getSlicePtr(name string, s1 *[]int) {
	s2 := (*reflect.SliceHeader)(unsafe.Pointer(s1))
	log.Printf(basePrintInfo, name, unsafe.Pointer(s1), unsafe.Pointer(s2.Data))
}

func getSliceAllPtr(name string, s1 *[]int) {
	s2 := (*reflect.SliceHeader)(unsafe.Pointer(s1))
	log.Printf(basePrintInfoV2, name, unsafe.Pointer(&s1), unsafe.Pointer(s2.Data), unsafe.Pointer(&s2.Len), unsafe.Pointer(&s2.Cap))
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

func myTestAppend() {
	//runtime.GOMAXPROCS(1)
	time := 1000
	wg := &sync.WaitGroup{}
	wg.Add(time)
	for i := 0; i < time; i++ {
		go func() {
			defer wg.Done()
			nums++
		}()
	}
	wg.Wait()
	log.Println(nums)
	return
}

func myRealMakeSlice() {
	var i int
	myType := reflect.TypeOf(i)
	slice := reflect.MakeSlice(reflect.SliceOf(myType), 10, 10).Interface()
	p := slice.([]int)
	p[0] = 1
}

func d() {

}

func testPrint2() {
	//log.Printf("reflect.SliceHeader结构占用的内存大小为:[%d]",unsafe.Sizeof(reflect.SliceHeader{}))
	//log.Println( 824634400792 - 824634400816)
	//log.Printf("reflect.SliceHeader的内存大小为:[%d]",unsafe.Sizeof(reflect.Int))
}
