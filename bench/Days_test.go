package main

import "testing"

var result int32

func BenchmarkBeautifulDays(b *testing.B) {
	var s int32
	for i := 0; i < b.N; i++ {
		s = beautifulDays(20, 23, 6)
	}
	result = s

}

func BenchmarkImprovedBeautifulDays(b *testing.B) {
	var s int32
	for i := 0; i < b.N; i++ {
		s = improvedBeautifulDays(20, 23, 6)
	}
	result = s

}
