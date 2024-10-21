package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Wanted : 2 | Got : %v \n", peakElement([]int{1, 2, 3}, 3))
	fmt.Printf("Wanted : 0 | Got : %v \n", peakElement([]int{1, 1, 1, 2, 1, 1}, 7))
	staircase(6)
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	birthdayCakeCandles([]int32{1, 2, 3, 3})
	timeConversion("07:05:45PM")
	timeConversion("07:05:45AM")
	out := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Printf("out: %v\n", out)
	out = breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42})
	fmt.Printf("out 2 : %v\n", out)

	out1 := birthdays([]int32{2, 2, 1, 3, 2}, 4, 2)
	fmt.Printf("out1 : %v\n", out1)

	out1 = birthdays([]int32{1, 2, 1, 3, 2}, 3, 2)
	fmt.Printf("out1 : %v\n", out1)

	out1 = birthdays([]int32{1, 1, 1, 1, 1, 1}, 3, 2)
	fmt.Printf("out1 : %v\n", out1)

	res := divisibleSumPairs(6, 5, []int32{1, 2, 3, 4, 5, 6})
	fmt.Printf("res: %v\n", res)

	res = divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	fmt.Printf("res: %v\n", res)

	res = migratoryBirds([]int32{1, 1, 2, 2, 3})
	fmt.Printf("res: %v\n", res)

	res = migratoryBirds([]int32{1, 4, 4, 4, 5, 3})
	fmt.Printf("res: %v\n", res)

	res = migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4})
	fmt.Printf("res: %v\n", res)

}

func peakElement(arr []int, n int) int {
	peak := arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > peak && arr[i] > arr[i+1] {
			peak = arr[i]
		} else {
			continue
		}
	}
	return peak
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	rows := len(arr[0])
	cols := len(arr)
	var principal, secondary int32
	//var int32 secondary
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j {
				principal += arr[i][j]
			}
			if j == cols-1-i {
				secondary += arr[i][j]
			}
		}
	}
	fmt.Printf("Pricipal %v | Secondary %v \n", principal, secondary)
	result := principal - secondary
	if result < 0 {
		return -result
	} else {
		return result
	}
}

func staircase(n int32) {
	// Write your code here
	var c int32 = 0
	for {
		for i := 0; i < int(n-c); i++ {
			fmt.Print(" ")
		}
		for j := 0; j < int(c); j++ {
			fmt.Print("#")
		}
		fmt.Println()
		c++
		if c > n {
			break
		}
	}

}

func miniMaxSum(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < arr[j] {
			return true
		} else {
			return false
		}
	})
	var minSum, maxSum int64
	for i, v := range arr {
		if i == 0 {
			minSum += int64(v)
		}
		if i == len(arr)-1 {
			maxSum += int64(v)
		}
		if i > 0 && i < len(arr)-1 {
			minSum += int64(v)
			maxSum += int64(v)
		}
	}

	fmt.Printf("%v  %v\n", minSum, maxSum)
}

func birthdayCakeCandles(candles []int32) int32 {
	var max, counter int32
	for _, v := range candles {
		if max == v {
			counter++
		}
		if v > max {
			max = v
		}
	}
	counter++
	fmt.Printf("counter: %v\n", counter)
	return counter
}

func timeConversion(s string) string {
	// Write your code here
	const layout = "01/02 03:04:05PM '06 -0700"
	const milForm = "15:04:05PM"
	t, _ := time.Parse(milForm, "17:34:14PM")
	fmt.Printf("t--->: %v\n", t)

	t, _ = time.Parse(time.Kitchen, "01/02 06:01:11AM")
	fmt.Printf("t: %v\n", t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	const shortForm1 = "2006-Jan-02 03:04:05AM"
	t, _ = time.Parse(shortForm1, "2013-Feb-03")
	fmt.Println(t)

	//String manipulation
	//12:00:00AM
	var sb strings.Builder
	var pm = false
	arr := strings.Split(s, ":")
	if strings.HasSuffix(arr[len(arr)-1], "PM") {
		pm = true
	}
	if pm {
		if arr[0] != "12" {
			h, _ := strconv.Atoi(arr[0])
			h += 12
			sb.WriteString(strconv.Itoa(h) + ":")
		} else {
			sb.WriteString(arr[0] + ":")
		}
	} else {
		if arr[0] == "12" {
			sb.WriteString("00" + ":")
		} else {
			sb.WriteString(arr[0] + ":")
		}
	}
	sb.WriteString(arr[1] + ":")
	sb.WriteString(arr[2][0:2])
	fmt.Printf("s----: %v\n", sb.String())
	return sb.String()
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var j int32
	j = (x2 - x1) / (v1 - v2)
	if (x2-x1)%(v1-v2) != 0 {
		return "NO"
	}
	if j < 0 {
		return "NO"
	}
	return "YES"
}
func breakingRecords(scores []int32) []int32 {
	// Write your code here
	min, max := scores[0], scores[0]
	counter := []int32{0, 0}

	for i := 1; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			counter[1]++
		}
		if scores[i] > max {
			max = scores[i]
			counter[0]++
		}
	}
	return counter
}

func birthdays(s []int32, d int32, m int32) int32 {
	// Write your code here
	// m = num of elements; d = sum of m elements;
	var i int32 = 0
	ways := 0
	for {
		if i+m-1 >= int32(len(s)) {
			break
		}
		var sum int32 = 0
		for j := i; j <= i+m-1; j++ {
			sum += s[j]
		}
		if sum == d {
			ways++
		}
		i++
	}
	return int32(ways)
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	birds := make(map[int32]int32, 0)

	for _, v := range arr {
		_, ok := birds[v]
		if ok {
			birds[v]++
		} else {
			birds[v] = 1
		}
	}

	var lowId, highSightings int32
	for k, v := range birds {
		if v > highSightings {
			highSightings = v
			lowId = k
		}
		if v == highSightings && k < lowId {
			lowId = k
		}
	}

	return lowId
}

func dayOfProgrammer(year int32) string {
	const c = 243 //Aug 31 on a leap year after 1918
	const pd = 256
	if year >= 1700 && year < 1918 {
		if year%4 == 0 {
			//leap year
			return "12.09." + fmt.Sprint(year)
		} else {
			return "13.09." + fmt.Sprint(year)
		}
	}
	if year > 1918 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			//leap year
			return "12.09." + fmt.Sprint(year)
		} else {
			return "13.09." + fmt.Sprint(year)
		}
	}

	return "26.09." + fmt.Sprint(year)
	//return

}

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	// k should not be included in calculation (position)
	var actualBill int32
	for i, v := range bill {
		if i != int(k) {
			actualBill += v
		}
	}
	eachContrib := actualBill / 2
	if eachContrib == b {
		fmt.Println("Bon Appetit")
	} else {
		refund := (actualBill+bill[k])/2 - b
		refund64 := math.Abs(float64(refund))
		fmt.Printf("refund: %v\n", refund64)
	}
}
