package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func intersect(s1, s2 map[string]bool) map[string]bool {
	intersection := map[string]bool{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] {
			intersection[k] = true
		}
	}
	return intersection
}

// 2018/day01
func part1(input string) int {

	sum := 0
	for _, l := range strings.Split(input, "\n") {
		println(l)
		numbers := strings.Split(strings.Split(l, ":")[1], "|")
		winningNumbers := numbers[0]
		elfNumbers := numbers[1]

		winningNumbersSet := make(map[string]bool)
		for _, n := range strings.Split(winningNumbers, " ") {
			winningNumbersSet[n] = true
		}

		elfNumbersSet := make(map[string]bool)
		for _, n := range strings.Split(elfNumbers, " ") {
			elfNumbersSet[n] = true
		}

		intersection := intersect(winningNumbersSet, elfNumbersSet)
		numberOfWinning := len(intersection) - 2
		println(numberOfWinning)
		i := int(math.Pow(2, float64(numberOfWinning)))
		println(i)
		sum += i

	}
	return sum
}

func part2(input string) int {

	sum := 0
	sMap := make(map[int]int)
	lines := strings.Split(input, "\n")
	for index, l := range lines {
		println(l)
		numbers := strings.Split(strings.Split(l, ":")[1], "|")
		winningNumbers := numbers[0]
		elfNumbers := numbers[1]

		winningNumbersSet := make(map[string]bool)
		for _, n := range strings.Split(winningNumbers, " ") {
			winningNumbersSet[n] = true
		}

		elfNumbersSet := make(map[string]bool)
		for _, n := range strings.Split(elfNumbers, " ") {
			elfNumbersSet[n] = true
		}

		intersection := intersect(winningNumbersSet, elfNumbersSet)
		numberOfWinning := len(intersection) - 1

		sMap[index] += 1
		println("winning:", numberOfWinning, "index:", index)
		for i := 0; i < numberOfWinning; i++ {
			x := index + i + 1
			sMap[x] += sMap[index]
		}

	}
	for k, v := range sMap {
		println(k+1, v)
		sum += v
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
