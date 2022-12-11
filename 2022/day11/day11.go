package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func parseItems(itemLine string) []int {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(itemLine, -1)

	items := make([]int, len(matches))
	for i, item := range matches {
		item := strings.TrimSpace(item)
		intItem, _ := strconv.Atoi(item)
		items[i] = intItem
	}
	return items
}

func parseOperation(opLine string) func(int) int {
	opLine = strings.Split(opLine, "=")[1]
	opLine = strings.TrimSpace(opLine)
	s := strings.Split(opLine, " ")

	second := s[2]
	secondNumber, err := strconv.Atoi(second)

	if err == nil {
		if s[1] == "*" {
			return func(old int) int {
				return old * secondNumber
			}
		} else if s[1] == "+" {
			return func(old int) int {
				return old + secondNumber
			}
		} else {
			panic("Invalid operation.")
		}

	} else {
		if s[1] == "*" {
			return func(old int) int {
				return old * old
			}
		} else if s[1] == "+" {
			return func(old int) int {
				return old + old
			}
		} else {
			panic("Invalid operation.")
		}
	}
}

func parseTest(testLines []string) func(int) []int {

	re := regexp.MustCompile("[0-9]+")
	divisibleString := re.FindAllString(testLines[0], -1)[0]
	divisible, _ := strconv.Atoi(divisibleString)
	// fmt.Println("divisible:", divisible)

	trueMonkeyString := re.FindAllString(testLines[1], -1)[0]
	trueMonkey, _ := strconv.Atoi(trueMonkeyString)
	// fmt.Println("trueMonkey:", trueMonkey)

	falseMonkeyString := re.FindAllString(testLines[2], -1)[0]
	falseMonkey, _ := strconv.Atoi(falseMonkeyString)

	return func(new int) []int {
		if new%divisible == 0 {
			return []int{new, trueMonkey}
		}
		return []int{new, falseMonkey}
	}

}

type Monkey struct {
	items     []int
	operation func(int) int
	test      func(int) []int
	inspected int
}

func part1(input string) int {

	lines := strings.Split(input, "\n")
	var monkeys []Monkey

	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if strings.HasPrefix(l, "Monkey") {
			items := parseItems(lines[i+1])
			operation := parseOperation(lines[i+2])
			test := parseTest(lines[i+3 : i+6])
			monkeys = append(monkeys, Monkey{items, operation, test, 0})
		}
	}

	for round := 0; round < 20; round++ {
		for j, monkey := range monkeys {
			numberOfMonkeyItems := len(monkey.items)
			for i := 0; i < numberOfMonkeyItems; i++ {
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				// fmt.Println(item, monkey.items)
				item = monkey.operation(item) / 3
				passTo := monkey.test(item)[1]

				monkeys[passTo].items = append(monkeys[passTo].items, item)
				monkey.inspected += 1

			}
			monkeys[j] = monkey
		}
	}

	for _, monkey := range monkeys {
		fmt.Println(monkey.items, monkey.inspected)
	}
	return 1
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	var monkeys []Monkey
	var praNumbers []int

	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if strings.HasPrefix(l, "Monkey") {
			items := parseItems(lines[i+1])
			operation := parseOperation(lines[i+2])
			test := parseTest(lines[i+3 : i+6])

			re := regexp.MustCompile("[0-9]+")
			divisibleString := re.FindAllString(lines[i+3], -1)[0]
			divisible, _ := strconv.Atoi(divisibleString)
			praNumbers = append(praNumbers, divisible)

			monkeys = append(monkeys, Monkey{items, operation, test, 0})
		}
	}

	fmt.Println(praNumbers)
	praNumberProduct := 1
	for _, praNumber := range praNumbers {
		praNumberProduct *= praNumber
	}

	for round := 0; round < 10000; round++ {
		for j, monkey := range monkeys {
			numberOfMonkeyItems := len(monkey.items)
			for i := 0; i < numberOfMonkeyItems; i++ {
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				// fmt.Println(item, monkey.items)
				item = monkey.operation(item)
				passTo := monkey.test(item)
				item = passTo[0]

				item = item % praNumberProduct

				// fmt.Println("Passing", item, "to", passTo[1])

				monkeys[passTo[1]].items = append(monkeys[passTo[1]].items, item)
				monkey.inspected += 1

			}
			monkeys[j] = monkey
		}
		/*
			for _, monkey := range monkeys {
				fmt.Printf("%5d", monkey.inspected)
				fmt.Print(" ")
			}
			fmt.Println()
		*/
	}

	var result []int
	for _, monkey := range monkeys {

		result = append(result, monkey.inspected)
	}

	sort.Ints(result)
	fmt.Println(result)
	return result[len(result)-1] * result[len(result)-2]
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
