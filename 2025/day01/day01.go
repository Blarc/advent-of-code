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
	count := 0
	dial := 50

	for _, l := range strings.Split(input, "\n") {

		dir := l[0]
		amount, _ := strconv.Atoi(l[1:])

		fmt.Printf("dir: %c amount: %d\n", dir, amount)

		if dir == 'R' {
			dial += amount % 100
		} else if dir == 'L' {
			dial -= amount % 100
		} else {
			panic("Wrong direction set!")
		}

		if dial < 0 {
			dial = 100 + dial
		} else if dial > 99 {
			dial = dial - 100
		}

		fmt.Printf("Dial: %d\n", dial)

		if dial == 0 {
			count++
		}
	}
	return count
}

func part2(input string) int {
	count := 0
	dial := 50
	prevDial := dial

	for _, l := range strings.Split(input, "\n") {

		dir := l[0]
		amount, _ := strconv.Atoi(l[1:])
		count += amount / 100
		amount %= 100

		fmt.Printf("dir: %c amount: %d\n", dir, amount)

		if dir == 'R' {
			dial += amount
		} else if dir == 'L' {
			dial -= amount
		} else {
			panic("Wrong direction set!")
		}

		over := false
		if dial < 0 {
			dial = 100 + dial
			if prevDial != 0 {
				over = true
			}
		} else if dial > 99 {
			dial = dial - 100
			if prevDial != 0 {
				over = true
			}
		}
		prevDial = dial

		fmt.Printf("Dial: %d\n Over: %v\n", dial, over)
		if dial == 0 {
			count++
		} else if over {
			count++
		}
	}
	return count
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
