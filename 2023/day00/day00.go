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

// 2018/day01
func part1(input string) int {
	sum := 0
	for _, l := range strings.Split(input, "\n") {

		var current, _ = strconv.Atoi(l)
		sum += current
	}
	return sum
}

func part2(input string) int {
	sum := 0
	seen := make(map[int]bool)

	input = strings.TrimSpace(input)
	inputArray := strings.Split(input, "\n")

	for {
		for _, l := range inputArray {
			var current, _ = strconv.Atoi(l)
			sum += current

			if seen[sum] {
				return sum
			}

			seen[sum] = true
		}
	}
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
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
