package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2})
	sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	fmt.Println(pageCount(6, 2))
	fmt.Println(pageCount(5, 4))
	out := countingValleys(0, "DDUUDDUDUUUD")
	fmt.Printf("valley: %v\n", out)

	getMoneySpent([]int32{40, 50, 60}, []int32{5, 8, 12}, 60)
	getMoneySpent([]int32{5, 8, 12}, []int32{40, 50, 60}, 60)
	getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10)
	getMoneySpent([]int32{4}, []int32{5}, 5)
	formingMagicSquare([][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}})
	gradingStudents([]int32{73, 67, 38, 33})
	getTotalX([]int32{2, 6}, []int32{24, 36})
	getTotalX([]int32{2, 4}, []int32{16, 32, 96})
	getTotalX([]int32{}, []int32{})
	getTotalX([]int32{3, 9, 6}, []int32{36, 72})

}
func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	pairs := make(map[int32]int32, 0)
	for _, v := range ar {
		_, ok := pairs[v]
		if ok {
			pairs[v]++
		} else {
			pairs[v] = 1
		}
	}
	var numPairs int32
	for _, v := range pairs {
		numPairs += v / 2
	}
	fmt.Printf("numPairs: %v\n", numPairs)
	return numPairs
}
func pageCount(n int32, p int32) int32 {
	//var toTurn int32

	x := p / 2
	y := n/2 - p/2

	min := math.Min(float64(x), float64(y))

	return int32(min)
}

func countingValleys(steps int32, path string) int32 {
	// Write your code here

	arr := make([]rune, 0)
	for _, r := range path {
		arr = append(arr, r)
	}
	fmt.Printf("arr: %v\n", arr)
	sum := 0
	valleys := 0
	for i := 0; i < len(arr); i++ {
		if sum == 0 && string(arr[i]) == "D" {
			valleys++
			sum--
			continue
		}
		if string(arr[i]) == "U" {
			sum++
		}
		if string(arr[i]) == "D" {
			sum--
		}
	}
	return int32(valleys)
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	slices.Sort(keyboards)
	slices.Sort(drives)
	var totalCost int32 = -1
	for i := len(keyboards) - 1; i >= 0; i-- {
		for j := len(drives) - 1; j >= 0; j-- {
			if keyboards[i]+drives[j] == b {
				return b
			}
			if keyboards[i]+drives[j] > b {
				continue
			}
			if keyboards[i]+drives[j] > totalCost {
				totalCost = keyboards[i] + drives[j]
			}
		}
	}
	fmt.Printf("totalCost: %v\n", totalCost)
	return totalCost
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	a := math.Abs(float64(z) - float64(x))
	b := math.Abs(float64(z) - float64(y))
	if a == b {
		fmt.Println("Mouse C")
		return "Mouse C"
	}
	if a < b {
		fmt.Println("Cat A")
		return "Cat A"
	}
	return "Cat B"

}
func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	sumMap := make(map[string]int32)
	magicNums := make(map[int32]int32)

	var magicNumber int32
	rows := len(s)
	cols := len(s[0])
	for i := 0; i < rows; i++ {
		var colSum int32
		for j := 0; j < cols; j++ {
			colSum += s[i][j]
		}
		sumMap["row"+fmt.Sprint(i)] = colSum
		_, ok := magicNums[colSum]
		if ok {
			magicNums[colSum]++
		} else {
			magicNums[colSum] = 0
		}
	}
	for j := 0; j < cols; j++ {
		var rowSum int32
		for i := 0; i < rows; i++ {
			rowSum += s[i][j]
		}
		sumMap["col"+fmt.Sprint(j)] = rowSum
		_, ok := magicNums[rowSum]
		if ok {
			magicNums[rowSum]++
		} else {
			magicNums[rowSum] = 0
		}
	}
	diag := s[0][0] + s[1][1] + s[2][2]
	sumMap["dia0"] = diag
	fmt.Printf("sumMap: %v\n", sumMap)
	magicNums[diag]++
	fmt.Printf("magicNums: %v\n", magicNums)
	var o int32
	for k, v := range magicNums {
		if v > o {
			o = v
			magicNumber = k
		}
	}

	fmt.Printf("magicNumber: %v\n", magicNumber)
	return magicNumber
}
func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var out = make([]int32, 0)
	for _, val := range grades {
		if val < 38 {
			out = append(out, val)
			continue
		}
		q := val / 5
		r := val % 5
		if r == 0 {
			out = append(out, val)
			continue
		}
		if r >= 3 {
			out = append(out, (q+1)*5)
		} else {
			out = append(out, val)
		}

	}
	fmt.Printf("new grades : %v\n", out)
	return out
}
func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	slices.Sort(a)
	slices.Sort(b)
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	start := a[len(a)-1]
	end := b[len(b)-1]
	var count int32
	for {
		var factor1, factor2 bool = false, false
		for _, f := range a {
			if start%f == 0 {
				factor1 = true
			} else {
				factor1 = false
				break
			}
		}
		for _, s := range b {
			r := s % start
			if r == 0 {
				factor2 = true
			} else {
				factor2 = false
				break
			}
		}
		if factor1 == true && factor2 == true {
			fmt.Printf("Matching number %v \n", start)
			count++
		}
		start++
		if start > end {
			break
		}
	}
	fmt.Printf("count: %v\n", count)
	return count

}
