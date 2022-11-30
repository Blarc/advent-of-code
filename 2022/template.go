package main

import (
	"flag"
	"fmt"
)

func part1() int {
	return 1
}

func part2() int {
	return 2
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var input string = "sample.txt"
	if *inputPtr {
		input = "input.txt"
	}

	fmt.Println("Running part", part, "on", input)
	if part == 1 {
		fmt.Println("Result:", part1())
	} else {
		fmt.Println("Result:", part2())
	}
}
