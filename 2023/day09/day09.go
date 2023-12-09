package main

import (
	"flag"
	"fmt"
	"regexp"
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

// -6 -5 -2 12 69 245 687 1656 3603 7306 14121 26447 48580 88239 159197 285682 509691 903614 1594019 2811443 5002638
// 10  13  16  21  30  45
func diff(a []int) int {
	//fmt.Printf("diff: %#v\n", a)
	if len(a) == 2 {
		return a[0] - a[1]
	}

	var b []int
	for i := 0; i < len(a)-1; i++ {
		b = append(b, a[i]-a[i+1])
	}

	return diff(b) + a[0]
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`-?\d+`)

	var histories [][]int
	for _, l := range lines {
		numbers := re.FindAllString(l, -1)
		slices.Reverse(numbers)
		history := make([]int, len(numbers))
		for i, n := range numbers {
			atoi, _ := strconv.Atoi(n)
			history[i] = atoi
		}
		histories = append(histories, history)
	}

	//fmt.Printf("Histories: %#v\n", histories)

	ans := 0
	for _, h := range histories {
		r := diff(h)
		//fmt.Printf("%#v: %d\n", h, r)
		ans += r
	}

	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`-?\d+`)

	var histories [][]int
	for _, l := range lines {
		numbers := re.FindAllString(l, -1)
		history := make([]int, len(numbers))
		for i, n := range numbers {
			atoi, _ := strconv.Atoi(n)
			history[i] = atoi
		}
		histories = append(histories, history)
	}

	//fmt.Printf("Histories: %#v\n", histories)

	ans := 0
	for _, h := range histories {
		r := diff(h)
		//fmt.Printf("%#v: %d\n", h, r)
		ans += r
	}

	return ans

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
