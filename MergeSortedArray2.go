package main

import "fmt"

func main() {
	fmt.Printf("Wanted : [1, 2, 2, 3, 5, 6] : Got %v \n", merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	fmt.Printf("Wanted : [1] : Got %v \n", merge([]int{1}, 1, []int{}, 0))
	fmt.Printf("Wanted : [1] : Got %v \n", merge([]int{0}, 0, []int{1}, 1))

}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if n == 0 {
		return nums1
	}
	if m == 0 {
		nums1 = nums2
		return nums1
	}
	j := 0
	for i := n; i < m+n; i++ {
		nums1[i] = nums2[j]
		//fmt.Printf("nums1[i] %v | nums2[j] %v\n", nums1[i], nums2[j])
		j++
	}
	fmt.Printf("nums1: %v\n", nums1)
	for i := 0; i < len(nums1)-1; i++ {
		if nums1[i] > nums1[i+1] {
			c := nums1[i+1]
			nums1[i+1] = nums1[i]
			nums1[i] = c
			continue
		}
	}
	return nums1
}
