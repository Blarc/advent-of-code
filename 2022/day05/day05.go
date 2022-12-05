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

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func part1(input string) string {
	movements := false

	lines := strings.Split(input, "\n")
	numberOfColumns := int((len(lines[0]) + 1) / 4)
	columns := make([]string, numberOfColumns)

	for _, line := range lines {

		if len(line) == 0 || strings.HasPrefix(line, " 1") {
			// Movements
			movements = true
			continue
		}

		if !movements {
			// Starting positions
			for column := 0; column < numberOfColumns; column += 1 {
				pos := column*4 + 1
				// fmt.Println(row, column, ":", string(line[pos]))
				if string(line[pos]) != " " {
					columns[column] += string(line[pos])
				}
			}

		} else {
			// Movements
			split := strings.Split(line, " ")
			amount, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])

			// fmt.Println("amount: ", amount, "from: ", from, "to: ", to)
			columns[to-1] = reverse(columns[from-1][:amount]) + columns[to-1]
			columns[from-1] = columns[from-1][amount:]

		}
	}

	result := ""
	for _, column := range columns {
		result += string(column[0])
	}

	return result
}

func part2(input string) string {
	movements := false

	lines := strings.Split(input, "\n")
	numberOfColumns := int((len(lines[0]) + 1) / 4)
	columns := make([]string, numberOfColumns)

	for _, line := range lines {

		if len(line) == 0 || strings.HasPrefix(line, " 1") {
			// Movements
			movements = true
			continue
		}

		if !movements {
			// Starting positions
			for column := 0; column < numberOfColumns; column += 1 {
				pos := column*4 + 1
				// fmt.Println(row, column, ":", string(line[pos]))
				if string(line[pos]) != " " {
					columns[column] += string(line[pos])
				}
			}

		} else {
			// Movements
			split := strings.Split(line, " ")
			amount, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])

			// fmt.Println("amount: ", amount, "from: ", from, "to: ", to)
			columns[to-1] = columns[from-1][:amount] + columns[to-1]
			columns[from-1] = columns[from-1][amount:]

		}
	}
	fmt.Println(columns)

	result := ""
	for _, column := range columns {
		result += string(column[0])
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
