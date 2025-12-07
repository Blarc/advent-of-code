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

type Point struct {
	X, Y int
}

func part1(input string) int {
	result := 0

	start := Point{}
	splitters := make(map[Point]bool)

	lines := strings.Split(input, "\n")
	for y, l := range lines {
		for x, c := range l {
			if c == 'S' {
				start = Point{x, y}
			} else if c == '^' {
				splitters[Point{x, y}] = true
			}
		}
	}

	queue := []Point{start}
	visited := make(map[Point]bool)

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		next := Point{pos.X, pos.Y + 1}

		if next.Y == len(lines)-1 {
			break
		}

		if _, ok := splitters[next]; ok {
			fmt.Printf("Found splitter at (%d, %d)\n", next.X, next.Y)
			result++
			left := Point{next.X - 1, next.Y}
			if !visited[left] {
				queue = append(queue, left)
				visited[left] = true
			}

			right := Point{next.X + 1, next.Y}
			if !visited[right] {
				queue = append(queue, right)
				visited[right] = true
			}
		} else {
			if !visited[next] {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}

	return result
}

func countTimelines(pos Point, n int, splitters map[Point]bool, memo map[Point]int) int {
	next := Point{pos.X, pos.Y + 1}
	if next.Y == n {
		memo[pos] = 1
		return 1
	}

	if v, ok := memo[next]; ok {
		return v
	}

	result := 0
	if _, ok := splitters[next]; ok {
		left := Point{next.X - 1, next.Y}
		result += countTimelines(left, n, splitters, memo)

		right := Point{next.X + 1, next.Y}
		result += countTimelines(right, n, splitters, memo)
	} else {
		result += countTimelines(next, n, splitters, memo)
	}

	memo[pos] = result
	return result
}

func part2(input string) int {
	result := 0

	start := Point{}
	splitters := make(map[Point]bool)

	lines := strings.Split(input, "\n")
	for y, l := range lines {
		for x, c := range l {
			if c == 'S' {
				start = Point{x, y}
			} else if c == '^' {
				splitters[Point{x, y}] = true
			}
		}
	}

	result = countTimelines(start, len(lines), splitters, make(map[Point]int))
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
