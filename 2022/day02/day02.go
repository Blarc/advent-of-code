package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

// A, X - Rock
// B, Y - Paper
// C, Z - Scissors
var mapper = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var winner = map[string]string{
	"X": "Y",
	"Y": "Z",
	"Z": "X",
}

var loser = map[string]string{
	"X": "Z",
	"Y": "X",
	"Z": "Y",
}

func part1(input string) int {

	result := 0
	for _, l := range strings.Split(input, "\n") {

		var split = strings.Split(l, " ")

		opponent := mapper[split[0]]
		me := split[1]

		fmt.Println(opponent, "vs", me)

		round := points[me]
		if opponent == me {
			round += 3
		} else if winner[opponent] == me {
			round += 6
		}

		fmt.Println(round)
		result += round

	}
	return result
}

func part2(input string) int {

	result := 0
	for _, l := range strings.Split(input, "\n") {

		var split = strings.Split(l, " ")

		opponent := mapper[split[0]]
		round_end := split[1]

		var me string
		if round_end == "X" {
			// lose
			me = loser[opponent]
		} else if round_end == "Y" {
			// draw
			me = opponent
		} else if round_end == "Z" {
			// win
			me = winner[opponent]
		} else {
			panic("Something went wrong!")
		}

		fmt.Println(opponent, "vs", me)
		round := points[me]
		if opponent == me {
			round += 3
		} else if winner[opponent] == me {
			round += 6
		}

		fmt.Println(round)
		result += round

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
