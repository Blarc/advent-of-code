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
	result := 0

	lines := strings.Split(input, "\n")
	numbers := make([][]int, len(strings.Fields(lines[0])))
	for _, l := range lines[:len(lines)-1] {
		for i, v := range strings.Fields(l) {
			n, _ := strconv.Atoi(strings.TrimSpace(v))
			numbers[i] = append(numbers[i], n)
		}
	}

	for i, op := range strings.Fields(lines[len(lines)-1]) {
		sum := numbers[i][0]
		for j := 1; j < len(numbers[i]); j++ {
			if op == `+` {
				sum += numbers[i][j]
			} else if op == `*` {
				sum *= numbers[i][j]
			} else {
				panic("Invalid operation!")
			}
		}
		result += sum
	}

	return result
}

func rotate90(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	width := len(lines[0])
	height := len(lines)
	rotated := make([]string, width)

	for col := 0; col < width; col++ {
		var builder strings.Builder
		for row := height - 1; row >= 0; row-- {
			builder.WriteByte(lines[row][col])
		}
		rotated[col] = builder.String()
	}

	return rotated
}

func part2(input string) int {
	result := 0

	lines := strings.Split(input, "\n")
	numberLines := lines[:len(lines)-1]
	operations := strings.Fields(lines[len(lines)-1])

	index := 0
	sum := 0
	rotated := rotate90(rotate90(rotate90(numberLines)))
	rotated = append(rotated, "\n")

	for _, nl := range rotated {
		if strings.TrimSpace(nl) == "" {
			// fmt.Printf("sum: %d\n", sum)
			index++
			result += sum
			sum = 0
			continue
		}

		number, _ := strconv.Atoi(strings.TrimSpace(nl))
		if sum == 0 {
			sum = number
		} else {
			if operations[len(operations)-1-index] == `+` {
				sum += number
			} else if operations[len(operations)-1-index] == `*` {
				sum *= number
			} else {
				panic("Invalid operation!")
			}
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
