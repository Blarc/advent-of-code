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

func show(b, w map[string]bool, min, max []int) {
	for y := min[0] - 1; y <= max[0]+1; y++ {
		for x := min[1] - 1; x <= max[1]+1; x++ {
			key := createKey(y, x)
			if w[key] {
				fmt.Print("#")
			} else if b[key] {
				fmt.Print("b")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Blizzard struct {
	y int
	x int
	// up, down, left, right
	dir int
}

var blizzardMoves = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var moves = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{0, 0},
}

func moveBlizzards(blizzards []Blizzard, min, max []int) map[string]bool {
	newPositions := make(map[string]bool)

	for i := 0; i < len(blizzards); i++ {
		blizzard := blizzards[i]
		newY := blizzard.y + blizzardMoves[blizzard.dir][0]
		newX := blizzard.x + blizzardMoves[blizzard.dir][1]

		// If blizzard hits the bottom wall
		if newY > max[0] {
			newY = min[0]
		}
		// If blizzard hits the top wall
		if newY < min[0] {
			newY = max[0]
		}
		// If blizzard hits the right wall
		if newX > max[1] {
			newX = min[1]
		}
		// If blizzard hits the left wall
		if newX < min[1] {
			newX = max[1]
		}

		blizzards[i].y = newY
		blizzards[i].x = newX

		newPositions[createKey(newY, newX)] = true

	}

	return newPositions
}

func part1(input string) int {

	lines := strings.Split(input, "\n")
	walls := map[string]bool{"-1,1": true}
	blizzards := []Blizzard{}
	blizzardPositions := map[string]bool{}

	for y := 0; y < len(lines); y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			key := createKey(y, x)
			if line[x] == '#' {
				walls[key] = true
			} else if line[x] == '^' {
				blizzards = append(blizzards, Blizzard{y, x, 0})
				blizzardPositions[key] = true
			} else if line[x] == 'v' {
				blizzards = append(blizzards, Blizzard{y, x, 1})
				blizzardPositions[key] = true
			} else if line[x] == '<' {
				blizzards = append(blizzards, Blizzard{y, x, 2})
				blizzardPositions[key] = true
			} else if line[x] == '>' {
				blizzards = append(blizzards, Blizzard{y, x, 3})
				blizzardPositions[key] = true
			}
		}
	}

	min := []int{1, 1}
	max := []int{len(lines) - 2, len(lines[0]) - 2}

	start := []int{0, 1, 0}
	end := []int{len(lines) - 1, len(lines[0]) - 2}

	queue := [][]int{start}

	time := 0
	for {
		// Pop
		pos := queue[0]
		queue = queue[1:]

		// fmt.Println(pos)

		// Made it to the end
		if pos[0] == end[0] && pos[1] == end[1] {
			fmt.Println("Made it to the end in", pos[2], "minutes.")
			break
		}

		// Move blizzards
		if pos[2] == time {
			fmt.Println("Minute", time)
			// show(blizzardPositions, walls, min, max)
			blizzardPositions = moveBlizzards(blizzards, min, max)
			time++
		}

		for i := 0; i < len(moves); i++ {
			newY := pos[0] + moves[i][0]
			newX := pos[1] + moves[i][1]
			newTime := pos[2] + 1

			key := createKey(newY, newX)
			if !walls[key] && !blizzardPositions[key] {
				queue = append(queue, []int{newY, newX, newTime})
			}
		}
	}

	return 1
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
