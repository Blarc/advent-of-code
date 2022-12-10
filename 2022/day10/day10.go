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

func part1(input string) int {

	cycles := 0
	x := 1
	result := 0
	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")
		command := s[0]

		fmt.Println(l)
		if command == "addx" {
			amount, _ := strconv.Atoi(s[1])
			for i := 0; i < 2; i++ {
				cycles += 1
				if cycles == 20 || (cycles+20)%40 == 0 {
					fmt.Println(cycles, x, x*cycles)
					result += x * cycles
				}
			}
			x += amount
			fmt.Println(amount, "=", x)
		} else if command == "noop" {
			cycles += 1
			if cycles == 20 || (cycles+20)%40 == 0 {
				fmt.Println(cycles, x, x*cycles)
				result += x * cycles
			}
		} else {
			panic(command)
		}

	}
	return result
}

func part2(input string) int {
	cycles := 0
	x := 1
	for _, l := range strings.Split(input, "\n") {

		s := strings.Split(l, " ")
		command := s[0]

		if command == "addx" {
			amount, _ := strconv.Atoi(s[1])
			for i := 0; i < 2; i++ {

				if x-1 <= cycles && cycles <= x+1 {
					fmt.Printf("#")
				} else {
					fmt.Printf(" ")
				}

				cycles += 1

				if cycles != 0 && cycles%40 == 0 {
					// fmt.Println(cycles, x, x*cycles)
					// fmt.Print(cycles)
					cycles -= 40
					fmt.Println()
				}
			}
			x += amount
		} else if command == "noop" {

			if x-1 <= cycles && cycles <= x+1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}

			cycles += 1

			if cycles != 0 && cycles%40 == 0 {
				// fmt.Println(cycles, x, x*cycles)
				// fmt.Print(cycles)
				cycles -= 40
				fmt.Println()
			}
		} else {
			panic(command)
		}

	}
	return 2
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
