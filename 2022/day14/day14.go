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

func createKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func part1(input string) int {

	occupied := make(map[string]bool)

	minX := 1000
	minY := 1000
	maxX := -1
	maxY := -1

	for _, l := range strings.Split(input, "\n") {

		previousX := 0
		previousY := 0
		s := strings.Split(l, " -> ")
		for i, coord := range s {
			xy := strings.Split(coord, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}

			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}

			if i > 0 {

				startX := previousX
				endX := x
				if previousX > x {
					startX = x
					endX = previousX
				}

				startY := previousY
				endY := y
				if previousY > y {
					startY = y
					endY = previousY
				}

				for j := startX; j <= endX; j++ {
					for k := startY; k <= endY; k++ {
						occupied[createKey(j, k)] = true
					}
				}

			}

			previousX = x
			previousY = y

		}

	}

	result := 0
	sandX := 500
	sandY := 0
	for {

		if sandX < minX || maxX < sandX || maxY < sandY {
			return result
		}

		downKey := createKey(sandX, sandY+1)
		leftKey := createKey(sandX-1, sandY+1)
		rightKey := createKey(sandX+1, sandY+1)

		if !occupied[downKey] {
			sandY += 1
		} else if !occupied[leftKey] {
			sandX -= 1
			sandY += 1
		} else if !occupied[rightKey] {
			sandX += 1
			sandY += 1
		} else {
			// fmt.Println(sandX, sandY)

			occupied[createKey(sandX, sandY)] = true
			result += 1
			sandX = 500
			sandY = 0
		}
	}

	return 1
}

func part2(input string) int {
	occupied := make(map[string]bool)

	minX := 1000
	minY := 1000
	maxX := -1
	maxY := -1

	for _, l := range strings.Split(input, "\n") {

		previousX := 0
		previousY := 0
		s := strings.Split(l, " -> ")
		for i, coord := range s {
			xy := strings.Split(coord, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])

			if x < minX {
				minX = x
			} else if x > maxX {
				maxX = x
			}

			if y < minY {
				minY = y
			} else if y > maxY {
				maxY = y
			}

			if i > 0 {

				startX := previousX
				endX := x
				if previousX > x {
					startX = x
					endX = previousX
				}

				startY := previousY
				endY := y
				if previousY > y {
					startY = y
					endY = previousY
				}

				for j := startX; j <= endX; j++ {
					for k := startY; k <= endY; k++ {
						occupied[createKey(j, k)] = true
					}
				}

			}

			previousX = x
			previousY = y

		}

	}

	maxY += 2
	result := 0
	sandX := 500
	sandY := 0
	for {

		downKey := createKey(sandX, sandY+1)
		leftKey := createKey(sandX-1, sandY+1)
		rightKey := createKey(sandX+1, sandY+1)

		if !occupied[downKey] && sandY+1 < maxY {
			sandY += 1
		} else if !occupied[leftKey] && sandY+1 < maxY {
			sandX -= 1
			sandY += 1
		} else if !occupied[rightKey] && sandY+1 < maxY {
			sandX += 1
			sandY += 1
		} else {
			occupied[createKey(sandX, sandY)] = true
			result += 1

			if sandX == 500 && sandY == 0 {
				return result
			}

			sandX = 500
			sandY = 0
		}
	}

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

	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
}
