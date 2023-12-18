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

type Node struct {
	y, x  int
	color string
}

func coloredString(text string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

//func part1(input string) int {
//	lines := strings.Split(input, "\n")
//
//	m := make(map[int]map[int]*Node)
//
//	minY, minX := 0, 0
//	maxY, maxX := 0, 0
//	currentY, currentX := 0, 0
//	for _, l := range lines {
//		s := strings.Split(l, " ")
//		dir := s[0]
//		amount, _ := strconv.Atoi(s[1])
//		color := s[2]
//		for i := 0; i < amount; i++ {
//			if m[currentY] == nil {
//				m[currentY] = make(map[int]*Node)
//			}
//
//			m[currentY][currentX] = &Node{currentY, currentX, color}
//
//			if currentY > maxY {
//				maxY = currentY
//			}
//			if currentX > maxX {
//				maxX = currentX
//			}
//
//			if currentY < minY {
//				minY = currentY
//			}
//
//			if currentX < minX {
//				minX = currentX
//			}
//
//			if dir == "R" {
//				currentX += 1
//			} else if dir == "L" {
//				currentX -= 1
//			} else if dir == "D" {
//				currentY += 1
//			} else if dir == "U" {
//				currentY -= 1
//			} else {
//				panic("No matches!")
//			}
//		}
//	}
//
//	fmt.Println(minY, minX)
//	fmt.Println(maxY, maxX)
//
//	sum := 0
//	visited := make(map[int]map[int]bool)
//	visited[1] = make(map[int]bool)
//	visited[1][1] = true
//
//	q := []int{1, 1}
//	for len(q) > 0 {
//		y, x := q[0], q[1]
//		q = q[2:]
//		sum += 1
//
//		if y < minY || maxY < y || x < minX || maxX < x {
//			panic("OUT OF BOUNDS")
//		}
//
//		dirs := []int{-1, 0, 1, 0, 0, -1, 0, 1}
//		for i := 0; i < len(dirs); i += 2 {
//			dy := dirs[i]
//			dx := dirs[i+1]
//
//			nextY := y + dy
//			nextX := x + dx
//			_, nextIsWall := m[nextY][nextX]
//			if !nextIsWall && !visited[nextY][nextX] {
//				if visited[nextY] == nil {
//					visited[nextY] = make(map[int]bool)
//				}
//				visited[nextY][nextX] = true
//				q = append(q, nextY)
//				q = append(q, nextX)
//			}
//		}
//	}
//
//	for y := minY; y <= maxY; y++ {
//		for x := minX; x <= maxX; x++ {
//			_, exists := m[y][x]
//			if exists {
//				sum += 1
//			}
//		}
//	}
//
//	//for y := minY; y <= maxY; y++ {
//	//	for x := minX; x <= maxX; x++ {
//	//		if _, exists := m[y][x]; exists {
//	//			fmt.Print("#")
//	//		} else if _, ok := visited[y][x]; ok {
//	//			fmt.Print("*")
//	//			//fmt.Print(coloredString("*", 94))
//	//		} else {
//	//			fmt.Print(".")
//	//			//fmt.Print(coloredString(".", 92))
//	//		}
//	//	}
//	//	fmt.Println()
//	//}
//
//	return sum
//}

func part1(input string) int {

	lines := strings.Split(input, "\n")

	var points [][]int
	edge := 0

	currentY, currentX := 0, 0
	for _, l := range lines {
		s := strings.Split(l, " ")
		dir := s[0]
		amountInt32, _ := strconv.Atoi(s[1])
		edge += amountInt32

		points = append(points, []int{currentY, currentX})

		if dir == "R" {
			currentX += amountInt32
		} else if dir == "L" {
			currentX -= amountInt32
		} else if dir == "D" {
			currentY += amountInt32
		} else if dir == "U" {
			currentY -= amountInt32
		} else {
			panic("No matches!")
		}
	}

	return edge/2 + computeArea(points) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeArea(corners [][]int) int {
	n := len(corners)
	area := 0

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += corners[i][0] * corners[j][1]
		area -= corners[j][0] * corners[i][1]
	}

	area = abs(area) / 2
	return area
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	var points [][]int
	edge := 0

	minY, minX := 0, 0
	maxY, maxX := 0, 0
	currentY, currentX := 0, 0
	for _, l := range lines {
		s := strings.Split(l, " ")

		color := s[2][2 : len(s[2])-1]
		dir := color[len(color)-1]
		amount, _ := strconv.ParseInt(color[:len(color)-1], 16, 32)
		amountInt32 := int(amount)
		edge += amountInt32

		points = append(points, []int{currentY, currentX})

		if currentY > maxY {
			maxY = currentY
		}
		if currentX > maxX {
			maxX = currentX
		}

		if currentY < minY {
			minY = currentY
		}

		if currentX < minX {
			minX = currentX
		}

		if dir == '0' {
			currentX += amountInt32
		} else if dir == '2' {
			currentX -= amountInt32
		} else if dir == '1' {
			currentY += amountInt32
		} else if dir == '3' {
			currentY -= amountInt32
		} else {
			panic("No matches!")
		}
	}

	return edge/2 + computeArea(points) + 1

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
