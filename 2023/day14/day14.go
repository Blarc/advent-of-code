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

func part1(input string) int {
	lines := strings.Split(input, "\n")

	empty := make([]int, len(lines[0]))

	sum := 0
	for y, l := range lines {
		for x, c := range strings.Split(l, "") {
			if c == "O" {
				sum += len(lines) - (y - empty[x])
				fmt.Print(y + 1 - empty[x])
			} else if c == "." {
				empty[x] += 1
				fmt.Print(".")
			} else if c == "#" {
				empty[x] = 0
				fmt.Print("#")
			}
		}
		fmt.Println()
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make([][]string, len(lines))

	for y, l := range lines {
		for _, c := range strings.Split(l, "") {
			m[y] = append(m[y], c)
		}
	}

	numOfCycles := 1000000000
	loads := make([]int, numOfCycles)
	memo := make(map[string]int)
	for cycle := 0; cycle < numOfCycles; cycle++ {
		memoString := ""

		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m[y]); x++ {
				memoString += m[y][x]
			}
		}

		moves := []int{-1, 0, 0, -1}
		for i := 0; i < len(moves)-1; i += 2 {
			dy := moves[i]
			dx := moves[i+1]

			emptyY := make([]int, len(lines))
			emptyX := make([]int, len(lines[0]))
			for y := 0; y < len(m); y++ {
				for x := 0; x < len(m[y]); x++ {
					c := m[y][x]
					if c == "O" {
						m[y][x] = "."

						dyt := dy * emptyX[x]
						dxt := dx * emptyY[y]

						m[y+dyt][x+dxt] = "O"
						//fmt.Print(y - empty[x])
					} else if c == "." {
						emptyY[y] += 1
						emptyX[x] += 1
						//fmt.Print(".")
					} else if c == "#" {
						emptyY[y] = 0
						emptyX[x] = 0
						//fmt.Print("#")
					}
				}
				//fmt.Println()
			}

			//for y := 0; y < len(m); y++ {
			//	for x := 0; x < len(m[y]); x++ {
			//		fmt.Print(m[y][x])
			//	}
			//	fmt.Println()
			//}
			//fmt.Println()
		}

		moves = []int{1, 0, 0, 1}
		for i := 0; i < len(moves)-1; i += 2 {
			dy := moves[i]
			dx := moves[i+1]

			emptyY := make([]int, len(lines))
			emptyX := make([]int, len(lines[0]))
			for y := len(m) - 1; y >= 0; y-- {
				for x := len(m[y]) - 1; x >= 0; x-- {
					c := m[y][x]
					if c == "O" {
						m[y][x] = "."

						dyt := dy * emptyX[x]
						dxt := dx * emptyY[y]

						m[y+dyt][x+dxt] = "O"
						//fmt.Print(y - empty[x])
					} else if c == "." {
						emptyY[y] += 1
						emptyX[x] += 1
						//fmt.Print(".")
					} else if c == "#" {
						emptyY[y] = 0
						emptyX[x] = 0
						//fmt.Print("#")
					}
				}
				//fmt.Println()
			}

			//for y := 0; y < len(m); y++ {
			//	for x := 0; x < len(m[y]); x++ {
			//		fmt.Print(m[y][x])
			//	}
			//	fmt.Println()
			//}
			//fmt.Println()
		}

		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m[y]); x++ {
				c := m[y][x]
				if c == "O" {
					loads[cycle] += len(m) - y
				}
			}
		}

		if val, ok := memo[memoString]; ok {
			fmt.Println("Cycle start:", val)
			fmt.Println("Cycle length:", cycle-val)
			tmp := val + (numOfCycles-val)%(cycle-val)
			return loads[tmp-1]
		} else {
			memo[memoString] = cycle
		}

		//for y := 0; y < len(m); y++ {
		//	for x := 0; x < len(m[y]); x++ {
		//		fmt.Print(m[y][x])
		//	}
		//	fmt.Println()
		//}
		//fmt.Println()

	}

	//fmt.Println(loads[(numOfCycles - 8):])

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
