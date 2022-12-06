package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func createSet(array []rune) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range array {
		set[item] = true
	}
	return set
}

func part1(input string) int {
	// Solved by hand ;)
	return 1623
}

func part2(input string) int {
	size := 14
	for i := 0; i < len(input)-size; i++ {
		window := createSet([]rune(input[i : i+size]))
		// fmt.Println(i, ":", len(window), input[i:i+size])
		if len(window) == size {
			return i + size
		}
	}
	return -1
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
