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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func coloredString(text string, color int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

func mustAtoi(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func intersect(a, b *Line) []float64 {
	x1, y1 := a.from[0], a.from[1]
	x2, y2 := a.to[0], a.to[1]
	x3, y3 := b.from[0], b.from[1]
	x4, y4 := b.to[0], b.to[1]

	denominator := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	if denominator == 0 {
		return nil
	}

	ua := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denominator
	if ua < 0 || ua > 1 {
		return nil
	}
	ub := ((x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)) / denominator
	if ub < 0 || ub > 1 {
		return nil
	}

	x := x1 + ua*(x2-x1)
	y := y1 + ua*(y2-y1)
	return []float64{x, y}
}

type Line struct {
	from []float64
	to   []float64
}

func part1(input string) int {
	textLines := strings.Split(input, "\n")

	var lines []*Line
	for _, l := range textLines {
		l = strings.Replace(l, " @", ",", 1)
		stringNumbers := strings.Split(l, ", ")
		var numbers []float64
		for _, stringNumber := range stringNumbers {
			float, _ := strconv.ParseFloat(stringNumber, 64)
			numbers = append(numbers, float)

		}

		a := []float64{numbers[0], numbers[1]}
		b := []float64{numbers[0] + numbers[3]*10e20, numbers[1] + numbers[4]*10e20}

		//fmt.Println(a, b)
		lines = append(lines, &Line{
			from: a,
			to:   b,
		})
	}

	ans := 0
	least := float64(200000000000000)
	most := float64(400000000000000)
	//least := float64(7)
	//most := float64(27)
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			result := intersect(lines[i], lines[j])
			if result != nil && least < result[0] && result[0] < most && least < result[1] && result[1] < most {
				ans += 1
			}
		}
	}

	// 2701 too low
	return ans
}

func part2(input string) int {
	return 0
}

func main() {

	// 6668 not correct
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
