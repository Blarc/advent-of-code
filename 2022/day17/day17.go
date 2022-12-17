package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func max(slice []int) int {
	max := -1
	for _, x := range slice {
		if x > max {
			max = x
		}
	}
	return max
}

func maxY(rock [][]int, x int) int {

	max := -1
	for _, point := range rock {
		if point[1] == x {
			if point[0] > max {
				max = point[0]
			}
		}
	}
	return max
}

func createKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func show(occupied map[string]bool, maxFloor int) {
	for i := maxFloor; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			if j == 0 {
				fmt.Printf("%2d ", i)
			}
			if occupied[createKey(i, j)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func canMove(x, y int, rock [][]int, occupied map[string]bool) bool {
	for _, point := range rock {
		if point[1]+x < 0 || point[1]+x > 6 || occupied[createKey(point[0]+y, point[1]+x)] {
			return false
		}
	}
	return true
}

func part1(input string) int {

	rocks := [][][]int{
		// horizontal line
		{{0, 2}, {0, 3}, {0, 4}, {0, 5}},
		// cross
		{{1, 2}, {0, 3}, {1, 3}, {1, 4}, {2, 3}},
		// L
		{{0, 2}, {0, 3}, {2, 4}, {1, 4}, {0, 4}},
		// vertical line
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		// square
		{{0, 2}, {0, 3}, {1, 2}, {1, 3}},
	}

	occupied := make(map[string]bool)
	for i := 0; i < 7; i++ {
		occupied[createKey(0, i)] = true
	}

	maxFloor := 0
	inputLen := len(input)
	index := 0
	for i := 0; i < 2022; i++ {
		rock := rocks[i%5]

		stop := false
		var j int
		k := 0
		for j = maxFloor + 4; j > 0; j-- {
			if stop {
				break
			}

			if string(input[index%inputLen]) == ">" && canMove(k+1, j, rock, occupied) {
				// fmt.Println("Right")
				k++
			} else if string(input[index%inputLen]) == "<" && canMove(k-1, j, rock, occupied) {
				// fmt.Println("Left")
				k--
			}

			for _, point := range rock {
				y := point[0] + j - 1
				x := point[1] + k

				if occupied[createKey(y, x)] {
					stop = true
					break
				}

			}
			index++
		}
		j++

		for _, point := range rock {
			if maxFloor < point[0]+j {
				maxFloor = point[0] + j
			}
			occupied[createKey(point[0]+j, point[1]+k)] = true
		}
	}

	//show(occupied, maxFloor+3)

	return maxFloor
}

func part2(input string) int {
	rocks := [][][]int{
		// horizontal line
		{{0, 2}, {0, 3}, {0, 4}, {0, 5}},
		// cross
		{{1, 2}, {0, 3}, {1, 3}, {1, 4}, {2, 3}},
		// L
		{{0, 2}, {0, 3}, {2, 4}, {1, 4}, {0, 4}},
		// vertical line
		{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		// square
		{{0, 2}, {0, 3}, {1, 2}, {1, 3}},
	}

	occupied := make(map[string]bool)
	for i := 0; i < 7; i++ {
		occupied[createKey(0, i)] = true
	}

	maxFloor := 0
	inputLen := len(input)
	index := 0
	previous := 0
	preivousIndex := 0

	for i := 0; i < 3600; i++ {

		saved := 0

		rock := rocks[i%5]

		stop := false
		var j int
		k := 0
		for j = maxFloor + 4; j > 0; j-- {
			if index%inputLen == 0 {
				// fmt.Println(maxFloor - previous)
				saved += maxFloor - previous
				previous = maxFloor
			}
			if stop {
				break
			}

			if string(input[index%inputLen]) == ">" && canMove(k+1, j, rock, occupied) {
				// fmt.Println("Right")
				k++
			} else if string(input[index%inputLen]) == "<" && canMove(k-1, j, rock, occupied) {
				// fmt.Println("Left")
				k--
			}

			for _, point := range rock {
				y := point[0] + j - 1
				x := point[1] + k

				if occupied[createKey(y, x)] {
					stop = true
					break
				}

			}
			index++
		}
		j++

		for _, point := range rock {
			if maxFloor < point[0]+j {
				maxFloor = point[0] + j
			}
			occupied[createKey(point[0]+j, point[1]+k)] = true
		}

		if saved > 0 {
			fmt.Println(i-preivousIndex, saved)
			preivousIndex = i
		}

	}

	// show(occupied, maxFloor+3)
	fmt.Println(588235292 * 2654 + 5601)
	return maxFloor
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
