package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"
	"unicode"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

type Number struct {
	value int
	used  bool
}

type Symbol struct {
	x, y   int
	symbol string
}

func coordToStr(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func leftNumber(x, y int, numberMap map[string]*Number) *Number {
	coords := coordToStr(x-1, y)
	if number, ok := numberMap[coords]; ok {
		return number
	}
	return nil
}

func part1(input string) int {
	sum := 0
	numberMap := make(map[string]*Number)
	symbolMap := make(map[string]*Symbol)
	for y, l := range strings.Split(input, "\n") {
		for x, c := range l {
			if unicode.IsDigit(c) {
				currentDigit := int(c - '0')

				leftNumber := leftNumber(x, y, numberMap)
				if leftNumber != nil {
					leftNumber.value = leftNumber.value*10 + currentDigit
					numberMap[coordToStr(x, y)] = leftNumber
				} else {
					numberMap[coordToStr(x, y)] = &Number{value: currentDigit}
				}

			} else if c != '.' {
				symbolMap[coordToStr(x, y)] = &Symbol{x, y, string(c)}
			}
		}
	}

	//for y := 0; y < 10; y++ {
	//	for x := 0; x < 10; x++ {
	//		coords := coordToStr(x, y)
	//		if number, ok := numberMap[coords]; ok {
	//			println(coords, number.value)
	//		}
	//	}
	//}

	for _, symbol := range symbolMap {
		for y := symbol.y - 1; y <= symbol.y+1; y++ {
			for x := symbol.x - 1; x <= symbol.x+1; x++ {
				coords := coordToStr(x, y)
				if number, ok := numberMap[coords]; ok && !number.used {
					sum += number.value
					number.used = true
				}
			}
		}
	}

	return sum
}

func part2(input string) int {
	sum := 0
	numberMap := make(map[string]*Number)
	symbolMap := make(map[string]*Symbol)
	for y, l := range strings.Split(input, "\n") {
		for x, c := range l {
			if unicode.IsDigit(c) {
				currentDigit := int(c - '0')

				leftNumber := leftNumber(x, y, numberMap)
				if leftNumber != nil {
					leftNumber.value = leftNumber.value*10 + currentDigit
					numberMap[coordToStr(x, y)] = leftNumber
				} else {
					numberMap[coordToStr(x, y)] = &Number{value: currentDigit}
				}

			} else if c != '.' {
				symbolMap[coordToStr(x, y)] = &Symbol{x, y, string(c)}
			}
		}
	}

	for _, symbol := range symbolMap {
		var numbers []Number
		for y := symbol.y - 1; y <= symbol.y+1; y++ {
			for x := symbol.x - 1; x <= symbol.x+1; x++ {
				coords := coordToStr(x, y)
				if number, ok := numberMap[coords]; ok && !number.used {
					numbers = append(numbers, *number)
					number.used = true
				}
			}
		}
		if len(numbers) == 2 && symbol.symbol == "*" {
			sum += numbers[0].value * numbers[1].value
			// println(numbers[0].value, "*", numbers[1].value, "=", numbers[0].value*numbers[1].value)
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
