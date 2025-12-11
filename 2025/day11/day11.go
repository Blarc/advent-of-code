package main

import (
	"flag"
	"fmt"
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

func countPaths(node string, goal string, visited map[string]bool, nodes map[string][]string, memo map[string]int) int {
	if node == goal {
		memo[node] = 1
		return 1
	}

	if v, ok := memo[node]; ok {
		return v
	}

	count := 0
	for _, n := range nodes[node] {
		if !visited[n] {
			visited[n] = true
			count += countPaths(n, goal, visited, nodes, memo)
			visited[n] = false
		}
	}

	memo[node] = count
	return count
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	nodes := make(map[string][]string)
	for _, l := range lines {
		s := strings.Split(l, ":")
		nodes[s[0]] = strings.Fields(s[1])
	}
	return countPaths("you", "out", make(map[string]bool), nodes, make(map[string]int))
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	nodes := make(map[string][]string)
	for _, l := range lines {
		s := strings.Split(l, ":")
		nodes[s[0]] = strings.Fields(s[1])
	}

	return countPaths("svr", "fft", make(map[string]bool), nodes, make(map[string]int))*
		countPaths("fft", "dac", make(map[string]bool), nodes, make(map[string]int))*
		countPaths("dac", "out", make(map[string]bool), nodes, make(map[string]int)) +
		countPaths("svr", "dac", make(map[string]bool), nodes, make(map[string]int))*
			countPaths("dac", "fft", make(map[string]bool), nodes, make(map[string]int))*
			countPaths("fft", "out", make(map[string]bool), nodes, make(map[string]int))
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
