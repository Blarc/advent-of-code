package main

import (
	"flag"
	"fmt"
	"slices"
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

func part1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		for _, s := range strings.Split(l, ",") {
			tmp := hash(s)
			sum += tmp
		}
	}

	return sum
}

func hash(s string) int {
	tmp := 0
	for _, c := range s {
		number := int(c)
		tmp += number
		tmp *= 17
		tmp %= 256
	}
	return tmp
}

type Lens struct {
	label string
	focal int
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	m := make(map[int][]*Lens)

	for _, l := range lines {
		for _, s := range strings.Split(l, ",") {

			if strings.HasSuffix(s, "-") {
				label := strings.TrimSuffix(s, "-")
				box := hash(label)
				m[box] = slices.DeleteFunc(m[box], func(lens *Lens) bool {
					return lens.label == label
				})
			} else {
				tmp := strings.Split(s, "=")
				label := tmp[0]
				box := hash(label)
				focal, _ := strconv.Atoi(tmp[1])

				index := slices.IndexFunc(m[box], func(lens *Lens) bool { return lens.label == label })
				if index == -1 {
					m[box] = append(m[box], &Lens{label, focal})
				} else {
					m[box][index].focal = focal
				}

			}
		}
	}

	sum := 0
	for i, box := range m {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focal
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
