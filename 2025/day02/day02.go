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

func part1(input string) int {
	sum := 0

	for _, l := range strings.Split(input, ",") {
		r := strings.Split(l, "-")
		firstId, _ := strconv.Atoi(r[0])
		lastId, _ := strconv.Atoi(r[1])

		for id := firstId; id <= lastId; id++ {
			stringId := strconv.Itoa(id)
			// Split in the middle
			middle := len(stringId) / 2
			firstPart := stringId[:middle]
			lastPart := stringId[middle:]

			// fmt.Printf("id: %s, first: %s, last: %s\n", stringId, firstPart, lastPart)
			if firstPart == lastPart {
				sum += id
			}
		}
	}
	return sum
}

func invalidId(first, last string) bool {
	for i := 0; i < len(last); i += len(first) {
		if i+len(first) > len(last) {
			return false
		}

		if last[i:i+len(first)] != first {
			return false
		}
	}
	return true
}

func part2(input string) int {
	sum := 0

	for _, l := range strings.Split(input, ",") {
		r := strings.Split(l, "-")
		firstId, _ := strconv.Atoi(r[0])
		lastId, _ := strconv.Atoi(r[1])

		for id := firstId; id <= lastId; id++ {
			stringId := strconv.Itoa(id)

			for i := 0; i < len(stringId)/2; i++ {
				firstPart := stringId[:i+1]
				lastPart := stringId[i+1:]

				// fmt.Printf("firstPart: %s, lastPart: %s\n", firstPart, lastPart)
				if invalidId(firstPart, lastPart) {
					// fmt.Printf("Invalid id: %d\n", id)
					sum += id
					break
				}
			}
		}
	}
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
