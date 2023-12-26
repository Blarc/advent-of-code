package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/Blarc/advent-of-code/shared"
	"math"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type State struct {
	pos, dir shared.Point
	distance int
}

func minimizeHeatLoss(start, end shared.Point, minDist, maxDist int, grid [][]int) int {
	queue := []State{
		{pos: start, dir: shared.Point{X: 1}, distance: 0},
		{pos: start, dir: shared.Point{Y: 1}, distance: 0},
	}
	visited := map[State]int{{start, shared.Point{X: 0, Y: 0}, 0}: 0}
	minimumHeatLoss := math.MaxInt

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos.X == end.X && current.pos.Y == end.Y && minDist <= current.distance {
			minimumHeatLoss = shared.Min(minimumHeatLoss, visited[current])
		}

		straightDir := current.dir.Straight()
		straightState := State{
			shared.Point{
				X: current.pos.X + straightDir.X,
				Y: current.pos.Y + straightDir.Y,
			},
			straightDir,
			current.distance + 1,
		}

		if straightState.pos.InsideArrayGrid(grid) && current.distance < maxDist {
			heatLoss := visited[current] + grid[straightState.pos.Y][straightState.pos.X]
			if existingHeatLoss, exists := visited[straightState]; !exists || existingHeatLoss > heatLoss {
				visited[straightState] = heatLoss
				queue = append(queue, straightState)
			}
		}

		leftDir := current.dir.Left()
		leftState := State{
			shared.Point{
				X: current.pos.X + leftDir.X,
				Y: current.pos.Y + leftDir.Y,
			},
			leftDir,
			1,
		}

		if leftState.pos.InsideArrayGrid(grid) && current.distance >= minDist {
			heatLoss := visited[current] + grid[leftState.pos.Y][leftState.pos.X]
			if existingHeatLoss, exists := visited[leftState]; !exists || existingHeatLoss > heatLoss {
				visited[leftState] = heatLoss
				queue = append(queue, leftState)
			}
		}

		rightDir := current.dir.Right()
		rightState := State{
			shared.Point{
				X: current.pos.X + rightDir.X,
				Y: current.pos.Y + rightDir.Y,
			},
			rightDir,
			1,
		}

		if rightState.pos.InsideArrayGrid(grid) && current.distance >= minDist {
			heatLoss := visited[current] + grid[rightState.pos.Y][rightState.pos.X]
			if existingHeatLoss, exists := visited[rightState]; !exists || existingHeatLoss > heatLoss {
				visited[rightState] = heatLoss
				queue = append(queue, rightState)
			}
		}
	}
	return minimumHeatLoss
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]int, len(lines))
	for y, l := range lines {
		grid[y] = make([]int, len(l))
		for x, s := range strings.Split(l, "") {
			grid[y][x] = shared.MustAtoi(s)
		}
	}

	start, end := shared.Point{X: 0, Y: 0}, shared.Point{X: len(grid[0]) - 1, Y: len(grid) - 1}
	fmt.Println("start:", start, "end:", end)
	return minimizeHeatLoss(start, end, 1, 3, grid)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]int, len(lines))
	for y, l := range lines {
		grid[y] = make([]int, len(l))
		for x, s := range strings.Split(l, "") {
			grid[y][x] = shared.MustAtoi(s)
		}
	}

	start, end := shared.Point{X: 0, Y: 0}, shared.Point{X: len(grid[0]) - 1, Y: len(grid) - 1}
	fmt.Println("start:", start, "end:", end)
	return minimizeHeatLoss(start, end, 4, 10, grid)
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
