package main

import (
	"flag"
	"fmt"
	"regexp"
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

// 2018/day01
func part1(input string) int {
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		var digits []rune
		for _, c := range l {

			digit := c - 48
			if c-48 < 10 {
				digits = append(digits, digit)
			}

		}
		sum += int(digits[0])*10 + int(digits[len(digits)-1])
	}
	return sum
}

func part2(input string) int {

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	sum := 0
	for _, l := range strings.Split(input, "\n") {
		minIndex := 999999
		minValue := 0
		maxIndex := -1
		maxValue := 0

		re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
		numberIndexes := re.FindAllStringIndex(l, -1)

		for number, _ := range numbers {
			re := regexp.MustCompile(number)
			numberIndexes := re.FindAllStringIndex(l, -1)

			for index := range numberIndexes {
				if numberIndexes[index][0] < minIndex {
					minIndex = numberIndexes[index][0]
					minValue = numbers[l[numberIndexes[index][0]:numberIndexes[index][1]]]
				}
				if numberIndexes[index][0] > maxIndex {
					maxIndex = numberIndexes[index][0]
					maxValue = numbers[l[numberIndexes[index][0]:numberIndexes[index][1]]]
				}
			}
		}

		re = regexp.MustCompile(`[0-9]`)
		numberIndexes = re.FindAllStringIndex(l, -1)

		for index := range numberIndexes {
			//fmt.Println("idx:", numberIndexes[index][0])
			if numberIndexes[index][0] < minIndex {
				minIndex = numberIndexes[index][0]
				atoi, _ := strconv.Atoi(l[numberIndexes[index][0]:numberIndexes[index][1]])
				minValue = atoi
				//println("minValue", l[numberIndexes[index][0]:numberIndexes[index][1]])
			}
			if numberIndexes[index][0] > maxIndex {
				maxIndex = numberIndexes[index][0]
				atoi, _ := strconv.Atoi(l[numberIndexes[index][0]:numberIndexes[index][1]])
				maxValue = atoi
				//println("maxValue", l[numberIndexes[index][0]:numberIndexes[index][1]])
			}
		}

		sum += minValue*10 + maxValue
		println(l, minValue, maxValue, minValue*10+maxValue)

		// 51140
		// 51128

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
