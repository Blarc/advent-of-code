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

func getReverseOp(op string) string {
	if op == "+" {
		return "-"
	} else if op == "-" {
		return "+"
	} else if op == "*" {
		return "/"
	} else if op == "/" {
		return "*"
	} else {
		panic("Invalid operation!")
	}
}

func doReverseOp(x int, y int, op string) int {
	if op == "+" {
		return x - y
	} else if op == "-" {
		return x + y
	} else if op == "*" {
		return x / y
	} else if op == "/" {
		return x * y
	} else {
		panic("Invalid operation!")
	}
}

func doOp(x, y int, op string) int {
	if op == "+" {
		return x + y
	} else if op == "-" {
		return x - y
	} else if op == "*" {
		return x * y
	} else if op == "/" {
		return x / y
	} else {
		panic("Invalid operation!")
	}
}

func rec(a string, ops map[string][]string, ops2 map[string][]string, vars map[string]int, visited map[string]bool) int {
	// Check if variable is already in map
	v, e := vars[a]
	if e {
		return v
	}

	// Check if variable can be computed directly
	aOperation, aOperationExists := ops[a]
	if aOperationExists && !visited[a] {
		visited[a] = true

		first := rec(aOperation[1], ops, ops2, vars, visited)
		second := rec(aOperation[2], ops, ops2, vars, visited)

		r := doOp(first, second, aOperation[3])
		vars[a] = r
		return r
	}

	// Find operation that contains the variable
	// and turn it around to get the value
	cOperation := ops2[a]
	x := rec(cOperation[0], ops, ops2, vars, visited)
	vars[cOperation[0]] = x
	var y int
	if cOperation[1] == a {
		y = rec(cOperation[2], ops, ops2, vars, visited)
		vars[cOperation[2]] = y
	} else {
		y = rec(cOperation[1], ops, ops2, vars, visited)
		vars[cOperation[1]] = y
	}

	r := doReverseOp(x, y, cOperation[3])
	vars[a] = r
	return r

}

func part2(input string) int {
	numberRe := regexp.MustCompile("[0-9]+")
	varRe := regexp.MustCompile("[a-z]+")

	root := []string{}
	vars := make(map[string]int)
	vars2 := make(map[string]int)
	ops := make(map[string][]string)
	ops2 := make(map[string][]string)

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
			ops[varNames[0]] = varNames
			ops2[varNames[1]] = varNames
			ops2[varNames[2]] = varNames

		} else {
			x, _ := strconv.Atoi(numberRe.FindString(l))
			vars[varNames[0]] = x
			vars2[varNames[0]] = x
		}
	}

	newStack := make([][]string, len(stack))
	for i := 0; i < len(stack); i++ {
		newStack[i] = make([]string, len(stack[i]))
		copy(newStack[i], stack[i])
	}

	for len(newStack) > 0 {
		// Top element
		n := len(newStack) - 1
		op := newStack[n]
		// Pop
		newStack = newStack[:n]

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
		} else {
			newStack = append([][]string{op}, newStack...)
		}
	}

	result := vars[root[1]]
	fmt.Println("Must be equal to:", result)

	delete(vars2, "humn")
	delete(vars2, "root")

	delete(ops, "root")
	delete(ops2, "root")

	vars2[root[0]] = result
	vars2[root[1]] = result

	return rec("humn", ops, ops2, vars2, map[string]bool{})
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
