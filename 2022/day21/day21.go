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

func varsExist(op []string, vars map[string]int) bool {
	for _, x := range op {
		_, exists := vars[x]
		if !exists {
			return false
		}
	}
	return true
}

func part1(input string) int {
	numberRe := regexp.MustCompile("[0-9]+")
	varRe := regexp.MustCompile("[a-z]+")

	vars := make(map[string]int)

	stack := [][]string{}

	for _, l := range strings.Split(input, "\n") {

		varNames := varRe.FindAllString(l, -1)

		if len(varNames) > 1 {
			op := string(l[11])
			varNames = append(varNames, op)
			stack = append(stack, varNames)

		} else {
			x, _ := strconv.Atoi(numberRe.FindString(l))
			vars[varNames[0]] = x
		}
	}

	for len(stack) > 0 {
		// Top element
		n := len(stack) - 1
		op := stack[n]
		// Pop
		stack = stack[:n]

		// fmt.Println(vars, op[1:len(op)-1])
		if varsExist(op[1:len(op)-1], vars) {
			if op[3] == "+" {
				vars[op[0]] = vars[op[1]] + vars[op[2]]
			} else if op[3] == "-" {
				vars[op[0]] = vars[op[1]] - vars[op[2]]
			} else if op[3] == "*" {
				vars[op[0]] = vars[op[1]] * vars[op[2]]
			} else if op[3] == "/" {
				vars[op[0]] = vars[op[1]] / vars[op[2]]
			} else {
				panic("Invalid operation!")
			}
			// fmt.Println(op, vars[op[1]], op[3], vars[op[2]], vars[op[0]])
		} else {
			stack = append([][]string{op}, stack...)
		}
	}

	return vars["root"]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(input string) int {
	numberRe := regexp.MustCompile("[0-9]+")
	varRe := regexp.MustCompile("[a-z]+")

	root := []string{}
	vars := make(map[string]int)

	stack := [][]string{}

	for _, l := range strings.Split(input, "\n") {

		varNames := varRe.FindAllString(l, -1)

		if varNames[0] == "root" {
			root = varNames[1:]
		}

		if len(varNames) > 1 {
			op := string(l[11])
			varNames = append(varNames, op)
			stack = append(stack, varNames)

		} else {
			x, _ := strconv.Atoi(numberRe.FindString(l))
			vars[varNames[0]] = x
		}
	}

	testNumber := 5069554000000

	for {

		// fmt.Println("Trying:", testNumber)
		vars["humn"] = testNumber
		newStack := make([][]string, len(stack))

		for i := 0; i < len(stack); i++ {
			newStack[i] = make([]string, len(stack[i]))
			copy(newStack[i], stack[i])
		}

		// fmt.Println(newStack)
		// fmt.Println(vars)

		for len(newStack) > 0 {
			// Top element
			n := len(newStack) - 1
			op := newStack[n]
			// Pop
			newStack = newStack[:n]

			// fmt.Println(vars, op[1:len(op)-1])
			if varsExist(op[1:len(op)-1], vars) {
				if op[3] == "+" {
					vars[op[0]] = vars[op[1]] + vars[op[2]]
				} else if op[3] == "-" {
					vars[op[0]] = vars[op[1]] - vars[op[2]]
				} else if op[3] == "*" {
					vars[op[0]] = vars[op[1]] * vars[op[2]]
				} else if op[3] == "/" {
					vars[op[0]] = vars[op[1]] / vars[op[2]]
				} else {
					panic("Invalid operation!")
				}
				// fmt.Println(op, vars[op[1]], op[3], vars[op[2]], vars[op[0]])
			} else {
				stack = append([][]string{op}, stack...)
			}
		}

		// fmt.Println(vars[root[0]], "?=", vars[root[1]])
		a, ae := vars[root[0]]
		b, be := vars[root[1]]
		if !ae || !be || a != b {
			testNumber += 1
		} else {
			break
		}

	}

	fmt.Println(root)
	fmt.Println(vars[root[0]])
	fmt.Println(vars[root[1]])
	return testNumber
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
