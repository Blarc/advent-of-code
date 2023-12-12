package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

import (
	_ "embed"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func compare(a []string, cond []int) bool {

	var test []int
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == "#" {
			count += 1
		} else if count > 0 {
			test = append(test, count)
			count = 0
		}
	}
	if count > 0 {
		test = append(test, count)
	}

	return slices.Compare(test, cond) == 0
}

func checkCond(a []string, cond int) bool {
	for i := 0; i < cond; i++ {
		if a[i] != "#" && a[i] != "?" {
			return false
		}
	}
	return true
}

func findOptions(a []string, cond []int, memo map[string]int) int {
	memoKey := fmt.Sprintf("%v %v", a, cond)

	if val, ok := memo[memoKey]; ok {
		return val
	}

	//fmt.Println(a, cond)
	if len(cond) == 0 && len(a) == 0 {
		//fmt.Println("goal")
		memo[memoKey] = 1
		return 1
	}

	if len(cond) != 0 && len(a) < cond[0] {
		memo[memoKey] = 0
		return 0
	}
	sum := 0

	if len(cond) != 0 {
		if checkCond(a[:cond[0]], cond[0]) {
			if len(a) == cond[0] {
				sum += findOptions(a[cond[0]:], cond[1:], memo)
			} else if a[cond[0]] != "#" {
				sum += findOptions(a[cond[0]+1:], cond[1:], memo)
			}
		}
	}
	if a[0] != "#" {
		sum += findOptions(a[1:], cond, memo)
	}

	memo[memoKey] = sum
	return sum
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		s := strings.Split(l, " ")
		first := strings.Split(s[0], "")
		second := strings.Split(s[1], ",")
		conditions := make([]int, len(second))
		for i := 0; i < len(second); i++ {
			n, _ := strconv.Atoi(second[i])
			conditions[i] = n
		}
		o := findOptions(first, conditions, make(map[string]int))
		fmt.Println(o)
		sum += o
		//var qPosition []int
		//for i := 0; i < len(first); i++ {
		//	if first[i] == "?" {
		//		qPosition = append(qPosition, i)
		//	}
		//}
		//
		////fmt.Println(len(qPosition))
		//for i := 1; i < int(math.Pow(2, float64(len(qPosition)))); i++ {
		//	b := strings.Split(strconv.FormatInt(int64(i), 2), "")
		//	for j := len(b); j < len(qPosition); j++ {
		//		b = append([]string{"0"}, b...)
		//	}
		//	//fmt.Println(b)
		//
		//	newFirst := slices.Clone(first)
		//	for j := 0; j < len(b); j++ {
		//		if b[j] != "0" {
		//			newFirst[qPosition[j]] = "#"
		//		}
		//	}
		//	//fmt.Println(newFirst, compare(newFirst, conditions))
		//	if compare(newFirst, conditions) {
		//		//fmt.Println(newFirst)
		//		sum += 1
		//	}
		//}
	}
	// 7728 too low
	// 8288 too high
	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		var f []string
		var c []int
		s := strings.Split(l, " ")
		first := strings.Split(s[0], "")
		second := strings.Split(s[1], ",")
		conditions := make([]int, len(second))
		for i := 0; i < len(second); i++ {
			n, _ := strconv.Atoi(second[i])
			conditions[i] = n
		}

		for i := 0; i < 5; i++ {
			f = append(f, first...)
			if i < 4 {
				f = append(f, "?")
			}
			c = append(c, conditions...)
		}

		o := findOptions(f, c, make(map[string]int))
		fmt.Println(o)
		sum += o
	}
	return sum
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var inputText string
	if *inputPtr {
		inputText = input
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = sample
		fmt.Println("Running part", part, "on sample.txt.")
	}

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(strings.TrimSpace(inputText)))
	} else {
		fmt.Println("Result:", part2(strings.TrimSpace(inputText)))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
