package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func part1(input string) int {
	maximum := -1
	current := 0
	for _, l := range strings.Split(input, "\n") {

		if len(l) > 0 {
			number, _ := strconv.Atoi(l)
			current += number
		} else {
			if current > maximum {
				maximum = current
			}
			current = 0
		}
	}
	return maximum
}

func part2(input string) int {
	var maximums [3]int
	current := 0
	for _, l := range strings.Split(input, "\n") {

		if len(l) > 0 {
			number, _ := strconv.Atoi(l)
			current += number
		} else {
			for i, maximum := range maximums {
				if current > maximum {
					for j := len(maximums) - 1; j > i; j-- {
						maximums[j] = maximums[j-1]
					}
					maximums[i] = current
					break
				}
			}
			// fmt.Println(maximums)
			current = 0
		}
	}
	fmt.Println(maximums)

	result := 0
	for _, v := range maximums {
		result += v
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
		inputText = input
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = sample
		fmt.Println("Running part", part, "on sample.txt.")
	}

	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
}
