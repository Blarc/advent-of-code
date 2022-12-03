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

func createSet(array []rune) map[rune]bool {
	set := make(map[rune]bool)
	for _, item := range array {
		set[item] = true
	}
	return set
}

func intersection(a map[rune]bool, b map[rune]bool) rune {
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

func intersectionSet(a map[rune]bool, b map[rune]bool) map[rune]bool {
	set_intersection := map[rune]bool{}
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

func part1(input string) int {
	result := 0
	for _, l := range strings.Split(input, "\n") {

		half := len(l) / 2
		firstHalf := createSet([]rune(l[:half]))
		secondHalf := createSet([]rune(l[half:]))
		// fmt.Println(firstHalf, " ", secondHalf)
		intersection := intersection(firstHalf, secondHalf)
		intersection_corrected := intersection - rune('a') + 1
		if intersection_corrected < 0 {
			intersection_corrected += 58
		}
		fmt.Println(intersection, intersection_corrected, " ", string(intersection))
		result += int(intersection_corrected)
	}
	return result
}

func part2(input string) int {
	result := 0
	groupSet := map[rune]bool{}
	for i, l := range strings.Split(input, "\n") {

		fmt.Println(i, " ", l)

		if i%3 == 0 {
			groupSet = createSet([]rune(l))
		} else {
			groupSet = intersectionSet(groupSet, createSet([]rune(l)))
		}

		i += 1
		if i%3 == 0 {
			for intersection := range groupSet {
				intersection_corrected := intersection - rune('a') + 1
				if intersection_corrected < 0 {
					intersection_corrected += 58
				}
				fmt.Println(intersection, intersection_corrected, " ", string(intersection))
				result += int(intersection_corrected)
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
