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

func checkNeighbours(point complex128, grid map[complex128]bool) bool {
	y := int(real(point))
	x := int(imag(point))

	neighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if grid[complex(float64(y+i), float64(x+j))] {
					neighbors++
				}
			}
		}
	}

	return neighbors < 4
}

func part1(input string) int {
	sum := 0
	grid := make(map[complex128]bool)

	for y, l := range strings.Split(input, "\n") {
		for x, c := range strings.Split(l, "") {
			if c == "@" {
				grid[complex(float64(y), float64(x))] = true
			}
		}
	}

	for point := range grid {
		if checkNeighbours(point, grid) {
			sum++
		}
	}

	return sum
}

func part2(input string) int {
	grid := make(map[complex128]bool)

	for y, l := range strings.Split(input, "\n") {
		for x, c := range strings.Split(l, "") {
			if c == "@" {
				grid[complex(float64(y), float64(x))] = true
			}
		}
	}

	result := 0
	prevSum := 0
	for {
		sum := 0
		newGrid := make(map[complex128]bool)
		for k, v := range grid {
			newGrid[k] = v
		}

		for point := range grid {
			if checkNeighbours(point, grid) {
				delete(newGrid, point)
				sum++
			}
		}

		grid = newGrid
		// fmt.Printf("%d\n", sum)
		result += sum
		if sum == prevSum {
			break
		}
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
