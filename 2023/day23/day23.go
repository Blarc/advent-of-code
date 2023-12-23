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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func coloredString(text string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

func mustAtoi(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func dfs(cy, cx, endY, endX, length int, visited [][]bool, m [][]string) int {

	if cy == endY && cx == endX {
		return length
	}

	best := 0
	neighbours := []int{1, 0, 0, 1, -1, 0, 0, -1}
	if m[cy][cx] == ">" {
		neighbours = []int{0, 1}
	} else if m[cy][cx] == "<" {
		neighbours = []int{0, -1}
	} else if m[cy][cx] == "v" {
		neighbours = []int{1, 0}
	} else if m[cy][cx] == "^" {
		neighbours = []int{-1, 0}
	}

	for i := 0; i < len(neighbours); i += 2 {
		dy := neighbours[i]
		dx := neighbours[i+1]

		nextY := cy + dy
		nextX := cx + dx

		if nextY < 0 || len(m) <= nextY || nextX < 0 || len(m[0]) <= nextX {
			continue
		}

		if m[nextY][nextX] != "#" && !visited[nextY][nextX] {
			visited[nextY][nextX] = true
			res := dfs(nextY, nextX, endY, endX, length+1, visited, m)
			if best < res {
				best = res
			}
			visited[nextY][nextX] = false

		}
	}

	return best
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	m := make([][]string, len(lines))
	visited := make([][]bool, len(lines))

	for y, l := range lines {
		if m[y] == nil {
			m[y] = make([]string, len(l))
			visited[y] = make([]bool, len(l))
		}
		for x, c := range strings.Split(l, "") {
			m[y][x] = c
			visited[y][x] = false
		}
	}

	startY := 0
	startX := 0
	endY := len(m) - 1
	endX := 0
	for x := 0; x < len(m[0]); x++ {
		if m[startY][x] == "." {
			startX = x
		}
		if m[endY][x] == "." {
			endX = x
		}
	}
	fmt.Println("start:", startY, startX)
	fmt.Println("end:", endY, endX)

	visited[startY][startX] = true
	return dfs(startY, startX, endY, endX, 0, visited, m)
}

type Node struct {
	y, x  int
	edges []*Edge
}

type Edge struct {
	to     *Node
	length int
}

func createGraph(prev, current, end *Node, length int, visited map[*Node]bool, m map[int]map[int]*Node) {
	// So we see end as one of the edges when searching for longest path
	if current.y == end.y && current.x == end.x {
		current.edges = append(current.edges, &Edge{prev, length})
		prev.edges = append(prev.edges, &Edge{current, length})
		return
	}

	var neighbours []*Node
	dydx := []int{1, 0, 0, 1, -1, 0, 0, -1}
	for i := 0; i < len(dydx); i += 2 {
		dy := dydx[i]
		dx := dydx[i+1]

		nextY := current.y + dy
		nextX := current.x + dx

		neighbour, neighbourExists := m[nextY][nextX]
		neighbourVisited, _ := visited[neighbour]
		if neighbourExists && !neighbourVisited {
			neighbours = append(neighbours, neighbour)
		} else if prev != neighbour && neighbourExists && neighbourVisited && len(neighbour.edges) > 1 {
			// This handles cases when neighbour is a crossroad, but was already visited
			// We still need to create an edge between previous crossroad and this one
			prev.edges = append(prev.edges, &Edge{neighbour, length + 2})
			neighbour.edges = append(neighbour.edges, &Edge{prev, length + 2})
		}
	}

	if len(neighbours) == 1 {
		neighbour := neighbours[0]
		visited[neighbour] = true
		createGraph(prev, neighbour, end, length+1, visited, m)
	} else if len(neighbours) > 1 {
		prev.edges = append(prev.edges, &Edge{current, length + 1})
		current.edges = append(current.edges, &Edge{prev, length + 1})
		for _, neighbour := range neighbours {
			visited[neighbour] = true
			createGraph(current, neighbour, end, 0, visited, m)
		}
	}
}

func findLongestPath(current, end *Node, length int, visited map[*Node]bool, m map[int]map[int]*Node, path []*Node) (int, []*Node) {
	if current.y == end.y && current.x == end.x {
		return length, path
	}

	longestLength := 0
	var longestPath []*Node
	for _, edge := range current.edges {
		neighbourVisited, _ := visited[edge.to]
		if !neighbourVisited {
			visited[edge.to] = true
			path = append(path, edge.to)
			currentLength, currentPath := findLongestPath(edge.to, end, length+edge.length, visited, m, path)
			if currentLength > longestLength {
				longestLength = currentLength
				for _, p := range currentPath {
					longestPath = append(longestPath, p)
				}
			}
			path = path[:len(path)-1]
			visited[edge.to] = false
		}
	}

	return longestLength, longestPath
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int]map[int]*Node)

	var start, end *Node
	for y, l := range lines {
		if m[y] == nil {
			m[y] = make(map[int]*Node)
		}
		for x, c := range strings.Split(l, "") {
			if c != "#" {
				m[y][x] = &Node{y: y, x: x}
				if y == 0 {
					start = m[y][x]
				}
				if y == len(lines)-1 {
					end = m[y][x]
				}
			}
		}

	}

	fmt.Println("start:", start)
	fmt.Println("end:", end)

	visited := map[*Node]bool{start: true}
	createGraph(start, start, end, 0, visited, m)

	//// Just print the nodes and their edges
	//for y := 0; y < len(lines); y++ {
	//	for x := 0; x < len(lines[0]); x++ {
	//		if node, nodeExists := m[y][x]; nodeExists {
	//			if len(node.edges) > 0 {
	//				fmt.Printf("%+v\n", node)
	//				for _, edge := range node.edges {
	//					fmt.Printf("  %+v: %d\n", edge.to, edge.length)
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//// Print the grid and mark crossroads with S
	//for y := 0; y < len(lines); y++ {
	//	fmt.Printf("%2d: ", y)
	//	for x := 0; x < len(lines[0]); x++ {
	//		if node, nodeExists := m[y][x]; nodeExists {
	//			if len(node.edges) > 0 {
	//				fmt.Print(coloredString("S", 92))
	//			} else {
	//				fmt.Print(".")
	//			}
	//		} else {
	//			fmt.Print("#")
	//		}
	//	}
	//	fmt.Println()
	//}

	visited = map[*Node]bool{start: true}
	length, _ := findLongestPath(start, end, 0, visited, m, []*Node{start})
	return length
}

func main() {

	// 6668 not correct
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
