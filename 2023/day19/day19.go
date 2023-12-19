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

func min(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b uint64) uint64 {
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

type Part struct {
	category   string
	x, m, a, s int
}

func (p *Part) result() int {
	return p.x + p.m + p.a + p.s
}

func createWorkflow(rulesString string) func(part *Part) string {
	rules := strings.Split(rulesString, ",")
	defaultRule := rules[len(rules)-1]
	return func(part *Part) string {
		for _, rule := range rules[:len(rules)-1] {
			conditionAndResult := strings.Split(rule, ":")
			condition := conditionAndResult[0]
			result := conditionAndResult[1]
			conditionAB := strings.Split(condition, string(condition[1]))

			a := conditionAB[0]
			b, _ := strconv.Atoi(conditionAB[1])

			if a == "x" {
				if string(condition[1]) == ">" && part.x > b {
					return result
				}
				if string(condition[1]) == "<" && part.x < b {
					return result
				}
			} else if a == "m" {
				if string(condition[1]) == ">" && part.m > b {
					return result
				}
				if string(condition[1]) == "<" && part.m < b {
					return result
				}
			} else if a == "a" {
				if string(condition[1]) == ">" && part.a > b {
					return result
				}
				if string(condition[1]) == "<" && part.a < b {
					return result
				}
			} else if a == "s" {
				if string(condition[1]) == ">" && part.s > b {
					return result
				}
				if string(condition[1]) == "<" && part.s < b {
					return result
				}
			} else {
				panic("no match")
			}
		}
		return defaultRule
	}
}

func part1(input string) int {

	workflowsAndParts := strings.Split(input, "\n\n")

	workflows := make(map[string]func(*Part) string)
	for _, l := range strings.Split(workflowsAndParts[0], "\n") {
		l = strings.TrimSuffix(l, "}")
		nameAndRules := strings.Split(l, "{")
		name := nameAndRules[0]
		rules := nameAndRules[1]
		workflows[name] = createWorkflow(rules)
	}

	sum := 0
	for _, l := range strings.Split(workflowsAndParts[1], "\n") {
		l = strings.TrimPrefix(l, "{")
		l = strings.TrimSuffix(l, "}")
		categories := strings.Split(l, ",")

		newPart := &Part{category: "in"}
		for _, category := range categories {
			nameAndValue := strings.Split(category, "=")
			name := nameAndValue[0]
			value, _ := strconv.Atoi(nameAndValue[1])
			if name == "x" {
				newPart.x = value
			} else if name == "m" {
				newPart.m = value
			} else if name == "a" {
				newPart.a = value
			} else if name == "s" {
				newPart.s = value
			} else {
				panic("no match")
			}
		}

		for newPart.category != "A" && newPart.category != "R" {
			workflow := workflows[newPart.category]
			newPart.category = workflow(newPart)
		}

		if newPart.category == "A" {
			sum += newPart.result()
		}
	}

	return sum
}

type Rule struct {
	a          string
	comparison string
	value      uint64
	result     string
}

func allCombinations(category string, ranges map[string][]uint64, workflows map[string][]*Rule) uint64 {
	//fmt.Println(category)
	//for _, r := range strings.Split("xmas", "") {
	//	fmt.Println(r, ranges[r])
	//}

	if category == "A" {
		sum := uint64(1)
		for _, r := range ranges {
			sum *= r[1] - r[0]
		}
		//fmt.Println(sum)
		return sum
	}
	if category == "R" {
		return 0
	}

	oldRanges := make(map[string][]uint64)
	for k, v := range ranges {
		for _, j := range v {
			oldRanges[k] = append(oldRanges[k], j)
		}
	}

	var sum uint64
	rules := workflows[category]
	for i := 0; i < len(rules); i++ {
		rule := rules[i]
		if rule.comparison == "<" {
			old := ranges[rule.a][1]
			ranges[rule.a][1] = min(ranges[rule.a][1], rule.value-1)
			sum += allCombinations(rule.result, ranges, workflows)
			ranges[rule.a][0] = max(ranges[rule.a][0], rule.value-1)
			ranges[rule.a][1] = old
		} else if rule.comparison == ">" {
			old := ranges[rule.a][0]
			ranges[rule.a][0] = max(ranges[rule.a][0], rule.value)
			sum += allCombinations(rule.result, ranges, workflows)
			ranges[rule.a][1] = min(ranges[rule.a][1], rule.value)
			ranges[rule.a][0] = old
		} else if rule.comparison == "default" {
			sum += allCombinations(rule.result, ranges, workflows)
		} else {
			panic("no match")
		}
	}

	for k, v := range oldRanges {
		ranges[k] = v
	}
	return sum
}

func part2(input string) uint64 {
	workflowsAndParts := strings.Split(input, "\n\n")

	workflows := make(map[string][]*Rule)
	for _, l := range strings.Split(workflowsAndParts[0], "\n") {
		l = strings.TrimSuffix(l, "}")
		nameAndRules := strings.Split(l, "{")
		name := nameAndRules[0]
		rulesString := nameAndRules[1]

		rules := strings.Split(rulesString, ",")
		defaultResult := rules[len(rules)-1]

		for _, rule := range rules[:len(rules)-1] {
			conditionAndResult := strings.Split(rule, ":")
			condition := conditionAndResult[0]
			result := conditionAndResult[1]
			conditionAB := strings.Split(condition, string(condition[1]))

			a := conditionAB[0]
			b, _ := strconv.Atoi(conditionAB[1])

			workflows[name] = append(workflows[name], &Rule{a, string(condition[1]), uint64(b), result})
		}
		workflows[name] = append(workflows[name], &Rule{"default", "default", 0, defaultResult})
	}

	return allCombinations("in", map[string][]uint64{"x": {0, 4000}, "m": {0, 4000}, "a": {0, 4000}, "s": {0, 4000}}, workflows)
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
