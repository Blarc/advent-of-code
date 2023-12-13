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
	y    int
	x    int
	t    string
	loop bool
}

func findStart(start *Node, m map[int]map[int]*Node) string {
	up, _ := m[start.y-1][start.x]
	down, _ := m[start.y+1][start.x]
	right, _ := m[start.y][start.x+1]
	left, _ := m[start.y][start.x-1]

	upValid := up != nil && (up.t == "F" || up.t == "|" || up.t == "7")
	downValid := down != nil && (down.t == "J" || down.t == "|" || down.t == "L")
	rightValid := right != nil && (right.t == "J" || right.t == "-" || right.t == "7")
	leftValid := left != nil && (left.t == "F" || left.t == "-" || left.t == "L")

	if upValid && downValid {
		return "|"
	} else if upValid && rightValid {
		return "L"
	} else if upValid && downValid {
		return "|"
	} else if upValid && leftValid {
		return "J"
	} else if downValid && rightValid {
		return "F"
	} else if downValid && leftValid {
		return "7"
	} else if rightValid && leftValid {
		return "-"
	} else {
		panic("No matches?")
	}
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

func coloredString(text string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
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
			m[y][x] = &Node{y, x, string(c), false}
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
			m[y][x] = &Node{y, x, string(c), false}
			if c == 'S' {
				start = m[y][x]
			}
		}
	}

	start.loop = true
	current, previous := next(start, nil, m)
	for current.t != "S" {
		current.loop = true
		current, previous = next(current, previous, m)
	}

	start.t = findStart(start, m)

	ans := 0
	for y := 0; y < len(lines); y++ {
		inside := false
		for x := 0; x < len(lines[y]); x++ {
			node := m[y][x]
			// Also works with (node.t == "|" || node.t == "J" || node.t == "L"); don't know why though...
			if node.loop && (node.t == "|" || node.t == "F" || node.t == "7") {
				inside = !inside
			} else if !node.loop && inside {
				ans += 1
			}

			//if node.loop {
			//	fmt.Print(coloredString(node.t, 36))
			//} else if !node.loop && inside {
			//	fmt.Print(coloredString("I", 31))
			//} else {
			//	fmt.Print(coloredString("O", 34))
			//}
		}
		//fmt.Println()
	}

	return ans
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

	part2(strings.TrimSpace(inputText))

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(strings.TrimSpace(inputText)))
	} else {
		fmt.Println("Result:", part2(strings.TrimSpace(inputText)))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
