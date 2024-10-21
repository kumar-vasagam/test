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

func reverseString() {
	input := "I am a man'"
	//output := "man' a am I"
	s := strings.Split(input, " ")
	//outputA := make([]string, len(inputA), len(inputA))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Printf("Reversed? %v ", s)
}

func main() {
	result := add(3, 7)
	fmt.Println(result)
	findSum()
	reverseString()
}
