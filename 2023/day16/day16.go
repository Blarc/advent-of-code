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

type Direction struct {
	dy, dx int
}

type Node struct {
	y, x int
	t    string
	dir  []*Direction
}

func containsDir(newDir *Direction, directions []*Direction) bool {
	for i := 0; i < len(directions); i++ {
		if newDir.dy == directions[i].dy && newDir.dx == directions[i].dx {
			return true
		}
	}
	return false
}

func findAll(current *Node, previousDir *Direction, m map[int]map[int]*Node) int {

	if current == nil {
		return 0
	}

	if containsDir(previousDir, current.dir) {
		return 0
	}

	ans := 0
	if len(current.dir) == 0 {
		ans += 1
	}

	current.dir = append(current.dir, previousDir)
	var directions []*Direction
	if current.t == "\\" {
		// right > down
		if previousDir.dy == 0 && previousDir.dx == 1 {
			directions = append(directions, &Direction{1, 0})
		}
		// left > up
		if previousDir.dy == 0 && previousDir.dx == -1 {
			directions = append(directions, &Direction{-1, 0})
		}
		// up > left
		if previousDir.dy == 1 && previousDir.dx == 0 {
			directions = append(directions, &Direction{0, 1})
		}
		// down > right
		if previousDir.dy == -1 && previousDir.dx == 0 {
			directions = append(directions, &Direction{0, -1})
		}

	} else if current.t == "/" {
		// right > up
		if previousDir.dy == 0 && previousDir.dx == 1 {
			directions = append(directions, &Direction{-1, 0})
		}
		// left > down
		if previousDir.dy == 0 && previousDir.dx == -1 {
			directions = append(directions, &Direction{1, 0})
		}
		// up > right
		if previousDir.dy == -1 && previousDir.dx == 0 {
			directions = append(directions, &Direction{0, 1})
		}
		// down > left
		if previousDir.dy == 1 && previousDir.dx == 0 {
			directions = append(directions, &Direction{0, -1})
		}

	} else if current.t == "-" && previousDir.dy != 0 {
		directions = append(directions, &Direction{0, -1})
		directions = append(directions, &Direction{0, 1})
	} else if current.t == "|" && previousDir.dx != 0 {
		directions = append(directions, &Direction{-1, 0})
		directions = append(directions, &Direction{1, 0})
	} else {
		directions = append(directions, previousDir)
	}

	sum := ans
	for _, dir := range directions {
		newY := current.y + dir.dy
		newX := current.x + dir.dx

		sum += findAll(m[newY][newX], dir, m)

	}
	return sum

}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]*Node)
	for y, l := range lines {
		for x, s := range strings.Split(l, "") {
			if m[y] == nil {
				m[y] = make(map[int]*Node)
			}
			m[y][x] = &Node{
				y: y,
				x: x,
				t: s,
			}
		}
	}

	ans := findAll(m[0][0], &Direction{0, 1}, m)

	//for y := 0; y < len(m); y++ {
	//	for x := 0; x < len(m[y]); x++ {
	//		node := m[y][x]
	//		if len(node.dir) != 0 {
	//			if len(node.dir) > 0 && node.t != "." {
	//				if len(node.dir) > 1 {
	//					fmt.Print(len(node.dir))
	//
	//				} else {
	//					fmt.Print(node.t)
	//				}
	//			} else if node.dir[0].dy == 0 && node.dir[0].dx == 1 {
	//				fmt.Print(">")
	//			} else if node.dir[0].dy == 0 && node.dir[0].dx == -1 {
	//				fmt.Print("<")
	//			} else if node.dir[0].dy == 1 && node.dir[0].dx == 0 {
	//				fmt.Print("v")
	//			} else if node.dir[0].dy == -1 && node.dir[0].dx == 0 {
	//				fmt.Print("^")
	//			} else {
	//				fmt.Print("#")
	//			}
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Println()
	//}

	return ans
}

func part2(input string) int {

	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]*Node)
	for y, l := range lines {
		for x, s := range strings.Split(l, "") {
			if m[y] == nil {
				m[y] = make(map[int]*Node)
			}
			m[y][x] = &Node{
				y: y,
				x: x,
				t: s,
			}
		}
	}

	best := -1
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			ans := -1
			if y == 0 {
				ans = findAll(m[y][x], &Direction{1, 0}, m)
				if ans > best {
					best = ans
				}
			}
			if x == 0 {
				ans = findAll(m[y][x], &Direction{0, 1}, m)
				if ans > best {
					best = ans
				}
			}
			if y == len(m)-1 {
				ans = findAll(m[y][x], &Direction{-1, 0}, m)
				if ans > best {
					best = ans
				}
			}
			if x == len(m[y])-1 {
				ans = findAll(m[y][x], &Direction{0, -1}, m)
				if ans > best {
					best = ans
				}
			}

			for yn := 0; yn < len(m); yn++ {
				for xn := 0; xn < len(m[yn]); xn++ {
					m[yn][xn].dir = nil
				}
			}
		}
	}

	// findAll(m[0][3], &Direction{1, 0}, m)

	return best
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
