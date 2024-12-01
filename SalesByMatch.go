package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
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
	fmt.Printf("The result : %v\n", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("The result : %v\n", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Printf("The result : %v\n", searchInsert([]int{1, 3, 5, 6}, 7))
	generate(6)
	fmt.Printf("singleNum %v \n", singleNumber([]int{2, 2, 1}))
	fmt.Printf("singleNum %v \n", singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Printf("singleNum %v \n", singleNumber([]int{4}))

	fmt.Printf("pick Nums %v \n", pickingNumbers([]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}))
	fmt.Printf("pick Nums %v \n", pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
	fmt.Printf("pick Nums %v \n", pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
	fmt.Printf("Remove duple elements (2) %v \n", removeDuplicateElements([]int{0, 1, 1, 2}))
	fmt.Printf("Remove duple elements (5) %v \n", removeDuplicateElements([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Printf("Reversal %v \n ", reverseAString("Geeks"))
	fmt.Printf("Reversal %v \n ", reverseAString("Geeks for Nerds"))

	fmt.Printf("checkEqualArrays (T) %v \n ", checkEqualArrays([]int{1, 2, 5, 4, 0}, []int{2, 4, 5, 0, 1}))
	fmt.Printf("checkEqualArrays (F) %v \n ", checkEqualArrays([]int{1, 2, 5}, []int{2, 4, 15}))
	//fmt.Printf(" %v \n")
	fmt.Printf(" findUnion (5) %v \n", findUnion([]int{1, 2, 3, 4, 5}, []int{1, 2, 3}))
	fmt.Printf(" findUnion (7) %v \n", findUnion([]int{85, 25, 1, 32, 54, 6}, []int{85, 2}))
	fmt.Printf(" findUnion (2) %v \n", findUnion([]int{1, 2, 1, 1, 2}, []int{2, 2, 1, 2, 1}))
	fmt.Printf(" isSubset (T) %v \n", isSubset([]int{11, 7, 1, 13, 21, 3, 7, 3}, []int{11, 3, 7, 1, 7}))
	fmt.Printf(" isSubset (T) %v \n", isSubset([]int{1, 2, 3, 4, 4, 5, 6}, []int{1, 2, 4}))
	fmt.Printf(" isSubset (F) %v \n", isSubset([]int{10, 5, 2, 23, 19}, []int{19, 5, 3}))
	fmt.Printf(" search 2 %v \n", search([]int{1, 2, 3, 4}, 3))
	fmt.Printf(" search 4 %v \n", search([]int{10, 8, 30, 4, 5}, 5))
	fmt.Printf(" search -1 %v \n", search([]int{10, 8, 30}, 6))

	rotateAtIndex([]int{1, 2, 3, 4, 5}, 3)
	rotateAtIndex([]int{5, 6, 8, 9}, 5)
	climbingLeaderboard([]int32{100, 90, 90, 80}, []int32{70, 80, 105})
	climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
	fmt.Printf("Want 3 | Got %v \n", saveThePrisoner(4, 6, 2))
	fmt.Printf("Want 2 | Got %v \n", saveThePrisoner(5, 2, 1))
	fmt.Printf("Want 3 | Got %v \n", saveThePrisoner(5, 2, 2))
	fmt.Printf("Want 6 | Got %v \n", saveThePrisoner(7, 19, 2))
	fmt.Printf("Want 3 | Got %v \n", saveThePrisoner(3, 7, 3))
	fmt.Printf("Want ? | Got %v \n", saveThePrisoner(6, 737005495, 6))
	fmt.Printf("Want ? | Got %v \n", saveThePrisoner(654809340, 204894365, 472730208))
	fmt.Printf("Want ? | Got %v \n", saveThePrisoner(12, 430895283, 10))

	circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2})

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

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
func generate(numRows int) [][]int {
	out := make([][]int, 0)
	fmt.Printf("out: %v\n", out)
	curRow := 1
	for {
		in := make([]int, 0)
		fmt.Printf("in: %v\n", in)

		for i := 0; i < curRow; i++ {
			if i == 0 || i == curRow-1 {
				in = append(in, 1)
				continue
			}
			fmt.Printf("out: %v\n", out)
			previous := out[len(out)-1][i] + out[len(out)-1][i-1]
			fmt.Printf("first %v | second %v | result %v \n", out[len(out)-1][i], out[len(out)-1][i-1], previous)
			in = append(in, previous)
		}
		out = append(out, in)
		curRow++
		if curRow > numRows {
			break
		}
	}
	fmt.Printf("out: %v\n", out)
	return out
}
func singleNumber(nums []int) int {
	items := make(map[int]int)
	var single int
	for _, item := range nums {
		_, ok := items[item]
		if ok {
			items[item]++
		} else {
			items[item] = 1
		}
	}
	for k, v := range items {
		if v == 1 {
			return k
		}
	}
	return single
}
func pickingNumbers(a []int32) int32 {
	// Write your code here
	//Sort
	var maxSize int32
	//

	slices.Sort(a)
	b := make([]int32, 0)
	for i := 0; i < len(a)-1; i++ {
		if math.Abs(float64(a[i])-float64(a[i+1])) == 1 || math.Abs(float64(a[i])-float64(a[i+1])) == 0 {
			//fmt.Printf("appending  %v \n ", a[i])
			b = append(b, a[i])
			if i == len(a)-2 {
				b = append(b, a[i+1])
			}
		} else {
			b = append(b, a[i])
			if i == len(a)-2 {
				b = append(b, a[i+1])
			}
			//fmt.Printf("b: %v\n", b)
			b = b[:0]
		}
		if len(b) > int(maxSize) {
			maxSize = int32(len(b))
		}
	}
	return maxSize

}
func removeDuplicateElements(nums []int) int {
	slices.Sort(nums)
	var seen = nums[0]
	i := 1

	for j := 1; j < len(nums); j++ {
		if nums[j] == seen {
			continue
		} else {
			nums[i] = nums[j]
			seen = nums[j]
			i++
		}
	}
	return i
}
func reverseAString(s string) string {
	var sb strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}

	fmt.Printf("sb.String(): %v\n", sb.String())

	return sb.String()

}

func checkEqualArrays(a1 []int, a2 []int) bool {
	slices.Sort(a1)
	slices.Sort(a2)
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
func findUnion(a1 []int, a2 []int) int {
	m := make(map[int]bool)
	var count int

	f := func(a []int) {
		for _, v := range a {
			_, ok := m[v]
			if !ok {
				m[v] = true
				count++
			}
		}
		fmt.Printf("m: %v\n", m)
	}

	f(a1)
	f(a2)
	fmt.Printf("count: %v\n", count)
	return count
}
func isSubset(a1 []int, a2 []int) bool {
	i1, i2, count := 0, 0, 0
	slices.Sort(a1)
	slices.Sort(a2)
	for {
		if a2[i2] == a1[i1] {
			i1++
			i2++
			count++
		} else {
			i1++
		}
		if i2 == len(a2)-1 || i1 == len(a1)-1 {
			break
		}
	}
	if count == len(a2)-1 {
		return true
	}
	return false
}

func search(a []int, x int) int {
	count := -1
	for i, v := range a {
		if v == x {
			count = i
			break
		}
	}
	return count

}

func rotateAtIndex(a []int, k int) {
	fmt.Printf("a BEFORE : %v\n", a)
	x := k / 2
	if k >= len(a) {
		x = len(a) / 2
		k = len(a)
	}
	j := 0
	for i := k - 1; i >= x; i-- {
		a[i], a[j] = a[j], a[i]
		j++
	}
	fmt.Printf("a AFTER: %v\n", a)
}

type position struct {
	score int32
	rank  int
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	output := make([]int32, 0)
	clearRanks := make([]int32, 0)
	for i, v := range ranked {
		if i == 0 || v != ranked[i-1] {
			clearRanks = append(clearRanks, v)
		}
	}
	fmt.Printf("clearRanks: %v\n", clearRanks)
	for i := 0; i < len(player); i++ {
		for j := 0; j < len(clearRanks); j++ {
			if player[i] == clearRanks[j] {
				output = append(output, int32(j)+1)
				break
			}
			if player[i] > clearRanks[j] {
				output = append(output, int32(j)+1)
				//Need to insert player[i] into ranked slice at jth location
				slices.Insert(clearRanks, j, player[i])
				//clearRanks = append(clearRanks, player[i])
				//slices.Sort(clearRanks)
				fmt.Printf("inserted %v at %v into clearRanks: %v\n", player[i], j, clearRanks)
				break
			}
			if player[i] < clearRanks[j] && j == len(clearRanks)-1 {
				clearRanks = append(clearRanks, player[i])
				output = append(output, int32(len(clearRanks)))
				break
			}
		}
		//fmt.Printf("output after each player : %v\n", output)
	}
	fmt.Printf("output--->>>> : %v\n", output)
	return output
}

func camelcase(s string) int32 {
	// Write your code here
	words := 1
	for _, v := range s {
		if unicode.IsUpper(v) {
			words++
		}
	}
	return int32(words)
}
func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	max := 0
	for _, v := range height {
		if v > int32(max) {
			max = int(v)
		}
	}
	if k > int32(max) {
		return 0
	}
	return int32(max) - k
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	l := len(word)
	tall := 0
	for _, w := range word {
		low := strings.ToLower(string(w))
		i := strings.Index(alpha, low)
		if h[i] > int32(tall) {
			tall = int(h[i])
		}
	}
	area := tall * l
	return int32(area)

}

func utopianTree(n int32) int32 {
	// Write your code here
	var height int32 = 1
	var x int32 = 0
	for {
		if x == n {
			break
		}
		x++
		if x%2 != 0 {
			//Odd -> double
			height = height * 2
		} else {
			//Even. Increment
			height++
		}
	}
	return height

}

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var threshold int32 = 0
	for _, v := range a {
		if v <= 0 {
			threshold++
		}
	}
	if threshold >= k {
		return "NO"
	}
	return "YES"
}

func regularProfessor() {
	p := func(i int) int {
		return i * i
	}

	x := p(4)
	fmt.Printf("x: %v\n", x)

}
func viralAdvertising(n int32) int32 {
	// Write your code here
	var shared, liked, cum int32 = 5, 2, 2
	for i := 2; i <= int(n); i++ {
		shared = liked * 3
		liked = shared / 2
		cum = cum + liked
	}
	return cum

}
func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	// n = number of candies
	//m/n = num of rotations
	//m%n = where to stop
	//s = starting point
	var out int32 = 0
	rotations := m / n
	stop := m % n
	start := s

	if rotations == 0 {
		out = start + m - 1
	} else {
		if start == 1 {
			out = n
		} else {
			out = start - 1 + stop
		}
	}
	if out > n {
		out = out - n
	}
	return out
}
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	out := make([]int32, 0)
	start := 0
	rotated := make([]int32, 0)
	if len(a)-int(k) > 0 {
		start = len(a) - int(k)
	} else {
		start = len(a) - len(a)%int(k)
	}
	for i := start; i < len(a); i++ {
		rotated = append(rotated, a[i])
	}
	for i := 0; i < start; i++ {
		rotated = append(rotated, a[i])
	}
	for _, v := range queries {
		out = append(out, rotated[v])
	}
	fmt.Printf("Queries out: %v\n", out)
	return out
}
