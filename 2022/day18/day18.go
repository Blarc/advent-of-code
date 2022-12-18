package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func createKey(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func part1(input string) int {

	cubes := make(map[string]int)

	for _, l := range strings.Split(input, "\n") {
		re := regexp.MustCompile("[0-9]+")
		coordsString := re.FindAllString(l, -1)

		x, _ := strconv.Atoi(coordsString[0])
		y, _ := strconv.Atoi(coordsString[1])
		z, _ := strconv.Atoi(coordsString[2])

		cubes[createKey(x, y, z)] = 0

		neighbours := [][]int{
			{-1, 0, 0},
			{1, 0, 0},
			{0, -1, 0},
			{0, 1, 0},
			{0, 0, -1},
			{0, 0, 1},
		}

		for i := 0; i < len(neighbours); i++ {
			nx := x + neighbours[i][0]
			ny := y + neighbours[i][1]
			nz := z + neighbours[i][2]

			_, exists := cubes[createKey(nx, ny, nz)]
			if exists {
				cubes[createKey(x, y, z)] += 1
				cubes[createKey(nx, ny, nz)] += 1
			}
		}
	}

	covered := 0
	for _, cube := range cubes {
		covered += cube
	}
	// 64
	return len(cubes)*6 - covered
}

func part2(input string) int {

	neighbours := [][]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}

	re := regexp.MustCompile("[0-9]+")
	lavaCubes := make(map[string]bool)

	minX := 1000
	minY := 1000
	minZ := 1000

	maxX := -1
	maxY := -1
	maxZ := -1

	for _, l := range strings.Split(input, "\n") {
		coordsString := re.FindAllString(l, -1)

		x, _ := strconv.Atoi(coordsString[0])
		y, _ := strconv.Atoi(coordsString[1])
		z, _ := strconv.Atoi(coordsString[2])

		lavaCubes[createKey(x, y, z)] = true

		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}

		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}

		if z < minZ {
			minZ = z
		} else if z > maxZ {
			maxZ = z
		}
	}

	minX -= 1
	minY -= 1
	minZ -= 1

	maxX += 1
	maxY += 1
	maxZ += 1

	start := []int{minX, minY, minZ}
	// end := []int{maxX, maxY, maxZ}

	result := 0
	visited := make(map[string]bool)

	queue := [][]int{start}
	visited[createKey(start[0], start[1], start[2])] = true

	for {

		if len(queue) == 0 {
			break
		}

		// Pop
		v := queue[0]

		// Discard top element
		queue = queue[1:]

		for i := 0; i < len(neighbours); i++ {
			nx := v[0] + neighbours[i][0]
			ny := v[1] + neighbours[i][1]
			nz := v[2] + neighbours[i][2]

			lava := lavaCubes[createKey(nx, ny, nz)]
			if lava {
				result += 1
			} else if !visited[createKey(nx, ny, nz)] &&
				nx >= minX && nx <= maxX &&
				ny >= minY && ny <= maxY &&
				nz >= minZ && nz <= maxZ {

				visited[createKey(nx, ny, nz)] = true
				queue = append(queue, []int{nx, ny, nz})
			}
		}

	}

	return result
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
