package main

import (
	"fmt"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func findSum() {
	sumOf := 3
	numbers := []int{4, 5, -1, 2, 6, -2}
	oldSum := 0
	finalIndexStart := 0
	count := len(numbers) + 1 - sumOf
	newSum := 0
	for i := 0; i < count; i++ {
		newSum = 0
		for j := i; j < sumOf+i; j++ {
			newSum += numbers[j]
		}
		if oldSum > newSum {
			//do nothing
		} else {
			oldSum = newSum
			finalIndexStart = i
		}
	}
	fmt.Printf("finalIndexStart %v and the final sum %v \n: ", finalIndexStart, newSum)
}

func reverseString(input string) {
	//input := "I am a man'"
	//output := "man' a am I"
	s := strings.Split(input, " ")
	//outputA := make([]string, len(inputA), len(inputA))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Printf("Reversed? %q", s)
}

func removeDuplicates(nums []int) int {
	return 0
}

func theLowerOf(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	var temp int64
	bb := int64(b)
	ww := int64(w)
	bbc := int64(bc)
	wwc := int64(wc)
	zz := int64(z)

	acqCost := bb*bbc + ww*wwc
	if z >= bc && z >= wc || bc == wc {
		//No difference
		return int64(acqCost)
	}
	if bc > wc {
		temp = ww*wwc + bb*wwc + bb*zz
		return theLowerOf(temp, acqCost)
	}
	if wc > bc {
		temp = bb*bbc + ww*bbc + ww*zz
		return theLowerOf(temp, acqCost)
	}
	return int64(acqCost)

}

func main() {
	result := add(3, 7)
	fmt.Println(result)
	findSum()
	reverseString("I am a man'")
	reverseString("Tom Weds Anne")

	fmt.Printf("Want 2 | Got %v : \n", removeDuplicates([]int{1, 1, 2}))
	fmt.Printf("Want 5 | Got %v : \n", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Printf("Want 20 | Got %v : \n", taumBday(10, 10, 1, 1, 1))
	fmt.Printf("Want 37 | Got %v : \n", taumBday(5, 5, 2, 3, 4))
	fmt.Printf("Want 12 | Got %v : \n", taumBday(3, 6, 9, 1, 1))
	fmt.Printf("Want 35 | Got %v : \n", taumBday(7, 7, 4, 2, 1))
	fmt.Printf("Want 12 | Got %v : \n", taumBday(3, 3, 1, 9, 2))
	fmt.Printf("Want 18192035842 | Got %v : \n", taumBday(27984, 1402, 619246, 615589, 247954))

}
