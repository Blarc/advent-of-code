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

func isVisible(xc int, yc int, grid [][]int) bool {
	left := true
	for x := xc - 1; x >= 0; x-- {
		if grid[yc][x] >= grid[yc][xc] {
			left = false
			break
		}
	}

	right := true
	for x := xc + 1; x < len(grid); x++ {
		if grid[yc][x] >= grid[yc][xc] {
			right = false
			break
		}
	}

	up := true
	for y := yc - 1; y >= 0; y-- {
		if grid[y][xc] >= grid[yc][xc] {
			up = false
			break
		}
	}

	down := true
	for y := yc + 1; y < len(grid); y++ {
		if grid[y][xc] >= grid[yc][xc] {
			down = false
			break
		}
	}

	return left || right || up || down

}

func isVisible2(xc int, yc int, grid [][]int) int {
	left := true
	var x int
	for x = xc - 1; x >= 0; x-- {
		if grid[yc][x] >= grid[yc][xc] {
			left = false
			break
		}
	}

	if x < 0 {
		x = 0
	}

	leftX := xc - x

	right := true
	for x = xc + 1; x < len(grid); x++ {
		if grid[yc][x] >= grid[yc][xc] {
			right = false
			break
		}
	}

	if x >= len(grid) {
		x = len(grid) - 1
	}

	rightX := x - xc

	up := true
	var y int
	for y = yc - 1; y >= 0; y-- {
		if grid[y][xc] >= grid[yc][xc] {
			up = false
			break
		}
	}

	if y < 0 {
		y = 0
	}

	upY := yc - y

	down := true
	for y = yc + 1; y < len(grid); y++ {
		if grid[y][xc] >= grid[yc][xc] {
			down = false
			break
		}
	}

	if y >= len(grid) {
		y = len(grid) - 1
	}

	downY := y - yc

	if left || right || up || down {
		return leftX * rightX * upY * downY
	}
	return -1

}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	size := len(lines[0])

	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	for y, l := range lines {
		for x, number := range strings.Split(l, "") {
			numberInt, _ := strconv.Atoi(number)
			grid[y][x] = numberInt
		}
	}

	result := 0
	for y := 1; y < size-1; y++ {
		for x := 1; x < size-1; x++ {
			visible := isVisible(x, y, grid)

			// fmt.Println(x, y, visible)
			if visible {
				result += 1
			}
		}
	}

	return 4*size - 4 + result
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	size := len(lines[0])

	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	for y, l := range lines {
		for x, number := range strings.Split(l, "") {
			numberInt, _ := strconv.Atoi(number)
			grid[y][x] = numberInt
		}
	}

	result := 0
	for y := 1; y < size-1; y++ {
		for x := 1; x < size-1; x++ {
			visible := isVisible2(x, y, grid)

			// fmt.Println(x, y, visible)
			if visible > result {
				result = visible
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
