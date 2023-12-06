package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func part1(input string) int {
	lines := strings.Split(input, "\n")

	numberRe := regexp.MustCompile(`\d+`)
	timeS := numberRe.FindAllString(lines[0], -1)
	t := make([]int, len(timeS))
	for i := range timeS {
		t[i], _ = strconv.Atoi(timeS[i])
	}
	distanceS := numberRe.FindAllString(lines[1], -1)
	d := make([]int, len(distanceS))
	for i := range distanceS {
		d[i], _ = strconv.Atoi(distanceS[i])
	}

	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", d)

	// x * (7 - x) > 9
	ans := 1
	for i := 0; i < len(t); i++ {
		ints := f(t[i], d[i])
		fmt.Printf("%d %v\n", t[i], ints)
		ans *= len(ints)

	}
	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	numberRe := regexp.MustCompile(`\d+`)
	timeS := numberRe.FindAllString(lines[0], -1)
	ts := ""
	for i := range timeS {
		ts += timeS[i]
	}
	t, _ := strconv.Atoi(ts)

	distanceS := numberRe.FindAllString(lines[1], -1)
	ds := ""
	for i := range distanceS {
		ds += distanceS[i]
	}
	d, _ := strconv.Atoi(ds)

	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", d)

	// x * (7 -x) > 9

	imin := fmin(t, d)
	imax := fmax(t, d)

	fmt.Printf("%d %d\n", imin, imax)

	// len(f(t, d)) this works as well... O.o
	return imax - imin + 1
}

func f(y, z int) []int {
	rs := make([]int, 0)
	for x := 0; x < y; x++ {
		r := x * (y - x)
		if r > z {
			rs = append(rs, x)
		}
	}
	return rs
}

func fmin(y, z int) int {
	for x := 0; x < y; x++ {
		r := x * (y - x)
		if r > z {
			return x
		}
	}
	return 0
}

func fmax(y, z int) int {
	for x := y; x > 0; x-- {
		r := x * (y - x)
		if r > z {
			return x
		}
	}
	return 0
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
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))
}
