package main

import (
	"fmt"
	"time"
)

func main() {
	timerTests()
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
