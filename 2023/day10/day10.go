package main

import (
	"flag"
	"fmt"
	"slices"
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
	y        int
	x        int
	t        string
	l        int
	previous *Node
	dir      int
	visited  bool
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]Node)

	var startY, startX int
	for y, l := range lines {
		m[y] = make(map[int]Node)
		for x, c := range l {
			m[y][x] = Node{y, x, string(c), 0, nil, 0, false}
			if c == 'S' {
				startY = y
				startX = x
			}
		}
	}

	s := make([]Node, 0)
	s = append(s, Node{startY, startX, "S", 0, nil, 0, false})

	// find the longest loop
	var lengths []int
	for len(s) > 0 {
		sl := len(s)
		current := s[sl-1]
		s = s[:sl-1]

		if current.t == "S" && current.l > 0 {
			println("found goal at", current.y, current.x, current.l/2)
			lengths = append(lengths, current.l/2)
			continue
		}

		neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
		for i := 0; i < len(neighbours); i += 2 {
			dy := neighbours[i]
			dx := neighbours[i+1]
			ny := current.y + dy
			nx := current.x + dx
			if ny < 0 || ny > len(lines) || nx < 0 || nx >= len(lines[0]) {
				continue
			}
			n := m[ny][nx]
			if n.t == "." {
				continue
			}
			if current.previous != nil && ny == current.previous.y && nx == current.previous.x {
				continue
			}

			n.l = current.l + 1
			n.previous = &current
			if n.t == "S" {
				s = append(s, n)
			} else if current.t == "|" && dx == 0 {
				if dy == 1 && (n.t == "J" || n.t == "L") {
					s = append(s, n)
				} else if n.t == "7" || n.t == "F" {
					s = append(s, n)
				} else if n.t == "|" {
					s = append(s, n)
				}
			} else if current.t == "-" && dy == 0 {
				if dx == 1 && (n.t == "J" || n.t == "7") {
					s = append(s, n)
				} else if n.t == "L" || n.t == "F" {
					s = append(s, n)
				} else if n.t == "-" {
					s = append(s, n)
				}
			} else if current.t == "L" && (dx == 1 || dy == -1) {
				if dx == 1 && (n.t == "J" || n.t == "7" || n.t == "-") {
					s = append(s, n)
				} else if dy == -1 && (n.t == "|" || n.t == "7" || n.t == "F") {
					s = append(s, n)
				}
			} else if current.t == "J" && (dx == -1 || dy == -1) {
				if dx == -1 && (n.t == "L" || n.t == "F" || n.t == "-") {
					s = append(s, n)
				} else if dy == -1 && (n.t == "|" || n.t == "F" || n.t == "7") {
					s = append(s, n)
				}
			} else if current.t == "7" && (dx == -1 || dy == 1) {
				if dx == -1 && (n.t == "-" || n.t == "L" || n.t == "F") {
					s = append(s, n)
				} else if dy == 1 && (n.t == "|" || n.t == "L" || n.t == "J") {
					s = append(s, n)
				}
			} else if current.t == "F" && (dx == 1 || dy == 1) {
				if dx == 1 && (n.t == "J" || n.t == "7" || n.t == "-") {
					s = append(s, n)
				} else if dy == 1 && (n.t == "|" || n.t == "J" || n.t == "L") {
					s = append(s, n)
				}
			} else if current.t == "S" && current.l == 0 {
				s = append(s, n)
			}

		}
	}
	return slices.Max(lengths)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]Node)

	var startY, startX int
	for y, l := range lines {
		m[y] = make(map[int]Node)
		for x, c := range l {
			m[y][x] = Node{y, x, string(c), 0, nil, 0, false}
			if c == 'S' {
				startY = y
				startX = x
			}
		}
	}

	s := make([]Node, 0)
	s = append(s, Node{startY, startX, "S", 0, nil, 0, false})

	// find the longest loop
	var maxLength = 0
	var maxPath [][]int

	yMap := make(map[int][]Node)

	pathMap := make(map[int]map[int]Node)

	for len(s) > 0 {
		sl := len(s)
		current := s[sl-1]
		s = s[:sl-1]

		if current.t == "S" && current.l > 0 {
			//println("found goal at", current.y, current.x, current.l/2)

			if current.l/2 > maxLength {
				maxLength = current.l / 2
				maxPath = make([][]int, 0)
				yMap = make(map[int][]Node)
				pathMap = make(map[int]map[int]Node)
				for current.previous != nil {
					maxPath = append(maxPath, []int{current.y, current.x, current.dir})
					current = *current.previous

					if yMap[current.y] == nil {
						yMap[current.y] = []Node{{current.y, current.x, current.t, 0, nil, current.dir, false}}
					} else {
						yMap[current.y] = append(yMap[current.y], Node{current.y, current.x, current.t, 0, nil, current.dir, false})
					}

					if pathMap[current.y] == nil {
						pathMap[current.y] = make(map[int]Node)
					}
					pathMap[current.y][current.x] = Node{current.y, current.x, current.t, 0, nil, current.dir, false}

				}
			}

			continue
		}

		neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
		for i := 0; i < len(neighbours); i += 2 {
			dy := neighbours[i]
			dx := neighbours[i+1]
			ny := current.y + dy
			nx := current.x + dx
			if ny < 0 || ny > len(lines) || nx < 0 || nx >= len(lines[0]) {
				continue
			}
			n := m[ny][nx]
			if n.t == "." {
				continue
			}
			if current.previous != nil && ny == current.previous.y && nx == current.previous.x {
				continue
			}

			if shouldAppend(n, current, dx, dy) {
				n.l = current.l + 1
				n.previous = &current

				//if n.t == "7" || n.t == "F" {
				//	n.dir = 1
				//} else if n.t == "J" || n.t == "L" {
				//	n.dir = -1
				//}

				if n.t == "F" && dy == 0 && dx == -1 {
					// left + F = down
					n.dir = 0
				} else if n.t == "J" && dy == 0 && dx == 1 {
					// right + J = up
					n.dir = 2
				} else if n.t == "L" && dy == 0 && dx == -1 {
					// left + L = up
					n.dir = 2
				} else if n.t == "7" && dy == 0 && dx == 1 {
					// right + 7 = down
					n.dir = 0
				} else if dy == 1 && dx == 0 {
					// down
					n.dir = 0
				} else if dy == 0 && dx == 1 {
					// right
					n.dir = -1
				} else if dy == -1 && dx == 0 {
					// up
					n.dir = 2
				} else if dy == 0 && dx == -1 {
					// left
					n.dir = 1
				}

				s = append(s, n)
			}

		}
	}

	//for i := len(maxPath) - 1; i >= 0; i-- {
	//	fmt.Printf("%s,%v ", m[maxPath[i][0]][maxPath[i][1]].t, maxPath[i])
	//}
	//fmt.Println()

	ylen := len(m)
	xlen := len(m[0])
	for y := -1; y <= ylen; y++ {
		if y == -1 || y == ylen {
			m[y] = make(map[int]Node)
			for x := -1; x <= xlen; x++ {
				m[y][x] = Node{y, x, ".", 0, nil, 0, false}
			}
		} else {
			m[y][-1] = Node{y, -1, ".", 0, nil, 0, false}
			m[y][xlen] = Node{y, xlen, ".", 0, nil, 0, false}
		}
	}

	//for y := -1; y < len(m); y++ {
	//	for x := -1; x < len(m[y]); x++ {
	//		fmt.Print(m[y][x].t)
	//	}
	//	fmt.Println()
	//}

	m[-1][-1] = Node{-1, -1, "O", 0, nil, 0, true}
	stack := []Node{m[-1][-1]}
	count := 0
	for len(stack) > 0 {
		sl := len(stack)
		current := stack[sl-1]
		stack = stack[:sl-1]

		if current.t == "O" {
			count += 1
		}

		neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
		for i := 0; i < len(neighbours); i += 2 {
			dy := neighbours[i]
			dx := neighbours[i+1]
			ny := current.y + dy
			nx := current.x + dx

			if (current.t == "J" || current.t == "L" || current.t == "7" || current.t == "F" || current.t == "|") && dx != 0 && m[ny][nx+1].t == "." {
				continue
			}

			n := m[ny][nx]
			_, contains := pathMap[ny][nx]
			_, contains2 := m[ny][nx]

			if !contains && !n.visited && contains2 {
				m[ny][nx] = Node{ny, nx, "O", 0, nil, 0, true}
				stack = append(stack, m[ny][nx])
			} else if n.y != 0 && n.y != len(m)-3 && n.x != 0 && n.x != len(m[0])-1 {
				if n.t == "J" && (m[n.y][n.x+1].t == "F" || m[n.y][n.x+1].t == "L" || m[n.y][n.x+1].t == "|" || m[n.y][n.x+1].t == ".") && !n.visited && dx == 0 {
					println(n.y)
					m[ny][nx] = Node{ny, nx, n.t, 0, nil, 0, true}
					stack = append(stack, m[ny][nx])
				} else if n.t == "|" && (m[n.y][n.x+1].t == "F" || m[n.y][n.x+1].t == "L" || m[n.y][n.x+1].t == "|" || m[n.y][n.x-1].t == "J" || m[n.y][n.x-1].t == "7" || m[n.y][n.x-1].t == "|") && !n.visited && dx == 0 {
					m[ny][nx] = Node{ny, nx, n.t, 0, nil, 0, true}
					stack = append(stack, m[ny][nx])
				} else if n.t == "L" && (m[n.y][n.x-1].t == "J" || m[n.y][n.x-1].t == "|") && !n.visited && dx == 0 {
					m[ny][nx] = Node{ny, nx, n.t, 0, nil, 0, true}
					stack = append(stack, m[ny][nx])
				} else if n.t == "7" && (m[n.y][n.x+1].t == "F" || m[n.y][n.x+1].t == "L" || m[n.y][n.x+1].t == "|") && !n.visited && dx == 0 {
					m[ny][nx] = Node{ny, nx, n.t, 0, nil, 0, true}
					stack = append(stack, m[ny][nx])
				} else if n.t == "F" && (m[n.y][n.x-1].t == "J" || m[n.y][n.x-1].t == "7" || m[n.y][n.x-1].t == "|") && !n.visited && dx == 0 && dy == -1 {
					m[ny][nx] = Node{ny, nx, n.t, 0, nil, 0, true}
					stack = append(stack, m[ny][nx])
				}
			}
		}
	}

	println(len(m) - 1)
	println("count", count)

	for y := 0; y < len(m)-2; y++ {
		for x := 0; x < len(m[y])-2; x++ {
			_, ok := pathMap[y][x]
			if m[y][x].visited {
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 94, m[y][x].t)
				fmt.Print(colored)
			} else if ok {
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 95, m[y][x].t)
				fmt.Print(colored)
			} else {
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, "I")
				fmt.Print(colored)
			}
		}
		fmt.Println()
	}
	//for y := 0; y <= len(yMap); y++ {
	//	g := yMap[y]
	//	slices.SortFunc(g, func(a, b Node) int {
	//		return a.x - b.x
	//	})
	//	fmt.Printf("%d: %v %d\n", y+1, g, len(g))
	//	//i := 0
	//	//for i < len(g) {
	//	//	j := i + 1
	//	//	for j < len(g) {
	//	//		if math.Abs(float64(g[j].dir-g[i].dir)) == 2 {
	//	//			fmt.Printf("adding %d %v %v %d\n", y+1, g[i], g[j], g[j].x-g[i].x-1)
	//	//			sum += g[j].x - g[i].x - 1
	//	//			break
	//	//		}
	//	//		j += 1
	//	//	}
	//	//	i = j + 1
	//	//}
	//
	//	//i := 0
	//	//for i < len(g)-1 {
	//	//	i2 := g[i+1].x - g[i].x - 1
	//	//
	//	//	if math.Abs(float64(g[i+1].dir-g[i].dir)) == 2 {
	//	//		fmt.Printf("adding %d %v %v %d\n", y+1, g[i], g[i+1], g[i+1].x-g[i].x-1)
	//	//		sum += i2
	//	//		i += 2
	//	//	} else {
	//	//		i += 1
	//	//	}
	//	//}
	//
	//	//for i := 0; i < len(g)-1; i += 2 {
	//	//	i2 := g[i+1] - g[i] - 1
	//	//	println("adding", y+1, g[i], g[i+1], i2)
	//	//}
	//}

	// 657 too high
	return len(m)*len(m[0]) - count - len(maxPath)
}

func shouldAppend(n Node, current Node, dx int, dy int) bool {
	if n.t == "S" {
		return true
	} else if current.t == "|" && dx == 0 {
		if dy == 1 && (n.t == "J" || n.t == "L") {
			return true
		} else if n.t == "7" || n.t == "F" {
			return true
		} else if n.t == "|" {
			return true
		}
	} else if current.t == "-" && dy == 0 {
		if dx == 1 && (n.t == "J" || n.t == "7") {
			return true
		} else if n.t == "L" || n.t == "F" {
			return true
		} else if n.t == "-" {
			return true
		}
	} else if current.t == "L" && (dx == 1 || dy == -1) {
		if dx == 1 && (n.t == "J" || n.t == "7" || n.t == "-") {
			return true
		} else if dy == -1 && (n.t == "|" || n.t == "7" || n.t == "F") {
			return true
		}
	} else if current.t == "J" && (dx == -1 || dy == -1) {
		if dx == -1 && (n.t == "L" || n.t == "F" || n.t == "-") {
			return true
		} else if dy == -1 && (n.t == "|" || n.t == "F" || n.t == "7") {
			return true
		}
	} else if current.t == "7" && (dx == -1 || dy == 1) {
		if dx == -1 && (n.t == "-" || n.t == "L" || n.t == "F") {
			return true
		} else if dy == 1 && (n.t == "|" || n.t == "L" || n.t == "J") {
			return true
		}
	} else if current.t == "F" && (dx == 1 || dy == 1) {
		if dx == 1 && (n.t == "J" || n.t == "7" || n.t == "-") {
			return true
		} else if dy == 1 && (n.t == "|" || n.t == "J" || n.t == "L") {
			return true
		}
	} else if current.t == "S" && current.l == 0 {
		return true
	}
	return false
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
