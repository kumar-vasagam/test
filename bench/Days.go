package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("b days %v \n", beautifulDays(20, 23, 6))
	fmt.Printf("b days %v \n", improvedBeautifulDays(20, 23, 6))
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var bdays int32 = 0

	for x := i; x <= j; x++ {
		ri := strconv.Itoa(int(x))
		mult := 1
		var revi = 0
		for z := 0; z < len(ri); z++ {
			basn := string(ri[z])
			zz, _ := strconv.Atoi(basn)
			revi += zz * mult
			mult *= 10
		}
		diff := x - int32(revi)
		//fmt.Printf("i %v reversed i %v diff %v \n", x, revi, diff)
		if diff%k == 0 {
			bdays++
		}
	}

	return bdays
}
func improvedBeautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var bdays int32 = 0
	for x := i; x <= j; x++ {
		var revi int32 = 0
		number := x
		for {
			tmp := number % 10
			revi = revi*10 + tmp
			number /= 10
			if number == 0 {
				break
			}
		}
		diff := x - int32(revi)
		fmt.Printf("i %v reversed i %v diff %v \n", x, revi, diff)
		if diff%k == 0 {
			bdays++
		}
	}
	return bdays
}
