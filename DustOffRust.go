package main

import "fmt"

func main() {
	fmt.Println("Hello Dustting Off Rust!!!")
	PlayWithSlices()
}

func PlayWithSlices() {
	//Declare a literatl slice
	numSlice := []int{1, 2, 3, 4, 5}

	//Declare in another way
	alphaSlice := make([]string, 0, 5)

	fmt.Printf("Num Slice %v \n", numSlice)
	fmt.Printf("alpha Slice len %v and contents %v \n", len(alphaSlice), alphaSlice)

	for _, item := range numSlice {
		fmt.Printf("item: %v\n", item)
	}
	alphaSlice = append(alphaSlice, "Tom")
	alphaSlice = append(alphaSlice, "Mary")

	count := 0
	for {
		fmt.Printf("alphaSlice[count]: %v\n", alphaSlice[count])
		count++
		if count == len(alphaSlice) {
			break
		}
	}

	x := make([]rune, 4)
	x[0] = 'u'
	x[2] = 't'
	fmt.Printf("x: %v\n", x)

}
