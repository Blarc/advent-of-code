package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func part1(input string) int {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		println(l)
		columnSplit := strings.Split(l, ":")
		gameSplit := strings.Split(columnSplit[0], " ")
		gameId, _ := strconv.Atoi(gameSplit[1])

		println(gameId)
		sets := strings.Split(strings.TrimSpace(columnSplit[1]), ";")
		invalid := false
		for i := range sets {
			cubesMap := make(map[string]int)

			cubes := strings.Split(strings.TrimSpace(sets[i]), ",")
			for j := range cubes {
				cube := strings.TrimSpace(cubes[j])
				cubeColor := strings.Split(cube, " ")[1]
				cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				cubesMap[cubeColor] += cubeCount
			}

			// only 12 red cubes, 13 green cubes, and 14 blue cubes?
			println(cubesMap["red"], cubesMap["green"], cubesMap["blue"])
			if cubesMap["red"] > 12 || cubesMap["green"] > 13 || cubesMap["blue"] > 14 {
				invalid = true
				println("invalid")
				break
			}
		}

		if !invalid {
			sum += gameId
		}

	}

	// 289
	return sum
}

func part2(input string) int {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		println(l)
		columnSplit := strings.Split(l, ":")
		gameSplit := strings.Split(columnSplit[0], " ")
		gameId, _ := strconv.Atoi(gameSplit[1])

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		println(gameId)
		sets := strings.Split(strings.TrimSpace(columnSplit[1]), ";")
		for i := range sets {
			cubesMap := make(map[string]int)

			cubes := strings.Split(strings.TrimSpace(sets[i]), ",")
			for j := range cubes {
				cube := strings.TrimSpace(cubes[j])
				cubeColor := strings.Split(cube, " ")[1]
				cubeCount, _ := strconv.Atoi(strings.Split(cube, " ")[0])
				cubesMap[cubeColor] += cubeCount
			}

			//// only 12 red cubes, 13 green cubes, and 14 blue cubes?
			//println(cubesMap["red"], cubesMap["green"], cubesMap["blue"])
			//if cubesMap["red"] > 12 || cubesMap["green"] > 13 || cubesMap["blue"] > 14 {
			//	invalid = true
			//	println("invalid")
			//	break
			//}
			if cubesMap["red"] > maxRed {
				maxRed = cubesMap["red"]
			}
			if cubesMap["blue"] > maxBlue {
				maxBlue = cubesMap["blue"]
			}
			if cubesMap["green"] > maxGreen {
				maxGreen = cubesMap["green"]
			}
		}

		println(maxRed, maxBlue, maxGreen)
		println(maxRed * maxBlue * maxGreen)
		sum += maxRed * maxBlue * maxGreen
	}

	// 289
	return sum
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
