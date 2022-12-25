package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

var digits = map[byte]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var snafu = map[int]rune{
	4: '2',
	3: '1',
	2: '0',
	1: '-',
	0: '=',
}

func part1(input string) string {
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		number := 0
		power := 1
		for i := len(l) - 1; i >= 0; i-- {
			number += power * digits[l[i]]
			power *= 5
		}
		sum += number
	}

	fmt.Println(sum)
	base5 := []int{}
	for sum > 0 {
		sum += 2
		base5 = append([]int{sum % 5}, base5...)
		sum = sum / 5
	}

	fmt.Println(base5)

	var result string
	for i := 0; i < len(base5); i++ {
		result += string(snafu[base5[i]])
	}

	return result
}

func part2(input string) int {
	return 2
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
