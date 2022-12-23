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

func createKey(y, x int) string {
	return fmt.Sprintf("%d,%d", y, x)
}

func show(elfs map[string]Elf, maxY, maxX, minY, minX int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, exists := elfs[createKey(y, x)]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Elf struct {
	y           int
	x           int
	proposedY   int
	proposedX   int
	hasProposal bool
}

func findAdjacent(elf Elf, elfs map[string]Elf) ([][]int, bool) {
	r := make([][]int, 12)
	i := 0
	e := false
	for y := elf.y - 1; y <= elf.y+1; y++ {
		for x := elf.x - 1; x <= elf.x+1; x++ {
			if y != elf.y || x != elf.x {
				if _, exists := elfs[createKey(y, x)]; exists {
					r[i] = []int{}
					e = true
				} else {
					r[i] = []int{y, x}
				}
			}
			i++
		}
	}

	// north, south, west, east

	// 0, 1, 2
	// 3, 4, 5
	// 6, 7, 8

	r = [][]int{r[0], r[1], r[2], r[6], r[7], r[8], r[0], r[3], r[6], r[2], r[5], r[8]}

	return r, e
}

func findProposal(dir int, adjacent [][]int) []int {
	// north, south, west, east
	for i := 0; i < 4; i++ {
		// Check if direction is free
		if len(adjacent[dir%12-1]) > 0 && len(adjacent[dir%12]) > 0 && len(adjacent[dir%12+1]) > 0 {
			return adjacent[dir%12]
		}
		dir += 3
	}
	return []int{}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	elfs := make(map[string]Elf)

	minY, minX := 0, 0
	maxY, maxX := len(lines), len(lines[0])

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				elfs[createKey(y, x)] = Elf{y, x, y, x, false}
			}
		}
	}

	// show(elfs, maxY, maxX, minY, minX)

	globalDir := 1
	for i := 0; i < 10; i++ {

		minY, minX = 1000, 1000
		maxY, maxX = -1000, -1000

		proposals := make(map[string]int)

		for k, elf := range elfs {
			adjacent, exists := findAdjacent(elf, elfs)
			if exists {
				// fmt.Println("elf", elf.y, elf.x)
				proposal := findProposal(globalDir, adjacent)
				if len(proposal) > 0 {
					// fmt.Println(elf.y, elf.x, "proposes", proposal[0], proposal[1])
					elf.proposedY = proposal[0]
					elf.proposedX = proposal[1]
					elf.hasProposal = true
					proposals[createKey(proposal[0], proposal[1])] += 1
				} else {
					elf.hasProposal = false
				}
			}

			elfs[k] = elf
		}

		newElfs := make(map[string]Elf)
		for _, elf := range elfs {
			if elf.hasProposal {
				p, pe := proposals[createKey(elf.proposedY, elf.proposedX)]
				if pe && p == 1 {
					elf.y = elf.proposedY
					elf.x = elf.proposedX
					elf.hasProposal = false
				}
			}

			if elf.y > maxY {
				maxY = elf.y
			}
			if elf.y < minY {
				minY = elf.y
			}

			if elf.x > maxX {
				maxX = elf.x
			}

			if elf.x < minX {
				minX = elf.x
			}

			newElfs[createKey(elf.y, elf.x)] = elf

		}

		elfs = newElfs

		globalDir += 3
		fmt.Printf("Round %d\n", i+1)
		show(elfs, maxY, maxX, minY, minX)
	}

	// show(elfs, maxY, maxX, minY, minX)

	// 3882
	// 1116

	return (maxX-minX+1)*(maxY-minY+1) - len(elfs)
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
