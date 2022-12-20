package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func mod(x, y int) int {
	for x < 0 {
		x += y
	}
	return x % y
}

func part1(input string) int {

	lines := strings.Split(input, "\n")
	a := make([][]int, len(lines))

	zeroIndex := 0
	for i, l := range lines {
		number, _ := strconv.Atoi(l)
		a[i] = []int{number % (len(lines) - 1), i, number}
		if number == 0 {
			zeroIndex = i
		}
	}

	for i := 0; i < len(a); i++ {
		// fmt.Println("moving", a[i][0], a[i][1])
		val := a[i][0]
		prev := a[i][1]

		corrector := 0

		if prev+val <= 0 {
			val--
		} else if prev+val >= len(a)-1 {
			val++
		}

		/*

			if val < 0 && prev+val < 0 {
				// back and underflow
				val--
				corrector = -1

			} else if val > 0 && prev+val >= len(a) {
				// forward and overflow
				val++
				corrector = 1
			} else if val > 0 {
				// forward
				corrector = -1
			}
		*/

		new := mod((prev + val), len(a))
		a[i][1] = new

		if prev > new {
			tmp := prev
			prev = new
			new = tmp
			corrector = 1
		} else {
			corrector = -1
		}

		for j := 0; j < len(a); j++ {
			if prev <= a[j][1] && a[j][1] <= new && i != j {
				a[j][1] = mod((a[j][1] + corrector), len(a))
			}
		}

	}

	// 1 2 -3 3 -2 0 4
	// 0 1  2 3  4 5 6

	// 1 2 -3 3 -2 0 4
	// 1 0  2 3  4 5 6

	// 1 2 -3 3 -2 0 4
	// 0 2  1 3  4 5 6

	// 1 2 -3 3 -2 0 4
	// 0 1  4 2  3 5 6

	// 1 2 -3 3 -2 0 4
	// 0 1  3 5  2 4 6

	// 1 2 -3 3 -2 0 4
	// 0 1  2 4  6 3 5

	// 1 2 -3 3 -2 0 4
	// 0 1  2 4  6 3 5

	// -2 0  4 1 -3 2 3
	// -2 -3 0 4  1 2 3

	fmt.Println(zeroIndex)
	fmt.Println(a[zeroIndex])
	result := 0
	for i := 1; i < 4; i++ {
		x := mod(a[zeroIndex][1]+i*1000, len(a))
		fmt.Println(x)

		for j := 0; j < len(a); j++ {
			if x == a[j][1] {
				fmt.Println(a[j])
				result += a[j][2]
			}
		}
	}

	test := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		test[a[i][1]] = true
	}

	// fmt.Println(a)
	fmt.Println(len(test) == len(a))

	// 2008 too low
	return result
}

func part2(input string) int {

	lines := strings.Split(input, "\n")
	a := make([][]int, len(lines))

	zeroIndex := 0
	for i, l := range lines {
		number, _ := strconv.Atoi(l)
		a[i] = []int{number * 811589153 % (len(lines) - 1), i, number * 811589153}
		if number == 0 {
			zeroIndex = i
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < len(a); i++ {
			// fmt.Println("moving", a[i][0], a[i][1])
			val := a[i][0]
			prev := a[i][1]

			corrector := 0

			if prev+val <= 0 {
				val--
			} else if prev+val >= len(a)-1 {
				val++
			}

			new := mod((prev + val), len(a))
			a[i][1] = new

			if prev > new {
				tmp := prev
				prev = new
				new = tmp
				corrector = 1
			} else {
				corrector = -1
			}

			for j := 0; j < len(a); j++ {
				if prev <= a[j][1] && a[j][1] <= new && i != j {
					a[j][1] = mod((a[j][1] + corrector), len(a))
				}
			}

		}

	}
	result := 0

	fmt.Println(zeroIndex)
	fmt.Println(a[zeroIndex])
	for i := 1; i < 4; i++ {
		x := mod(a[zeroIndex][1]+i*1000, len(a))
		fmt.Println(x)

		for j := 0; j < len(a); j++ {
			if x == a[j][1] {
				fmt.Println(a[j])
				result += a[j][2]
			}
		}
	}

	test := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		test[a[i][1]] = true
	}

	// fmt.Println(a)
	fmt.Println(len(test) == len(a))

	// 2008 too low
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

	start := time.Now()
	if part == 1 {
		fmt.Println("Result:", part1(inputText))
	} else {
		fmt.Println("Result:", part2(inputText))
	}
	fmt.Println("Time:", fmt.Sprintf("%d ms", time.Since(start).Milliseconds()))

}
