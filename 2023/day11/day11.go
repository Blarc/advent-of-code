package main

import (
	"flag"
	"fmt"
	"math"
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

func manhattan(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	galaxies := make([][]int, 0)
	rows := make([]int, len(lines))
	cols := make([]int, len(lines[0]))

	for y, l := range lines {
		row := strings.Split(l, "")
		fmt.Println(row)
		for x, c := range row {
			if c == "." {
				rows[y]++
				cols[x]++
			}
			if c == "#" {
				galaxies = append(galaxies, []int{y, x})
			}
		}

	}
	fmt.Println("rows:", rows)
	fmt.Println("cols:", cols)

	all := 0
	sum := 0
	for _, galaxyA := range galaxies {
		for _, galaxyB := range galaxies {
			if galaxyA[0] != galaxyB[0] || galaxyA[1] != galaxyB[1] {
				dist := manhattan(galaxyA, galaxyB)

				for y := min(galaxyA[0], galaxyB[0]); y < max(galaxyA[0], galaxyB[0]); y++ {
					if rows[y] == len(rows) {
						dist += 1
					}
				}
				for x := min(galaxyA[1], galaxyB[1]); x < max(galaxyA[1], galaxyB[1]); x++ {
					if cols[x] == len(cols) {
						dist += 1
					}
				}
				//fmt.Println(indexA+1, indexB+1, dist)
				all += 1
				sum += dist
			}
		}
	}

	fmt.Println(all / 2)
	return sum / 2
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	galaxies := make([][]int, 0)
	rows := make([]int, len(lines))
	cols := make([]int, len(lines[0]))

	for y, l := range lines {
		row := strings.Split(l, "")
		for x, c := range row {
			if c == "." {
				rows[y]++
				cols[x]++
			}
			if c == "#" {
				galaxies = append(galaxies, []int{y, x})
			}
		}

	}

	all := 0
	sum := 0
	for _, galaxyA := range galaxies {
		for _, galaxyB := range galaxies {
			if galaxyA[0] != galaxyB[0] || galaxyA[1] != galaxyB[1] {
				dist := manhattan(galaxyA, galaxyB)

				for y := min(galaxyA[0], galaxyB[0]); y < max(galaxyA[0], galaxyB[0]); y++ {
					if rows[y] == len(rows) {
						dist += (1000000 - 1)
					}
				}
				for x := min(galaxyA[1], galaxyB[1]); x < max(galaxyA[1], galaxyB[1]); x++ {
					if cols[x] == len(cols) {
						dist += (1000000 - 1)
					}
				}
				//fmt.Println(indexA+1, indexB+1, dist)
				all += 1
				sum += dist
			}
		}
	}

	fmt.Println(all / 2)
	return sum / 2
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
		fmt.Println("Result:", part1(strings.TrimSpace(inputText)))
	} else {
		fmt.Println("Result:", part2(strings.TrimSpace(inputText)))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
