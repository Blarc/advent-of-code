package main

import (
	"flag"
	"fmt"
	"regexp"
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

type Node struct {
	left  string
	right string
}

func GCD(a, b uint) uint {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b uint, integers ...uint) uint {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	nav := strings.Split(lines[0], "")
	fmt.Printf("%#v\n", nav)

	// Works without numbers? hehe
	re := regexp.MustCompile(`[A-Z][A-Z][A-Z]`)

	nodes := make(map[string]Node)
	for _, l := range lines[2:] {
		n := re.FindAllString(l, -1)
		nodes[n[0]] = Node{left: n[1], right: n[2]}
	}
	//fmt.Printf("%#v\n", nodes)

	i := 0
	current := "AAA"
	for {
		index := i % len(nav)
		fmt.Println(index, nav[index])

		cn := nodes[current]

		if current == "ZZZ" {
			return i
		}

		if nav[index] == "L" {
			current = cn.left
		} else {
			current = cn.right
		}

		i++
	}
}

func part2(input string) uint {
	lines := strings.Split(input, "\n")
	nav := strings.Split(lines[0], "")
	fmt.Printf("Nav: %#v\n", nav)

	re := regexp.MustCompile(`[A-Z|0-9][A-Z|0-9][A-Z|0-9]`)

	nodes := make(map[string]Node)
	var cN []string
	for _, l := range lines[2:] {
		n := re.FindAllString(l, -1)
		nodes[n[0]] = Node{left: n[1], right: n[2]}

		if n[0][2] == 'A' {
			cN = append(cN, n[0])
		}
	}

	fmt.Printf("Start: %#v\n", cN)

	//i := 0
	//for {
	//	index := i % len(nav)
	//	//fmt.Println(index, nav[index])
	//
	//	if allFinish(cN, eN) {
	//		return i
	//	}
	//
	//	newCn := make(map[string]bool)
	//	for c, _ := range cN {
	//		cn := nodes[c]
	//		if nav[index] == "L" {
	//			newCn[cn.left] = true
	//		} else {
	//			newCn[cn.right] = true
	//		}
	//	}
	//	cN = newCn
	//	i++
	//}

	var steps []int
	for _, c := range cN {
		i := 0
		current := c
		for {
			index := i % len(nav)
			//fmt.Println(index, nav[index])

			node := nodes[current]

			if current[2] == 'Z' {
				steps = append(steps, i)
				fmt.Println("Found:", current, i)
				break
			}

			if nav[index] == "L" {
				current = node.left
			} else {
				current = node.right
			}

			i++
		}
	}
	ans := uint(steps[0])
	for i := 1; i < len(steps); i++ {
		ans = LCM(ans, uint(steps[i]))
		println(ans)
	}

	return ans

}

func allFinish(cN map[string]bool, eN map[string]bool) bool {
	for c, _ := range cN {
		if !eN[c] {
			return false
		}
	}
	return true
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
