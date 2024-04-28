package main

import (
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(time.Millisecond)
}

func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}

// 1 cpu
// BenchmarkPrint1         131558955                8.894 ns/op
// BenchmarkGoPrint1        1152337              1009 ns/op
// BenchmarkPrint2             6583            346539 ns/op
// BenchmarkGoPrint2        1000000              7905 ns/op

// 2 cpu
// BenchmarkPrint1-2       135320718                8.866 ns/op
// BenchmarkGoPrint1-2      3076652               384.4 ns/op
// BenchmarkPrint2-2          11536            109327 ns/op
// BenchmarkGoPrint2-2       293672              5808 ns/op