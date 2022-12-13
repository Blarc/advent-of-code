package main

import (
	_ "embed"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func parse(input string) []interface{} {
	var cur []interface{}

	for i := 0; i < len(input); i++ {
		cs := string(input[i])
		if cs == "," {
			continue
		} else if cs == "[" {
			cur = append(cur, parse(input[i+1:]))
			stack := 1
			for stack > 0 {
				i++
				cs = string(input[i])
				if cs == "[" {
					stack++
				} else if cs == "]" {
					stack--
				}
			}
		} else if cs == "]" {
			return cur
		} else {
			// This handles number 10
			if i+1 < len(input) && input[i+1] == '0' {
				cs = string(input[i : i+2])
				i += 1
			}
			num, _ := strconv.Atoi(cs)
			cur = append(cur, num)
		}

	}

	// [[0 [] 2 10 8] [[]] [7 [[7 10] 6 2 [5 10 9 7 6] [8 8 2]] 9] [6 5 [7 0 10 []]]]
	// [[0,[],2,10,8],[[]],[7,[[7,10],6,2,[5,10,9,7,6],[8,8,2]],9],[6,5,[7,0,10,[]]]]

	// [[[[7] [7 10 9 1] [9 1 0 1 9] [9]] [2 3 [8 1 9] 9] 7 7 7] [[6]] [3 1 1 [[6 9] [7 6 0]] 10]]
	// [[[[7],[7,10,9,1],[9,1,0,1,9],[9]],[2,3,[8,1,9],9],7,7,7],[[6]],[3,1,1,[[6,9],[7,6,0]],10]]

	return cur
}

func compare(left []interface{}, right []interface{}) bool {

	for i := 0; i < len(left); i++ {

		// If the right list runs out of items first, the inputs are not in the right order.
		if i > len(right)-1 {
			return false
		}

		leftType := reflect.TypeOf(left[i]).Kind()
		rightType := reflect.TypeOf(right[i]).Kind()

		// Compare integer and integer
		if leftType == reflect.Int && rightType == reflect.Int {
			leftValue, _ := left[i].(int)
			rightValue, _ := right[i].(int)

			// fmt.Println("Compare:", leftValue, "vs", rightValue)
			if leftValue != rightValue {
				return leftValue < rightValue
			}

			// Compare slice and slice
		} else if leftType == reflect.Slice && rightType == reflect.Slice {
			leftValue, _ := left[i].([]interface{})
			rightValue, _ := right[i].([]interface{})

			comparison := compare(leftValue, rightValue)
			// fmt.Println("Compare:", leftValue, "vs", rightValue, comparison)

			return comparison
			// Compare slice and integer
		} else if leftType == reflect.Slice && rightType == reflect.Int {
			leftValue, _ := left[i].([]interface{})
			rightValue := []interface{}{right[i]}

			comparison := compare(leftValue, rightValue)
			// fmt.Println("Compare:", leftValue, "vs", rightValue, comparison)
			return comparison

			// Compare integer and slice
		} else if leftType == reflect.Int && rightType == reflect.Slice {
			leftValue := []interface{}{left[i]}
			rightValue, _ := right[i].([]interface{})

			comparison := compare(leftValue, rightValue)
			// fmt.Println("Compare:", leftValue, "vs", rightValue, comparison)
			return comparison

		} else {
			panic("Wrong types!" + leftType.String() + rightType.String())
		}

	}

	return true
}

func part1(input string) int {

	lines := strings.Split(input, "\n")

	result := 0
	pairIndex := 1
	for i := 0; i < len(lines); i += 3 {
		a := parse(lines[i][1 : len(lines[i])-1])
		b := parse(lines[i+1][1 : len(lines[i+1])-1])
		// fmt.Println(a)
		// fmt.Println(b)
		comparison := compare(a, b)

		// fmt.Println(comparison)

		if comparison {
			fmt.Println(pairIndex - 1)

			result += pairIndex
		}

		pairIndex += 1

	}

	// 5604
	return result

}

func part2(input string) int {
	return 2
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
