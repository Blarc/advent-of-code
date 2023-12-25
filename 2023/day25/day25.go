package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
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

func count(node *Node, visited map[string]bool, m map[string]*Node) int {
	sum := 0
	for _, neighbour := range node.edges {
		if !visited[neighbour.to] {
			visited[neighbour.to] = true
			sum += count(m[neighbour.to], visited, m) + 1
		}

	}
	return sum
}

type Node struct {
	name  string
	edges []*Edge
}

type Edge struct {
	from string
	to   string
}

type Subset struct {
	parent string
	rank   int
}

func find(subsets map[string]*Subset, name string) string {
	if subsets[name].parent != name {
		subsets[name].parent = find(subsets, subsets[name].parent)
	}
	return subsets[name].parent
}

func union(subsets map[string]*Subset, a, b string) {
	aRoot := find(subsets, a)
	bRoot := find(subsets, b)

	if subsets[aRoot].rank > subsets[bRoot].rank {
		subsets[bRoot].parent = aRoot
	} else if subsets[aRoot].rank < subsets[bRoot].rank {
		subsets[aRoot].parent = bRoot
	} else {
		subsets[bRoot].parent = aRoot
		subsets[aRoot].rank += 1
	}
}

func karger(v map[string]*Node, e []*Edge) []*Edge {
	subsets := make(map[string]*Subset, len(v))
	for name, _ := range v {
		subsets[name] = &Subset{name, 0}
	}

	vertices := len(v)
	for vertices > 2 {
		randomEdge := e[rand.Intn(len(e))]

		subset1 := find(subsets, randomEdge.from)
		subset2 := find(subsets, randomEdge.to)

		if subset1 == subset2 {
			continue
		} else {
			vertices -= 1
			union(subsets, subset1, subset2)
		}
	}

	var cutEdges []*Edge
	for i := 0; i < len(e); i++ {
		subset1 := find(subsets, e[i].from)
		subset2 := find(subsets, e[i].to)
		if subset1 != subset2 {
			cutEdges = append(cutEdges, e[i])
		}
	}
	return cutEdges
}

func part1(input string) int {
	textLines := strings.Split(input, "\n")

	nodes := make(map[string]*Node)
	var edges []*Edge
	for _, l := range textLines {
		l = strings.Replace(l, ":", "", 1)
		components := strings.Split(l, " ")
		from := components[0]
		for _, to := range components[1:] {
			if nodes[from] == nil {
				nodes[from] = &Node{name: from}
			}
			if nodes[to] == nil {
				nodes[to] = &Node{name: to}
			}

			edgeFrom := &Edge{from, to}
			nodes[from].edges = append(nodes[from].edges, edgeFrom)
			edges = append(edges, edgeFrom)

			edgeTo := &Edge{to, from}
			nodes[to].edges = append(nodes[to].edges, edgeTo)
			edges = append(edges, edgeTo)
		}
	}

	var bestCut []*Edge
	for len(bestCut) != 6 {
		bestCut = karger(nodes, edges)
	}

	for _, cut := range bestCut {
		fromNode := nodes[cut.from]
		fromNode.edges = slices.DeleteFunc(fromNode.edges, func(edge *Edge) bool {
			return edge == cut
		})
	}

	firstGroupCount := count(nodes["cmg"], map[string]bool{"cmg": true}, nodes) + 1
	secondGroupCount := len(nodes) - firstGroupCount
	return firstGroupCount * secondGroupCount
}

func copyGraph(m map[string]*Node) (map[string]*Node, []*Edge) {
	graphCopy := make(map[string]*Node)
	var edgesCopy []*Edge
	for k1, v1 := range m {
		graphCopy[k1] = &Node{name: k1}
		for _, e := range v1.edges {
			edgeCopy := &Edge{e.from, e.to}
			graphCopy[k1].edges = append(graphCopy[k1].edges, edgeCopy)
			edgesCopy = append(edgesCopy, edgeCopy)
		}
	}
	return graphCopy, edgesCopy
}

func part2(input string) int {
	return 0
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
