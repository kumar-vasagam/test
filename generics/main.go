package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  30,
		"second": 25,
	}

	floats := map[string]float64{
		"first":  24.99,
		"second": 29.59,
	}

	fmt.Printf("Summing non generics way %v & %v \n", SumInts(ints), SumFloats(floats))

	fmt.Printf("Sum the Ints or Floats on Ints %v \n", SumIntsOrFloats(ints))
}

type Number interface {
	int64 | float64
}

func SumInts(item map[string]int64) int64 {
	var s int64
	for _, v := range item {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var f float64
	for _, v := range m {
		f += v
	}
	return f
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloatsByInterface[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
