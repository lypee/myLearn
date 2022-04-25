package main

import (
	"encoding/json"
	"log"
	"unsafe"
)

type Args struct {
	PkCommonID
}
type ArgsWithClassId struct {
	PkCommonID
	ClassId int `json:"classId"`
}
type PkCommonID struct {
	BizID      int `json:"bizId"`
	PlanID     int `json:"planId"`
	PkType     int `json:"pkType"`     // pk类型 常量定义搜索PkTypeDefault
	IsPlayback int `json:"isPlayback"` // 0-直播 1-回放

	PkID         int64 `json:"pkId"`
	GroupID      int64 `json:"groupId"`
	RivalGroupID int64 `json:"rivalGroupId"`
}

type Person struct {
	Mmap1 map[int]int
	Mmap2 map[int]int
	Mmap3 map[int]int
}

func main() {
	sizeOf()
}

func sizeOf(){
	mmap1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	str1  , _ := json.Marshal(mmap1)

	log.Println(unsafe.Sizeof(PkCommonID{}) , " " , unsafe.Sizeof(Args{})  , unsafe.Sizeof(ArgsWithClassId{}))
}
func unMarshalMap() {
	mmap1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	mmap2 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	mmap3 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	mmap4 := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	p := Person{
		Mmap1: mmap1,
		Mmap2: mmap2,
		Mmap3: mmap3,
	}

	log.Println(unsafe.Sizeof(mmap1), " ", unsafe.Sizeof(p), " ", unsafe.Sizeof(mmap4))
	mmapStr, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}
	pp := Person{}
	err = json.Unmarshal(mmapStr, &pp)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(pp)
	return
}
func testMapNil() {
	mmap := map[int]Person{}
	mmap[1] = Person{}
}

func testSliceNil() {
	var s1 []int
	s1 = []int{}
	s2 := []int{}
	log.Println(nil == s1)
	log.Println(nil == s2)
}
