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

type Node struct {
	y int
	x int
	t string
}

func next(current *Node, previous *Node, m map[int]map[int]*Node) (*Node, *Node) {
	up, _ := m[current.y-1][current.x]
	down, _ := m[current.y+1][current.x]
	right, _ := m[current.y][current.x+1]
	left, _ := m[current.y][current.x-1]

	if up != nil && up != previous && (up.t == "S" || up.t == "F" || up.t == "|" || up.t == "7") && (current.t == "S" || current.t == "J" || current.t == "L" || current.t == "|") {
		return up, current
	} else if down != nil && down != previous && (down.t == "S" || down.t == "J" || down.t == "|" || down.t == "L") && (current.t == "S" || current.t == "F" || current.t == "7" || current.t == "|") {
		return down, current
	} else if right != nil && right != previous && (right.t == "S" || right.t == "J" || right.t == "-" || right.t == "7") && (current.t == "S" || current.t == "L" || current.t == "F" || current.t == "-") {
		return right, current
	} else if left != nil && left != previous && (left.t == "S" || left.t == "F" || left.t == "-" || left.t == "L") && (current.t == "S" || current.t == "J" || current.t == "7" || current.t == "-") {
		return left, current
	} else {
		fmt.Println("Current: ", current)
		fmt.Println("Up: ", up)
		fmt.Println("Down: ", down)
		fmt.Println("Right: ", right)
		fmt.Println("Left: ", left)
		panic("No matches!")
	}
}

// |
// -
// L
// J
// 7
// F
// .
func part1(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]*Node)

	var start *Node
	for y, l := range lines {
		m[y] = make(map[int]*Node)
		for x, c := range l {
			m[y][x] = &Node{y, x, string(c)}
			if c == 'S' {
				start = m[y][x]
			}
		}
	}

	ans := 1
	fmt.Println("Start: ", start)
	current, previous := next(start, nil, m)
	for current.t != "S" {
		current, previous = next(current, previous, m)
		ans++
	}

	return ans / 2
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]*Node)

	var start *Node
	for y, l := range lines {
		m[y] = make(map[int]*Node)
		for x, c := range l {
			m[y][x] = &Node{y, x, string(c)}
			if c == 'S' {
				start = m[y][x]
			}
		}
	}

	ans := 1
	fmt.Println("Start: ", start)
	current, previous := next(start, nil, m)
	for current.t != "S" {
		current, previous = next(current, previous, m)
		ans++
	}

	return ans / 2
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
