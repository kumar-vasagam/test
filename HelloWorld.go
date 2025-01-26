package main

import (
	"container/list"
	"fmt"
	"strconv"
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
	seen := nums[0]
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != seen {
			j++
			nums[j] = nums[i]
			seen = nums[i]
		}
	}
	fmt.Printf("j: %v\n", j)
	return j + 1
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
func linkedList() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertAfter(2, e1)
	l.InsertBefore(3, e4)

	fmt.Printf("l: %v\n", l)
	// for e := l.Front(); e != nil; e.Next() {
	// 	fmt.Printf("e.Value: %v\n", e.Value)
	// }

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func timeInWords(h, m int32) string {
	hour := int(h)
	mins := int(m)

	hours := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six",
		7: "seven", 8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve",
		13: "thirteen", 14: "fourteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
		20: "twenty", 21: "twenty one", 22: "twenty two", 23: "twenty three", 24: "twenty four",
		25: "twenty five", 26: "twenty six", 27: "twenty seven", 28: "twenty eight", 29: "twenty nine",
	}
	skips := map[int]string{15: "quarter", 30: "half"}
	switch mins {
	case 0:
		return hours[hour] + " o' " + "clock"
	case 15:
		return skips[mins] + " past " + hours[hour]
	case 30:
		return skips[mins] + " past " + hours[hour]
	case 45:
		return skips[60-mins] + " to " + hours[hour+1]
	}

	switch {
	case mins == 1:
		return hours[mins] + " minute past " + hours[hour]
	case mins > 1 && mins < 30:
		return hours[mins] + " minutes past " + hours[hour]
	default:
		return hours[60-mins] + " minutes to " + hours[hour+1]
	}
}

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	p_ := int(p)
	q_ := int(q)
	var sb strings.Builder
	for i := p_; i <= q_; i++ {
		numDigits := strconv.Itoa(int(i))

		//Square it
		j := i * i
		//Convert into a string
		jaas := strconv.Itoa(int(j))
		newlgt := len(jaas) - len(numDigits)
		left := jaas[:newlgt]
		right := jaas[newlgt:]

		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)

		if l+r == int(i) {
			//fmt.Printf("left : %v | right : %v\n", left, right)
			//fmt.Printf("l : %v | r: %v\n", l, r)
			sb.WriteString(numDigits + " ")
		}

	}
	if sb.String() == "" {
		fmt.Print("INVALID RANGE")
	} else {
		fmt.Print(strings.TrimSpace(sb.String()))
	}

}

func ReverseStringAnotherWay(s string) string {
	trimmed := strings.TrimSpace(s)
	ss := strings.Split(trimmed, " ")
	var sb strings.Builder
	for i := len(ss) - 1; i >= 0; i-- {
		t_item := strings.TrimSpace(ss[i])
		if t_item != "" {
			sb.WriteString(ss[i] + " ")
		}
	}
	fmt.Printf(" \n Reversed ? %v \n", sb.String())
	return strings.TrimSpace(sb.String())
}
func TwoDSlices() {
	holder := make([][]int, 5)
	fmt.Printf("len %v | cap %v | Holder %v \n", len(holder), cap(holder), holder)
	fmt.Printf("Len of 0th %v \n", len(holder[0]))

	subHolder := make([]int, 3)
	holder[0] = append(holder[0], subHolder...)

	fmt.Printf("Len of 0th after subHolder %v | Hodler %v \n", len(holder[0]), holder)

	secondHolder := []int{0, 1, 2, 3, 4}
	holder[1] = append(holder[1], secondHolder...)

	fmt.Printf("Len of 1st after secondHolder %v | Hodler %v \n", len(holder[1]), holder)

	fmt.Printf("Element at 1,2 : %v \n", holder[1][2])

}
func convert(s string, numRows int) string {
	if numRows >= len(s) || numRows == 1 {
		return s
	}
	holder := make([][]string, 0)
	idx := numRows - 2
	s1 := s[0:numRows]
	temp := make([]string, numRows)
	for i, s11 := range s1 {
		temp[i] = string(s11)
	}
	cPos := numRows
	//Init
	holder = append(holder, temp)
	for {
		subHolder := make([]string, numRows)
		if idx == 0 {
			//Hydrate all
		Inner1:
			for i := 0; i < numRows; i++ {
				subHolder[i] = s[cPos : cPos+1]
				cPos++
				if cPos >= len(s) {
					break Inner1
				}
			}
			holder = append(holder, subHolder)
		} else {
			//Hydrate only that position
		Inner2:
			for i := 0; i < numRows-1; i++ {
				if i == idx {
					subHolder[i] = s[cPos : cPos+1]
					cPos++
					if cPos >= len(s) {
						break Inner2
					}
				} else {
					subHolder[i] = ""
				}
			}
			holder = append(holder, subHolder)
		}
		if idx == 0 && cPos < len(s) {
			//Reset idx
			idx = numRows - 2
			continue
		}
		idx--
		//cPos++
		if cPos >= len(s) {
			break
		}
	}

	//Loop thru the subHolders and pick same spot letters and smush them together for the final word!
	var sb strings.Builder
	max := len(holder[0])
	for j := 0; j < max; j++ {
		for i := 0; i < len(holder); i++ {
			item := holder[i][j]
			if item != "" {
				sb.WriteString(item)
			}
		}
	}
	fmt.Printf("sb.String(): %v\n", sb.String())
	return sb.String()
}

func main() {
	result := add(3, 7)
	fmt.Println(result)
	findSum()
	reverseString("I am  a man'")
	reverseString(" Hello  World   ")

	fmt.Printf("Want 2 | Got %v  \n", removeDuplicates([]int{1, 1, 2}))
	fmt.Printf("Want 5 | Got %v  \n", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Printf("Want 20 | Got %v : \n", taumBday(10, 10, 1, 1, 1))
	fmt.Printf("Want 37 | Got %v : \n", taumBday(5, 5, 2, 3, 4))
	fmt.Printf("Want 12 | Got %v : \n", taumBday(3, 6, 9, 1, 1))
	fmt.Printf("Want 35 | Got %v : \n", taumBday(7, 7, 4, 2, 1))
	fmt.Printf("Want 12 | Got %v : \n", taumBday(3, 3, 1, 9, 2))
	fmt.Printf("Want 18192035842 | Got %v : \n", taumBday(27984, 1402, 619246, 615589, 247954))
	linkedList()
	fmt.Printf("Want five o' clock | Got %q \n", timeInWords(5, 0))
	fmt.Printf("Want one minute past 5 | Got %q \n", timeInWords(5, 1))
	fmt.Printf("Want ten minutes past five | Got %q \n", timeInWords(5, 10))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 15))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 30))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 40))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 45))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 47))
	fmt.Printf("Want  | Got %q \n", timeInWords(5, 28))
	fmt.Println("----Kaprekar----")
	kaprekarNumbers(1, 99999)

	ReverseStringAnotherWay("the sky is blue")
	ReverseStringAnotherWay("  hello world  ")
	ReverseStringAnotherWay("a good   example")
	fmt.Println("----Convert----")
	convert("PAYPALISHIRING", 3)
	convert("PAYPALISHIRING", 4)
	convert("A", 2)
	convert("AB", 2)
	//convert("PAYPALISHIRING", 5)

	TwoDSlices()

}
