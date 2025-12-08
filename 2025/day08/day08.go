package main

import (
	"container/heap"
	"flag"
	"fmt"
	"math"
	"sort"
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

type Point struct {
	X, Y, Z   int
	circuitId int
}

type Edge struct {
	Start, End *Point
	Distance   int
}

type PriorityQueue []*Edge

func (pq *PriorityQueue) Len() int { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Distance < (*pq)[j].Distance
}
func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (p Point) Euclidian(other *Point) int {
	dx := p.X - other.X
	dy := p.Y - other.Y
	dz := p.Z - other.Z
	return int(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

func (p Point) ClosestPoint(points []Point) *Point {
	minDistance := math.MaxInt32
	closest := &Point{}
	for _, point := range points {
		dist := p.Euclidian(&point)
		if dist < minDistance {
			minDistance = dist
			closest = &point
		}
	}
	return closest
}

func part1(input string) int {
	result := 0
	lines := strings.Split(input, "\n")

	boxes := make([]*Point, len(lines))
	for i, l := range lines {
		s := strings.Split(l, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		z, _ := strconv.Atoi(s[2])
		boxes[i] = &Point{x, y, z, i}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			if i != j {
				edge := &Edge{boxes[i], boxes[j], boxes[i].Euclidian(boxes[j])}
				heap.Push(&pq, edge)
			}
		}
	}

	count := 0
	for pq.Len() > 0 && count < 1000 {
		edge := heap.Pop(&pq).(*Edge)
		count++

		if edge.Start.circuitId == edge.End.circuitId {
			continue
		}

		if edge.Start.circuitId < edge.End.circuitId {
			tmp := edge.End.circuitId
			for _, box := range boxes {
				if box.circuitId == tmp {
					box.circuitId = edge.Start.circuitId
				}
			}
		} else {
			tmp := edge.Start.circuitId
			for _, box := range boxes {
				if box.circuitId == tmp {
					box.circuitId = edge.End.circuitId
				}
			}
		}
	}

	circuitSizes := make([]int, len(boxes))
	for _, box := range boxes {
		circuitSizes[box.circuitId]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))
	fmt.Printf("circuit sizes: %v\n", circuitSizes[:3])

	result = 1
	for _, size := range circuitSizes[:3] {
		result *= size
	}

	return result
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	boxes := make([]*Point, len(lines))
	for i, l := range lines {
		s := strings.Split(l, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		z, _ := strconv.Atoi(s[2])
		boxes[i] = &Point{x, y, z, i}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			if i != j {
				edge := &Edge{boxes[i], boxes[j], boxes[i].Euclidian(boxes[j])}
				heap.Push(&pq, edge)
			}
		}
	}

	count := 0
	lastEdge := pq[0]
	for pq.Len() > 0 && count < len(boxes)-1 {
		edge := heap.Pop(&pq).(*Edge)

		if edge.Start.circuitId == edge.End.circuitId {
			continue
		}

		if edge.Start.circuitId < edge.End.circuitId {
			tmp := edge.End.circuitId
			for _, box := range boxes {
				if box.circuitId == tmp {
					box.circuitId = edge.Start.circuitId
				}
			}
		} else {
			tmp := edge.Start.circuitId
			for _, box := range boxes {
				if box.circuitId == tmp {
					box.circuitId = edge.End.circuitId
				}
			}
		}

		lastEdge = edge
		count++
	}

	fmt.Printf("last edge: %v - %v\n", lastEdge.Start, lastEdge.End)
	return lastEdge.Start.X * lastEdge.End.X
}

func main() {

	inputPtr := flag.Bool("input", false, "sample or input")

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")

	flag.Parse()

	var inputText string
	if *inputPtr {
		inputText = strings.TrimSpace(input)
		fmt.Println("Running part", part, "on input.txt.")
	} else {
		inputText = strings.TrimSpace(sample)
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
