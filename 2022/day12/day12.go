package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func getMoves(pos []int) [][]int {
	moves := make([][]int, 4)
	// left
	moves[0] = []int{pos[0] - 1, pos[1], pos[2] + 1}
	// right
	moves[1] = []int{pos[0] + 1, pos[1], pos[2] + 1}
	// down
	moves[2] = []int{pos[0], pos[1] - 1, pos[2] + 1}
	// up
	moves[3] = []int{pos[0], pos[1] + 1, pos[2] + 1}

	return moves
}

func createKey(pos []int) string {
	return fmt.Sprintf("%d,%d", pos[0], pos[1])
}

func part1(input string) int {

	var grid [][]rune

	start := []int{0, 0}
	end := []int{0, 0}

	for j, l := range strings.Split(input, "\n") {

		row := make([]rune, len(l))
		for i, char := range strings.Split(l, "") {
			r := []rune(char)[0]
			row[i] = r
			if r == []rune("S")[0] {
				start = []int{i, j}
			} else if r == []rune("E")[0] {
				end = []int{i, j}
			}

		}
		grid = append(grid, row)

	}

	grid[start[1]][start[0]] = []rune("a")[0]
	grid[end[1]][end[0]] = []rune("z")[0]
	// fmt.Println("start:", start, []rune("a")[0])
	// fmt.Println("end:", end, []rune("z")[0])

	visited := make(map[string]bool)
	visited[createKey(start)] = true

	queue := [][]int{{start[0], start[1], 0}}

	for {
		// fmt.Println(visited)

		// Pop
		pos := queue[0]

		// Discard top element
		queue = queue[1:]

		if pos[0] == end[0] && pos[1] == end[1] {
			// fmt.Println("Found the end!", pos[2])
			break
		}

		for _, move := range getMoves(pos) {

			if 0 <= move[0] && move[0] < len(grid[0]) && 0 <= move[1] && move[1] < len(grid) {
				visitedMove := visited[createKey(move)]

				if grid[pos[1]][pos[0]]+1 >= grid[move[1]][move[0]] && !visitedMove {
					visited[createKey(move)] = true
					queue = append(queue, move)

				}
			}

		}

	}

	return 1

}

func part2(input string) int {
	var grid [][]rune

	var starts [][]int
	end := []int{0, 0}

	for j, l := range strings.Split(input, "\n") {

		row := make([]rune, len(l))
		for i, char := range strings.Split(l, "") {
			r := []rune(char)[0]
			row[i] = r
			if r == []rune("S")[0] {
				starts = append(starts, []int{i, j})
			} else if r == []rune("E")[0] {
				end = []int{i, j}
			} else if r == []rune("a")[0] {
				starts = append(starts, []int{i, j})
			}

		}
		grid = append(grid, row)

	}

	var results []int
	for _, start := range starts {
		grid[start[1]][start[0]] = []rune("a")[0]
		grid[end[1]][end[0]] = []rune("z")[0]
		// fmt.Println("start:", start, []rune("a")[0])
		// fmt.Println("end:", end, []rune("z")[0])

		visited := make(map[string]bool)
		visited[createKey(start)] = true

		queue := [][]int{{start[0], start[1], 0}}

		for {
			// Pop
			if len(queue) == 0 {
				break
			}
			pos := queue[0]

			// Discard top element
			queue = queue[1:]

			if pos[0] == end[0] && pos[1] == end[1] {
				// fmt.Println("Found the end!", pos[2])
				results = append(results, pos[2])
				break
			}

			for _, move := range getMoves(pos) {

				if 0 <= move[0] && move[0] < len(grid[0]) && 0 <= move[1] && move[1] < len(grid) {
					visitedMove := visited[createKey(move)]

					if grid[pos[1]][pos[0]]+1 >= grid[move[1]][move[0]] && !visitedMove {
						visited[createKey(move)] = true
						queue = append(queue, move)

					}
				}

			}

		}

	}
	sort.Ints(results)
	return results[0]
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

	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
}
