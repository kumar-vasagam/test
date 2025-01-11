package main

import (
	"fmt"
	"time"
)

func main() {
	timerTests()
	fmt.Printf("equalizeArray Want 2 | Got %v \n", equalizeArray([]int32{3, 3, 2, 1, 3}))
	fmt.Printf("equalizeArray Want 2 | Got %v \n", equalizeArray([]int32{3, 3, 3, 1, 1}))
	fmt.Printf("equalizeArray Want 5 | Got %v \n", equalizeArray([]int32{3, 3, 3, 1, 1, 4, 5, 6}))

}

func timerTests() {
	layout := "2006-01/02 15:04:05 PM"
	date := "2013-02/03 07:45:23 AM"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error in time")
		fmt.Printf("err.Error(): %v\n", err.Error())
	}
	fmt.Printf("t: %v\n", t)
}
func equalizeArray(arr []int32) int32 {
	// Write your code here
	dmap := make(map[int]int)
	for _, v := range arr {
		_, ok := dmap[int(v)]
		if ok {
			dmap[int(v)]++
		} else {
			dmap[int(v)] = 1
		}
	}
	var min = 0
	for _, v := range dmap {
		if v > 1 {
			if v > min {
				min = v
			}
		}
	}
	if min == 0 {
		min = len(arr) - 1
	} else {
		min = len(arr) - min
	}
	return int32(min)
}
