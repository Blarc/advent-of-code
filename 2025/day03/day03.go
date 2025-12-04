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

func toIntArray(s string) []int {
	var arr []int
	for _, v := range strings.Split(s, "") {
		i, _ := strconv.Atoi(v)
		arr = append(arr, i)
	}
	return arr
}

func part1(input string) int {
	sum := 0

	for _, l := range strings.Split(input, "\n") {
		r := toIntArray(l)
		leftMax := 0
		rightMax := 0

		for i := 0; i < len(r); i++ {
			if r[i] > leftMax && i < len(r)-1 {
				leftMax = r[i]
				rightMax = 0
			} else if r[i] > rightMax {
				rightMax = r[i]
			}
		}

		// fmt.Printf("%v\n", r)
		// fmt.Printf("Left: %d, Right: %d\n", leftMax, rightMax)
		sum += leftMax*10 + rightMax
	}
	return sum
}

func toNumber(a []int) int {
	var n int
	for _, v := range a {
		n = n*10 + v
	}
	return n
}

func part2(input string) int {
	// change to 2 and it works for part 1 as well
	const length = 12
	sum := 0

	for _, l := range strings.Split(input, "\n") {
		r := toIntArray(l)
		maxes := make([]int, length)

		maxJ := 0
		for i := 0; i < length; i++ {
			for j := maxJ; j < len(r); j++ {
				//fmt.Printf("%d,%d\n", i, j)
				//fmt.Printf("%d > %d\n", r[j], maxes[i])
				if r[j] > maxes[i] && j < len(r)-length+i+1 {
					maxes[i] = r[j]
					maxJ = j + 1
				}
			}
			//fmt.Printf("Max: %d\n", maxes[i])

		}

		//fmt.Printf("%v\n", r)
		//fmt.Printf("%v\n", maxes)
		sum += toNumber(maxes)
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
