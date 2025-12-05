package main

import (
	"flag"
	"fmt"
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

func part1(input string) int {
	result := 0
	ranges := make([][]int, 0)

	readRanges := true
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			readRanges = false
		}

		if readRanges {
			s := strings.Split(l, "-")
			from, _ := strconv.Atoi(s[0])
			to, _ := strconv.Atoi(s[1])
			ranges = append(ranges, []int{from, to})

		} else {
			ingredient, _ := strconv.Atoi(l)

			for _, r := range ranges {
				if ingredient >= r[0] && ingredient <= r[1] {
					result++
					break
				}
			}
		}
	}
	return result
}

func expand(from, to int, r [][]int) [][]int {
	newRange := []int{from, to}
	merged := false

	for i := 0; i < len(r); i++ {
		if r[i][0] <= newRange[1] && newRange[0] <= r[i][1] {
			newRange[0] = min(newRange[0], r[i][0])
			newRange[1] = max(newRange[1], r[i][1])
			merged = true
		}
	}

	if !merged {
		return append(r, newRange)
	}

	result := make([][]int, 0)
	for i := 0; i < len(r); i++ {
		if newRange[1] < r[i][0] || r[i][1] < newRange[0] {
			result = append(result, r[i])
		}
	}
	result = append(result, newRange)
	return result
}

func part2(input string) int {
	result := 0
	ranges := make([][]int, 0)

	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			break
		}

		s := strings.Split(l, "-")
		from, _ := strconv.Atoi(s[0])
		to, _ := strconv.Atoi(s[1])
		ranges = expand(from, to, ranges)
	}

	for _, r := range ranges {
		result += r[1] - r[0] + 1
		// fmt.Printf("%d-%d: %d\n", r[0], r[1], r[1]-r[0]+1)
	}

	return result
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var inputText string
	if *inputPtr {
		inputText = strings.TrimSpace(input)
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = strings.TrimSpace(sample)
		fmt.Println("Running part", part, "on sample.txt.")
	}

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
