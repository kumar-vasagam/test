package main

import "fmt"

func main() {

	fmt.Printf("Wanted  2 | Got %v \n ", removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Printf("Wanted  5 | Got %v \n ", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}

func removeElement(nums []int, val int) int {
	i, j, k := 0, len(nums), 0

	for {
		if nums[i] == val {
			if nums[j] == val {
				j--

			} else {
				//Swap it to the last position
			}
		}
		i++
		if i == len(nums) {
			break
		}
	}
	return k
}
