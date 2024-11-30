package main

import (
	"flag"
	"fmt"
	"math"
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

func part1(input string) int {

	lines := strings.Split(input, "\n")

	m := make([][]string, len(lines))

	var startY, startX int
	for y, l := range strings.Split(input, "\n") {
		for x, c := range strings.Split(l, "") {
			m[y] = append(m[y], c)
			if c == "S" {
				startY = y
				startX = x
			}
		}
	}

	fmt.Println(startY, startX)
	c := make(map[int]map[int]bool)
	c[startY] = make(map[int]bool)
	c[startY][startX] = true

	for i := 0; i < 64; i++ {
		newC := make(map[int]map[int]bool)
		for ky, vy := range c {
			for kx, _ := range vy {

				neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
				for k := 0; k < len(neighbours); k += 2 {
					dy := neighbours[k]
					dx := neighbours[k+1]

					newY := ky + dy
					newX := kx + dx

					if newY < 0 || len(m) <= newY || newX < 0 || len(m[0]) <= newX {
						continue
					}

					if m[newY][newX] != "#" {
						if _, exists := newC[newY]; !exists {
							newC[newY] = make(map[int]bool)
						}
						newC[newY][newX] = true
					}

				}
			}
		}
		c = newC

		//tmp := 0
		//fmt.Println("-------", i, "--------")
		//for y := 0; y < len(m); y++ {
		//	for x := 0; x < len(m[0]); x++ {
		//		if _, exists := c[y][x]; exists {
		//			fmt.Print(coloredString("O", 92))
		//			tmp += 1
		//		} else {
		//			fmt.Print(m[y][x])
		//		}
		//	}
		//	fmt.Println()
		//}
		//fmt.Println(tmp)

	}

	sum := 0
	for _, vy := range c {
		sum += len(vy)
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make([][]string, len(lines))

	var startY, startX int
	for y, l := range strings.Split(input, "\n") {
		for x, c := range strings.Split(l, "") {
			m[y] = append(m[y], c)
			if c == "S" {
				startY = y
				startX = x
			}
		}
	}
	mSize := len(m)
	c := make(map[int]map[int]bool)
	c[startY] = make(map[int]bool)
	c[startY][startX] = true
	numberOfGardensFull := numberOfGardenPlots(startY, startX, 0, 201, m, c)
	fmt.Println(mSize, mSize-startX, mSize*mSize, numberOfGardensFull)

	//steps := 26501365
	steps := 26501365 - startY
	remainderOfSteps := steps % mSize
	steps -= remainderOfSteps
	fmt.Println("remainder", remainderOfSteps)
	fmt.Println(steps / mSize)
	multiplier := ((steps/mSize)*2 + 1) * ((steps/mSize)*2 + 1)

	fmt.Println(multiplier, multiplier*numberOfGardensFull, remainderOfSteps)

	//fmt.Println(numberOfStepsToFull(startY, startX, 0, 0, m))
	//fmt.Println(numberOfStepsToFull(startY, startX, 1, 0, m))

	visited := make(map[int]map[int]bool)
	visited[startY] = make(map[int]bool)
	visited[startY][startX] = true
	//size := 1
	//fmt.Println(numberOfStepsFromTo(startY, startX, -size*len(m)+1, -size*len(m)+1, 0, size, visited, m))
	//fmt.Println(numberOfStepsFromTo(startY, startX, 0, size*len(m)-1, 0, size, visited, m))
	//fmt.Println(numberOfStepsFromTo(startY, startX, size*len(m)-1, 0, 0, size, visited, m))
	//fmt.Println(numberOfStepsFromTo(startY, startX, size*len(m)-1, size*len(m)-1, 0, size, visited, m))

	//steps := 1000
	//gardenPlots0 := numberOfGardenPlots(startY, startX, 0, steps, m)
	//fmt.Println(gardenPlots0)
	//gardenPlots1 := numberOfGardenPlots(startY, startX, 1, steps, m)
	//fmt.Println(gardenPlots1 - gardenPlots0)
	//fmt.Println(gardenPlots1)
	//gardenPlots2 := numberOfGardenPlots(startY, startX, 2, steps, m)
	//fmt.Println(gardenPlots2 - gardenPlots1)
	//fmt.Println(gardenPlots2)

	// 3014472300 too low
	// 1489332600
	// 2978672562
	// 1205167939920000 too high
	// 1205161982596962 too high
	return 0
}

func numberOfStepsFromTo(fromY, fromX, toY, toX, steps, size int, visited map[int]map[int]bool, m [][]string) int {

	fmt.Println(fromY, fromX)
	if fromY == toY && fromX == toX {
		return steps
	}

	best := math.MaxInt
	neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
	for k := 0; k < len(neighbours); k += 2 {
		dy := neighbours[k]
		dx := neighbours[k+1]

		newY := fromY + dy
		newX := fromX + dx

		if newY < -size*len(m) || len(m)+size*len(m) <= newY || newX < -size*len(m) || len(m)+size*len(m) <= newX {
			continue
		}

		newYFixed, newXFixed := newY, newX
		if newYFixed < 0 {
			newYFixed = len(m) + (newYFixed % len(m))
		}
		if len(m) <= newYFixed {
			newYFixed = newYFixed % len(m)
		}
		if newXFixed < 0 {
			newXFixed = len(m[0]) + (newXFixed % len(m[0]))
		}
		if len(m[0]) <= newXFixed {
			newXFixed = newXFixed % len(m[0])
		}

		alreadyVisited, _ := visited[newY][newX]
		if m[newYFixed][newXFixed] != "#" && !alreadyVisited {
			if visited[newY] == nil {
				visited[newY] = make(map[int]bool)
			}

			visited[newY][newX] = true
			res := numberOfStepsFromTo(newY, newX, toY, toX, steps+1, size, visited, m)
			visited[newY][newX] = false

			if res < best {
				best = res
			}
		}
	}
	return best
}

func numberOfGardenPlots(startY, startX, size, steps int, m [][]string, c map[int]map[int]bool) int {
	for i := 0; i < steps; i++ {
		newC := make(map[int]map[int]bool)
		for ky, vy := range c {
			for kx, _ := range vy {

				neighbours := []int{1, 0, -1, 0, 0, 1, 0, -1}
				for k := 0; k < len(neighbours); k += 2 {
					dy := neighbours[k]
					dx := neighbours[k+1]

					newY := ky + dy
					newX := kx + dx

					if newY < -size*len(m) || len(m)+size*len(m) <= newY || newX < -size*len(m) || len(m)+size*len(m) <= newX {
						continue
					}

					newYFixed, newXFixed := newY, newX
					if newYFixed < 0 {
						newYFixed = len(m) + (newYFixed % len(m))
					}
					if len(m) <= newYFixed {
						newYFixed = newYFixed % len(m)
					}
					if newXFixed < 0 {
						newXFixed = len(m[0]) + (newXFixed % len(m[0]))
					}
					if len(m[0]) <= newXFixed {
						newXFixed = newXFixed % len(m[0])
					}

					if m[newYFixed][newXFixed] != "#" {
						if _, exists := newC[newY]; !exists {
							newC[newY] = make(map[int]bool)
						}
						newC[newY][newX] = true
					}

				}
			}
		}
		c = newC
	}

	//for y := 0; y < len(m); y++ {
	//	for x := 0; x < len(m[0]); x++ {
	//		if _, exists := c[y][x]; exists {
	//			fmt.Print(coloredString("O", 92))
	//		} else {
	//			fmt.Print(m[y][x])
	//		}
	//	}
	//	fmt.Println()
	//}

	sum := 0
	for _, vy := range c {
		sum += len(vy)
	}
	return sum
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
