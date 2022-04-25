package main

import (
	"testing"
)

type Person struct {
	Age            int
	Name           string
	InteractionId  string `json:"interactionId"`
	PlanId         int64  `json:"planId"`
	BizId          int    `json:"bizId"`
	BusinesslineId int    `json:"businesslineId"`
	BusinessType   int    `json:"businessType"`
}

//func BenchmarkTestSprintf(b *testing.B) {
//	nums := 2 << 20
//	log.Println(nums)
//	err := errors.New("errors")
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		fmt.Sprintf("%+v", err)
//	}
//
//
//	//p := Person{
//	//	Age:            11,
//	//	Name:           "312",
//	//	InteractionId:  "3412123",
//	//	PlanId:         1231,
//	//	BizId:          3124,
//	//	BusinesslineId: 1231,
//	//	BusinessType:   031241,
//	//}
//	//
//	//b.ResetTimer()
//	//for i := 0; i < nums ; i++ {
//	//	fmt.Sprintf("%v", p)
//	//}
//}

func BenchmarkTestRangeStruct1(b *testing.B) {
	arrs := make([]Person, 10000, 10000)
	var tmp Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arrs); j++ {
			tmp = arrs[j]
		}
		_ = tmp
	}
}

func BenchmarkTestRangeStruct2(b *testing.B) {
	arrs := make([]Person, 10000, 10000)
	var tmp Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range arrs {
			tmp = item
		}
		_ = tmp
	}
}

func BenchmarkTestRangeStruct3(b *testing.B) {
	arrs := make([]Person, 10000, 10000)
	var tmp Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index, _ := range arrs {
			tmp = arrs[index]
		}
		_ = tmp
	}
}

