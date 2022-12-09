package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isFar(a []int, b []int) bool {
	return abs(a[0]-b[0]) > 1 || abs(a[1]-b[1]) > 1
}

func isDiag(a []int, b []int) bool {
	return abs(a[0]-b[0]) > 1 && abs(a[1]-b[1]) == 1 ||
		abs(a[0]-b[0]) == 1 && abs(a[1]-b[1]) > 1
}

func part1(input string) int {

	visited := make(map[string]bool)

	// x, y
	head := []int{0, 0}
	tail := []int{0, 0}

	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")

		direction := s[0]
		amount, _ := strconv.Atoi(s[1])

		fmt.Println(direction, amount)
		for i := 0; i < amount; i++ {

			previousHead := []int{head[0], head[1]}

			if direction == "R" {
				head[0] += 1
			} else if direction == "U" {
				head[1] += 1
			} else if direction == "L" {
				head[0] -= 1
			} else if direction == "D" {
				head[1] -= 1
			} else {
				panic("Invalid direction!")
			}

			tailX := strconv.Itoa(tail[0])
			tailY := strconv.Itoa(tail[1])

			if isFar(tail, head) {

				fmt.Println(tailX, ",", tailY)

				tail[0] = previousHead[0]
				tail[1] = previousHead[1]
			} else {
				fmt.Println(tailX, ",", tailY)
			}

			visited[tailX+","+tailY] = true

		}

	}
	return len(visited)
}

func part2(input string) int {
	visited := make(map[string]bool)

	// x, y
	snakeLength := 10
	snake := make([][]int, snakeLength)
	for i := 0; i < len(snake); i++ {
		snake[i] = []int{0, 0}
	}

	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")

		direction := s[0]
		amount, _ := strconv.Atoi(s[1])

		fmt.Println(direction, amount)
		for i := 0; i < amount; i++ {

			// previousHead := snake[0]

			if direction == "R" {
				snake[0][0] += 1
			} else if direction == "U" {
				snake[0][1] += 1
			} else if direction == "L" {
				snake[0][0] -= 1
			} else if direction == "D" {
				snake[0][1] -= 1
			} else {
				panic("Invalid direction!")
			}

			for j := 1; j < len(snake); j++ {
				if snake[j][0]+2 == snake[j-1][0] && snake[j][1] == snake[j-1][1] {
					// Right
					snake[j][0] += 1
				} else if snake[j][0]-2 == snake[j-1][0] && snake[j][1] == snake[j-1][1] {
					// Left
					snake[j][0] -= 1
				} else if snake[j][0] == snake[j-1][0] && snake[j][1]+2 == snake[j-1][1] {
					// Up
					snake[j][1] += 1
				} else if snake[j][0] == snake[j-1][0] && snake[j][1]-2 == snake[j-1][1] {
					// Down
					snake[j][1] -= 1
				} else if snake[j][0]+1 <= snake[j-1][0] && snake[j][1]+2 <= snake[j-1][1] ||
					snake[j][0]+2 <= snake[j-1][0] && snake[j][1]+1 <= snake[j-1][1] {
					// Right up
					snake[j][0] += 1
					snake[j][1] += 1
				} else if snake[j][0]+1 <= snake[j-1][0] && snake[j][1]-2 >= snake[j-1][1] ||
					snake[j][0]+2 <= snake[j-1][0] && snake[j][1]-1 >= snake[j-1][1] {
					// Right down
					snake[j][0] += 1
					snake[j][1] -= 1
				} else if snake[j][0]-1 >= snake[j-1][0] && snake[j][1]+2 <= snake[j-1][1] ||
					snake[j][0]-2 >= snake[j-1][0] && snake[j][1]+1 <= snake[j-1][1] {
					// Left up
					snake[j][0] -= 1
					snake[j][1] += 1
				} else if snake[j][0]-1 >= snake[j-1][0] && snake[j][1]-2 >= snake[j-1][1] ||
					snake[j][0]-2 >= snake[j-1][0] && snake[j][1]-1 >= snake[j-1][1] {
					// Left down
					snake[j][0] -= 1
					snake[j][1] -= 1
				}
			}

			tailX := strconv.Itoa(snake[snakeLength-1][0])
			tailY := strconv.Itoa(snake[snakeLength-1][1])

			for index, node := range snake {
				fmt.Println(index, node)
			}
			fmt.Println("---")

			visited[tailX+","+tailY] = true

		}

	}
	return len(visited)
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
