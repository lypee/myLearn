package main

import (
	"bytes"
	"testing"
)

func BenchmarkGenRandNumberByRange(b *testing.B) {
	GenRandNumberByRange(100, 999)

	b.ReportAllocs()
}


var strLen = 1000

func BenchmarkConcatString(b *testing.B) {
	var str string
	i := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"

		i++
		if i >= strLen {
			i = 0
			str = ""
		}
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer

	i := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")

		i++
		if i >= strLen {
			i = 0
			buffer = bytes.Buffer{}
		}
	}
}