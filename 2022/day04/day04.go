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

func createIntSet(fromTo []string) map[int]bool {

	from, _ := strconv.Atoi(fromTo[0])
	to, _ := strconv.Atoi(fromTo[1])

	set := make(map[int]bool)
	for i := from; i <= to; i++ {
		set[i] = true
	}
	return set
}

func intersection(a map[int]bool, b map[int]bool) int {
	if len(a) > len(b) {
		a, b = b, a
	}
	for key := range a {
		if b[key] {
			return key
		}
	}
	return -1
}

func intersectionSet(a map[int]bool, b map[int]bool) map[int]bool {
	set_intersection := map[int]bool{}
	if len(a) > len(b) {
		a, b = b, a
	}
	for key := range a {
		if b[key] {
			set_intersection[key] = true
		}
	}
	return set_intersection
}

func equalSet(a map[int]bool, b map[int]bool) bool {
	if len(a) == len(b) {
		for keyA := range a {
			if _, ok := b[keyA]; !ok {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func part1(input string) int {
	result := 0

	for _, l := range strings.Split(input, "\n") {

		sections := strings.Split(l, ",")
		first := strings.Split(sections[0], "-")
		second := strings.Split(sections[1], "-")

		firstSet := createIntSet(first)
		secondSet := createIntSet(second)

		intersection := intersectionSet(firstSet, secondSet)

		if len(firstSet) > len(secondSet) {
			if equalSet(secondSet, intersection) {
				result += 1
			}
		} else {
			if equalSet(firstSet, intersection) {
				result += 1
			}
		}

	}
	return result
}

func part2(input string) int {
	result := 0

	for _, l := range strings.Split(input, "\n") {

		sections := strings.Split(l, ",")
		first := strings.Split(sections[0], "-")
		second := strings.Split(sections[1], "-")

		firstSet := createIntSet(first)
		secondSet := createIntSet(second)

		intersection := intersectionSet(firstSet, secondSet)

		if len(intersection) > 0 {
			result += 1
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
