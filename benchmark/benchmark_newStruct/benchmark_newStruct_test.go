package benchmark_newStruct

import (
	"testing"
	"time"
)

type Energy struct {
	EnergyType   int
	EnergyText   string
	Now          time.Time
	EnergyType2  int
	EnergyText2  string
	Now2         time.Time
	EnergyType33 int
	EnergyText3  string
	Now3         time.Time
	EnergyType4  int
	EnergyText5  string
	Now5         time.Time
	EnergyType6  int
	EnergyText6  string
	Now6         time.Time
}

func Benchmark_newStruct1(b *testing.B) {
	eList := make([]Energy, 0, 1000)
	for i := 0; i < 1000; i++ {
		eList = append(eList, Energy{EnergyText2: "1", Now: time.Now()})
	}
	//var energy Energy
	//log.Println(unsafe.Pointer(&e1))
	b.ResetTimer()
	for idx := 0; idx < b.N*100; idx++ {
		for i := range eList {
			energy := eList[i]
			_ = energy
			//log.Println(fmt.Sprintf("%p", &energy))
		}
	}
	return
}

func Benchmark_newStruct2(b *testing.B) {

	eList := make([]Energy, 0, 1000)
	for i := 0; i < 1000; i++ {
		eList = append(eList, Energy{EnergyText2: "1", Now: time.Now()})
	}
	var energy Energy
	//log.Println(unsafe.Pointer(&e1))
	b.ResetTimer()
	for idx := 0; idx < b.N*100; idx++ {
		for i := range eList {
			energy = eList[i]
			_ = energy
			//log.Println(fmt.Sprintf("%p", &energy))
		}
	}

	return
}
