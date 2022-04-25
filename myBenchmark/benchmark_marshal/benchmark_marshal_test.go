package benchmark_marshal

import (
	"encoding/json"
	"testing"
)

type BenchPerson struct {
	Family map[int]int `json:"family"`
	Grow   map[int]int `json:"grow"`
	Age    int         `json:"age"`
	Name   string      `json:"name"`
}
//
func BenchmarkTestNoMarshal(b *testing.B) {
	size := 16
	family := make(map[int]int, size)
	grow := make(map[int]int, size)
	for i := 0; i < size; i++ {
		family[i] = i
		grow[i] = i
	}
	var bp, bp2Un BenchPerson
	var str1 []byte
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		// 1. map
		bp = BenchPerson{
			Family: family,
			Grow:   grow,
			Age:    1,
			Name:   "111",
		}
		// 2. 2json
		str1, _ = json.Marshal(&bp)
		// 3. tobytes
		bytes1 := []byte(str1)
		bp2Un = BenchPerson{}
		// 4.unmarshal
		_ = json.Unmarshal(bytes1, &bp2Un)
	}
}

type BenchPersonMarshal struct {
	Family string `json:"family"`
	Grow   string `json:"grow"`
	Age    int    `json:"age"`
	Name   string `json:"name"`
}

func BenchmarkTestMarshal(b *testing.B) {
	size := 16
	family := make(map[int]int, size)
	grow := make(map[int]int, size)
	for i := 0; i < size; i++ {
		family[i] = i
		grow[i] = i
	}
	var bp ,bp2Un BenchPersonMarshal
	var familyStr ,growStr,str1 []byte

	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		familyStr , _ = json.Marshal(family)
		growStr , _ = json.Marshal(grow)
		bp.Family = string(familyStr)
		bp.Grow = string(growStr)
		bp.Age = 1
		bp.Name = "111"

		str1, _ = json.Marshal(&bp)
		_ = json.Unmarshal(str1 , &bp2Un)
	}

}
